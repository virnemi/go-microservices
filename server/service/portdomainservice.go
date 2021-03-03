package service

import (
	"context"
	"fmt"
	"reflect"

	gorpc "../../gorpc"
	"../../models"
	repository "../repository"
)

// Service to provide Ports via gRPC.
// It has an InMemory database
type PortDomainService struct {
	gorpc.UnimplementedPortDomainServer
	InMemoryDb *repository.PortsMap
}

// Service constructor
func NewPortDomainService() *PortDomainService {
	db := repository.NewPortsMap()
	s := &PortDomainService{
		InMemoryDb: db,
	}
	return s
}

// Persists Ports in its database, called via gRPC
func (server *PortDomainService) Save(ctx context.Context, in *gorpc.Data) (*gorpc.Result, error) {
	data := server.ConvertGrpc(in)
	_, err := server.InMemoryDb.Add(in.Key, *data)
	if err != nil {
		err = server.InMemoryDb.Update(in.Key, *data)
	}
	if err != nil {
		return nil, err
	}
	result := &gorpc.Result{
		Code:    0,
		Message: fmt.Sprintf("Port '%s' saved.", in.Key),
	}
	return result, nil
}

// Returns a specific Port, searched by the provided key, called via gRPC
func (server *PortDomainService) GetPort(ctx context.Context, portKey *gorpc.PortKey) (*gorpc.Data, error) {
	port, err := server.InMemoryDb.Get(portKey.Key)
	if err != nil {
		return nil, err
	}
	model := (*port)[portKey.Key]
	data := server.ConvertModel(portKey.Key, &model)
	return data, nil
}

// Returns a list of Ports, filtered by a page information (start and size), called via gRPC
func (server *PortDomainService) GetPorts(ctx context.Context, page *gorpc.Page) (*gorpc.PortArray, error) {
	p := models.Page{
		Start: page.Start,
		Size:  page.Size,
	}
	list, err := server.InMemoryDb.GetPage(p)
	if err != nil {
		return nil, err
	}
	arr := gorpc.PortArray{
		Ports: make([]*gorpc.Data, 0),
	}
	for _, value := range list {
		key := reflect.ValueOf(value).MapKeys()[0].String()
		model := value[key]
		data := server.ConvertModel(key, &model)
		arr.Ports = append(arr.Ports, data)
	}
	return &arr, nil
}

// Converts a gRPC Port to the model, used to help persisting data.
func (server *PortDomainService) ConvertGrpc(in *gorpc.Data) *models.Data {
	data := &models.Data{
		Name:        in.Name,
		City:        in.City,
		Country:     in.Country,
		Alias:       in.Alias,
		Regions:     in.Regions,
		Coordinates: in.Coordinates,
		Province:    in.Province,
		Timezone:    in.Timezone,
		Unlocs:      in.Unlocs,
		Code:        in.Code,
	}
	return data
}

// Converts a Port model to a gRPC Port, used to help providing data.
func (server *PortDomainService) ConvertModel(key string, in *models.Data) *gorpc.Data {
	data := &gorpc.Data{
		Key:         key,
		Name:        in.Name,
		City:        in.City,
		Country:     in.Country,
		Alias:       in.Alias,
		Regions:     in.Regions,
		Coordinates: in.Coordinates,
		Province:    in.Province,
		Timezone:    in.Timezone,
		Unlocs:      in.Unlocs,
		Code:        in.Code,
	}
	return data
}
