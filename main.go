package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/sunjayBhatia/yages/yages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	ipnport := "0.0.0.0:9000"
	if ie := os.Getenv("YAGES_BIND"); ie != "" {
		ipnport = ie
	}

	ln, err := net.Listen("tcp", ipnport)
	if err != nil {
		log.Fatalf("Failed to listen on %s due to %v", ipnport, err)
	}

	server := grpc.NewServer()
	yages.RegisterEchoServer(server, &EchoService{})
	grpc_health_v1.RegisterHealthServer(server, &HealthCheckService{})
	reflection.Register(server)

	log.Printf("Starting YAGES serve on %s\n", ipnport)
	if err := server.Serve(ln); err != nil {
		log.Fatalf("YAGES serve failed due to %v", err)
	}
}

// HealthCheckService implements grpc_health_v1.HealthServer
type HealthCheckService struct{}

func (s *HealthCheckService) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (s *HealthCheckService) Watch(req *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	return server.Send(&grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	})
}

// EchoService implements yages.EchoServer
type EchoService struct {
	yages.UnimplementedEchoServer
}

// Ping returns a "pong" (constant message).
func (s *EchoService) Ping(ctx context.Context, _ *yages.Empty) (*yages.Content, error) {
	return &yages.Content{Text: "pong"}, nil
}

// Reverse returns the message it received in reverse order.
func (s *EchoService) Reverse(ctx context.Context, msg *yages.Content) (*yages.Content, error) {
	revstr := func(s string) string {
		r := []rune(s)
		for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}
	return &yages.Content{Text: revstr(msg.Text)}, nil
}
