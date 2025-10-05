package main

import (
	"agent-model-ms/cmd"
	"agent-model-ms/internal/infrastructure/configuration"
	"log"
	"log/slog"
	"net"

	pb_agent "agent-model-ms/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// configuration.LoadEnv()
	configuration.LoadLogger()
	configuration.LoadDatabase()

	router_compose := cmd.NewRouterCompose()

	aiAgentHandler := router_compose.HandlerAIAgentConfiguration()
	aiModelHandler := router_compose.HandlerAIModelConfiguration()

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Falha ao iniciar o servidor gRPC: %v", err)
	}

	s := grpc.NewServer()

	pb_agent.RegisterAIAgentServiceServer(s, aiAgentHandler)
	pb_agent.RegisterAIModelServiceServer(s, aiModelHandler)

	reflection.Register(s)

	slog.Info("Servidor gRPC do Agent-Model iniciado na porta 50052...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falha ao servir: %v", err)
	}
}
