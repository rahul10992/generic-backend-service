package grpcserver

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "generic-backend-service/internal/gen/health/v1"
)

type GrpcServer struct {
	addr   string
	server *grpc.Server
}

func GetNewGrpcServer(addr string) *GrpcServer {
	if addr == "" {
		panic("Addr is empty for Grpc Server")
	}
	return &GrpcServer{addr: addr}
}

func (g *GrpcServer) ListenAndServe() error {
	listener, err := net.Listen("tcp", g.addr)
	if err != nil {
		return err
	}
	g.server = grpc.NewServer()
	if err := g.RegisterRoutes(); err != nil {
		log.Fatalf("failed to register routes: %v", err)
	}
	if err := g.server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}

var errServerIsNil = fmt.Errorf("server is not initialised")

func (g *GrpcServer) RegisterRoutes() error {
	if g.server == nil {
		return errServerIsNil
	}

	pb.RegisterHealthServiceServer(g.server, pb.UnimplementedHealthServiceServer{})
	return nil
}

func (g *GrpcServer) Shutdown() {
	g.server.Stop()
}
