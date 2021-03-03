package clientrpc

import (
	"testing"

	"../../models"
)

func TestClientGrpc(t *testing.T) {
	params, er := models.NewEnvParams(false)
	if er != nil {
		t.Fatal(er)
	}
	_, err := NewClientGrpc(params.GrpcAddr)
	if err != nil {
		t.Fatal(err)
	}
}
