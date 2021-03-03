package clientrpc

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
)

// gRPC client to communicate with the PortDomainService
type ClientGrpc struct {
	Conn           *grpc.ClientConn
	ProcessedPorts int
}

// gRPC client's constructor. Returns an error if it can't connect with the gRPC service
func NewClientGrpc(grpcAddr string) (*ClientGrpc, error) {
	grpcResolver := manual.NewBuilderWithScheme("portDomain")
	address := []resolver.Address{}
	address = append(address, resolver.Address{Addr: grpcAddr})
	conn, err := grpc.Dial(grpcResolver.Scheme()+":///test.server",
		grpc.WithInsecure(),
		grpc.WithResolvers(grpcResolver),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		log.Printf("Error creating grpc.Dial: %v", err)
		return nil, err
	}
	grpcResolver.UpdateState(resolver.State{Addresses: address})

	client := &ClientGrpc{
		Conn:           conn,
		ProcessedPorts: 0,
	}
	return client, nil
}

func (client *ClientGrpc) Close() {
	if client.Conn != nil {
		defer client.Conn.Close()
	}
}
