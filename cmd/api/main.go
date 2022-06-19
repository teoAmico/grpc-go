package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)


func main() {
	log := hclog.Default()

	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs := grpc.NewServer()
	
	hs := health.NewServer()

	hs.SetServingStatus("",healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(gs, hs)
	
	reflection.Register(gs)
	gs.Serve(l)
}
