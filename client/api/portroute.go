package api

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	gorpc "../../gorpc"
	clientgrpc "../grpc"

	"../../models"
	"../reader"
)

// Router to handle the API's requests
type PortRoute struct {
	helper      *Helper
	Params      *models.EnvParams
	ContentType string
}

// Error abstraction to pass information to the ResponseWriter
type HandlerError struct {
	Error string      `json:"error"`
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
}

// Success abstraction to pass information to the ResponseWriter
type SuccessMessage struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}

// Router constructor
func NewPortRoute(params *models.EnvParams) *PortRoute {
	route := &PortRoute{
		helper:      NewHelper(),
		Params:      params,
		ContentType: "application/json",
	}
	return route
}

// Answer to the POST request and handles the 'ports.json' file then sends
// to the PortDomainService via gRPC asynchronouslly
func (route *PortRoute) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", route.ContentType)
	if r.Method != http.MethodPost {
		NotAllowed(w, errors.New("Method not Allowed"))
		return
	}
	// get file from HTTP header
	r.ParseMultipartForm(32 << 20)
	f, handler, err := r.FormFile("file")
	if err != nil {
		FileError(w, err, "")
		return
	}
	separator := string(os.PathSeparator)
	now := time.Now()
	filename := fmt.Sprintf("%s%stmp%s%d%s", route.Params.BaseDir, separator, separator, now.Unix(), handler.Filename)
	log.Printf("uploaded filename: %s\n", filename)

	out, err := os.Create(filename)
	if err != nil {
		FileError(w, err, "")
		return
	}
	if _, err := io.Copy(out, f); err != nil {
		FileError(w, err, "")
		return
	}
	// close both files
	f.Close()
	out.Close()

	clientGrpc, grpcErr := clientgrpc.NewClientGrpc(route.Params.GrpcAddr)
	if grpcErr != nil {
		InternalServerError(w, grpcErr)
		return
	}
	defer clientGrpc.Close()

	reader := reader.NewPortReader()

	if er := reader.Init(filename); er != nil {
		InternalServerError(w, er)
		return
	}

	portQtd := 0
	for {
		port, erp := reader.NextPort()
		if erp != nil {
			if erp != io.EOF {
				InternalServerError(w, erp)
				return
			}
			break
		}
		portQtd++
		key := reflect.ValueOf(port).MapKeys()[0].String()
		dbPort := gorpc.Data{
			Key:         key,
			Name:        port[key].Name,
			City:        port[key].City,
			Country:     port[key].Country,
			Alias:       port[key].Alias,
			Regions:     port[key].Regions,
			Coordinates: port[key].Coordinates,
			Province:    port[key].Province,
			Timezone:    port[key].Timezone,
			Unlocs:      port[key].Unlocs,
			Code:        port[key].Code,
		}
		go func(grpcCli *clientgrpc.ClientGrpc, portToSave *gorpc.Data) {
			client := gorpc.NewPortDomainClient(grpcCli.Conn)

			rpcContext, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer cancel()

			resp, err := client.Save(rpcContext, portToSave)
			if err != nil {
				log.Printf("Port %s not persisted: %v\n", portToSave.Key, err)
			}
			if resp != nil {
				log.Printf("Result: [Code: %d, Message: %s]", resp.Code, resp.Message)
				if resp.Code == 0 {
					grpcCli.ProcessedPorts++
				}
			}

		}(clientGrpc, &dbPort)
	}

	message := fmt.Sprintf("Ports processed: %d of %d.\n", clientGrpc.ProcessedPorts, portQtd)
	code := http.StatusOK
	if clientGrpc.ProcessedPorts != portQtd {
		code = http.StatusPartialContent
	}
	portsProcessed := &SuccessMessage{
		Code:    code,
		Message: message,
	}
	encodeErr := route.helper.Encoder(w, &portsProcessed)
	log.Printf("Result: %v.\n", portsProcessed)

	if encodeErr != nil {
		BadRequest(w, encodeErr)
	}
	reader.Close()
}

// Answer to the GET request and returns the Port fond from the key provided
// or an error (for example: 404 - Not found, if the key doesn't exists in
// PortDomainService, 400 - BadRequest, if the key wasn't informed or 500 -
// Invalid Page information, if the start position is bigger than the size of the persisted list).
// To provide this Port it asks to PortDomainService via gRPC
func (route *PortRoute) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", route.ContentType)
	if r.Method != http.MethodGet {
		NotAllowed(w, errors.New("Method not Allowed"))
		return
	}

	id, err := route.helper.ParseKey(r, "id")
	if err != nil {
		BadRequest(w, err)
		return
	}
	log.Printf("Port ID to look for: %s", id)
	clientGrpc, grpcErr := clientgrpc.NewClientGrpc(route.Params.GrpcAddr)
	if grpcErr != nil {
		InternalServerError(w, grpcErr)
		return
	}
	defer clientGrpc.Close()

	client := gorpc.NewPortDomainClient(clientGrpc.Conn)

	// Setting a 150ms timeout on the RPC.
	rpcContext, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()

	key := &gorpc.PortKey{
		Key: id,
	}
	portFound, err := client.GetPort(rpcContext, key)
	if err != nil {
		log.Printf("Ports not persisted: %v\n", err)
	}
	if portFound == nil {
		erMsg := fmt.Sprintf("Port %s not found.", id)
		NotFound(w, errors.New(erMsg))
		return
	}

	encodeErr := route.helper.Encoder(w, &portFound)
	log.Printf("Port found: %v.\n", portFound)

	if encodeErr != nil {
		BadRequest(w, encodeErr)
	}
}

