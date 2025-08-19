package internal

import (
	"generic-backend-service/internal/config"
	"generic-backend-service/internal/grpcserver"
	"generic-backend-service/internal/httpserver"
)

type Server interface {
	ListenAndServe() error
	Shutdown()
}

func CreateServers(config *config.Config) []Server {
	servers := make([]Server, 0, 2)
	servers = append(servers, grpcserver.GetNewGrpcServer(config.GrpcInfo.GRPCAddr))
	servers = append(servers, httpserver.GetNewHttpServer(config.GrpcInfo.GRPCAddr))

	return servers
}
