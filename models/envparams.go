package models

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

// Handles the environment variables with the microservices
type EnvParams struct {
	BaseDir  string
	GrpcAddr string
	ApiAddr  string
}

// Constructor of the environment parameters handler
func NewEnvParams(ignoreClientApi bool) (*EnvParams, error) {
	path, er := os.Executable()
	if er != nil {
		log.Printf("Error: Client API %s\n", er.Error())
		return nil, er
	}
	dir, e_r := filepath.Abs(filepath.Dir(path))
	if e_r != nil {
		log.Printf("Error: Client API %s\n", e_r.Error())
		return nil, e_r
	}
	host, hostExists := os.LookupEnv("GRPC_ADDR")
	if !hostExists {
		err := errors.New("Error: GRPC_ADDR not set.")
		log.Println(err.Error())
		return nil, err
	}
	api, apiExists := os.LookupEnv("CLIENT_API_ADDR")
	if !ignoreClientApi && !apiExists {
		erApi := errors.New("Error: CLIENT_API_ADDR not set.")
		log.Println(erApi.Error())
		return nil, erApi
	}
	ep := &EnvParams{
		BaseDir:  dir,
		GrpcAddr: host,
		ApiAddr:  api,
	}
	return ep, nil
}