// Answer to the GET request and returns a list of Ports within the Page provided
// or an error (for example: 400 - Bad Request, if the page information wasn't
// informed or 500 - Invalid Page information, if the start position is bigger
// than the size of the persisted list).
// To provide this list it asks to PortDomainService via gRPC
func (route *PortRoute) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", route.ContentType)
	if r.Method != http.MethodGet {
		NotAllowed(w, errors.New("Method not Allowed"))
		return
	}
	start, err := route.helper.ParseInt32Key(r, "start")
	if err != nil {
		BadRequest(w, err)
		return
	}
	size, err := route.helper.ParseInt32Key(r, "size")
	if err != nil {
		BadRequest(w, err)
		return
	}
	log.Printf("List %d ports starting at %d.", size, start)
	clientGrpc, grpcErr := clientgrpc.NewClientGrpc(route.Params.GrpcAddr)
	if grpcErr != nil {
		InternalServerError(w, grpcErr)
		return
	}
	defer clientGrpc.Close()

	client := gorpc.NewPortDomainClient(clientGrpc.Conn)

	// Setting a 150ms timeout on the RPC.
	rpcContext, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()

	page := &gorpc.Page{
		Start: start,
		Size:  size,
	}
	portArray, err := client.GetPorts(rpcContext, page)
	if err != nil {
		log.Printf("Ports not listed: %v\n", err)
		InternalServerError(w, err)
		return
	}
	if portArray == nil || len(portArray.Ports) == 0 {
		NoContent(w, errors.New("No ports found."))
		return
	}
	portList := make(map[string]models.Data, len(portArray.Ports))
	for _, data := range portArray.Ports {
		portList[data.Key] = models.Data{
			Name:        data.Name,
			City:        data.City,
			Country:     data.Country,
			Alias:       data.Alias,
			Regions:     data.Regions,
			Coordinates: data.Coordinates,
			Province:    data.Province,
			Timezone:    data.Timezone,
			Unlocs:      data.Unlocs,
			Code:        data.Code,
		}
	}
	encodeErr := route.helper.Encoder(w, &portList)
	log.Printf("Ports list found: %v.\n", portList)

	if encodeErr != nil {
		InternalServerError(w, encodeErr)
	}
}

// Writes a 400 - Bad Request custom message in the Response Body
func BadRequest(w http.ResponseWriter, err error) {
	he := &HandlerError{
		Error: err.Error(),
		Code:  http.StatusBadRequest,
	}
	writeToResponse(w, he)
}

// Writes a "405 - Method not allowed" message in the Response Body
func NotAllowed(w http.ResponseWriter, err error) {
	he := &HandlerError{
		Error: err.Error(),
		Code:  http.StatusMethodNotAllowed,
	}
	writeToResponse(w, he)
}

// Writes a 404 - Not Found custom message in the Response Body
func NotFound(w http.ResponseWriter, err error) {
	he := &HandlerError{
		Error: err.Error(),
		Code:  http.StatusNotFound,
	}
	writeToResponse(w, he)
}

// Writes a 206 - Not Content custom message in the Response Body
func NoContent(w http.ResponseWriter, err error) {
	he := &HandlerError{
		Error: err.Error(),
		Code:  http.StatusNoContent,
	}
	writeToResponse(w, he)
}

// Writes a 500 - Internal Server Error custom message in the Response Body
func InternalServerError(w http.ResponseWriter, err error) {
	he := &HandlerError{
		Error: err.Error(),
		Code:  http.StatusInternalServerError,
	}
	writeToResponse(w, he)
}

// Writes a 404 - Bad Request custom message with the file name in the Response Body
func FileError(w http.ResponseWriter, err error, filename string) {
	he := &HandlerError{
		Code:  http.StatusBadRequest,
		Error: err.Error(),
		Data:  map[string]interface{}{"Filename": filename},
	}
	writeToResponse(w, he)
}

// Writes a the error object in the Response Body
func writeToResponse(w http.ResponseWriter, err *HandlerError) {
	helper := NewHelper()
	if err.Code == http.StatusOK || err.Code == http.StatusAccepted ||
		err.Code == http.StatusCreated {
		http.Error(w, http.StatusText(err.Code), err.Code)
	} else {
		if err != nil {
			log.Printf("error: %v", err)
		}

		encodeErr := helper.Encoder(w, err)
		if encodeErr != nil {
			http.Error(w, http.StatusText(err.Code), err.Code)
		}
	}
}
