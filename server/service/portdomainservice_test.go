package service

import (
	"context"
	"log"
	"net"
	"testing"

	gorpc "../../gorpc"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	server := NewPortDomainService()
	gorpc.RegisterPortDomainServer(s, server)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestPortDomainService(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := gorpc.NewPortDomainClient(conn)
	data1 := dummyData()
	resp, err := client.Save(ctx, &data1)
	if err != nil {
		t.Fatal(err)
	}
	code, expected := int32(0), "Port 'TESTE1' saved."
	if code != resp.Code || expected != resp.Message {
		t.Fatalf("Result diff from expected. Got: %d, '%s' Want: %d, '%s'",
			resp.Code, resp.Message, code, expected)
	}
	portKey := &gorpc.PortKey{
		Key: data1.Key,
	}
	data2, err2 := client.GetPort(ctx, portKey)
	if err2 != nil {
		t.Fatal(err2)
	}
	data2.Key = "TEST2"
	resp, err = client.Save(ctx, data2)
	if err != nil {
		t.Fatal(err)
	}
	expected = "Port 'TEST2' saved."
	if code != resp.Code || expected != resp.Message {
		t.Fatalf("Result diff from expected. Got: %d, '%s' Want: %d, '%s'",
			resp.Code, resp.Message, code, expected)
	}
	page := &gorpc.Page{
		Start: 0,
		Size:  10,
	}
	ports, errP := client.GetPorts(ctx, page)
	if err != nil {
		t.Fatal(errP)
	}
	if len(ports.Ports) != 2 {
		t.Errorf("Service listed invalid page size: %d", len(ports.Ports))
	}
}

func dummyData() gorpc.Data {
	return gorpc.Data{
		Key:         "TESTE1",
		Name:        "Data 1",
		City:        "City 1",
		Country:     "Country 1",
		Coordinates: []float32{12.0, 32.9},
		Province:    "Prov 1",
		Timezone:    "Timezone 1",
		Code:        "Code 1",
	}
}
