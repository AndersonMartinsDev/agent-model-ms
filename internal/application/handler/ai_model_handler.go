package handler

import (
	"agent-model-ms/internal/application/service"
	"agent-model-ms/internal/domain/model"
	pb "agent-model-ms/proto"
	"context"
	"fmt"
)

// AIModelHandler implementa o serviço gRPC
type AIModelHandler struct {
	pb.UnimplementedAIModelServiceServer
	service *service.AIModelService
}

func NewAIModelHandler(s *service.AIModelService) *AIModelHandler {
	return &AIModelHandler{
		service: s,
	}
}

func (h *AIModelHandler) CreatePerfilModel(ctx context.Context, req *pb.CreateAIModelRequest) (*pb.GetAIModelsResponse, error) {
	aiModel := model.AIModel{
		PerFilName:   req.GetPerfilName(),
		Instructions: req.GetInstructions(),
	}

	err := h.service.CreateAgent(aiModel)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar modelo de IA: %w", err)
	}
	return &pb.GetAIModelsResponse{Status: "success"}, nil
}

func (h *AIModelHandler) GetPerfilModels(ctx context.Context, req *pb.GetPerfilModelsRequest) (*pb.GetPerfilModelsResponse, error) {
	aiModels, err := h.service.GetModels()
	if err != nil {
		return nil, fmt.Errorf("falha ao listar modelos de IA: %w", err)
	}

	var protoModels []*pb.AIModel
	for _, aiModel := range aiModels {
		protoModels = append(protoModels, &pb.AIModel{
			Id:           uint32(aiModel.ID),
			PerfilName:   aiModel.PerFilName,
			Instructions: aiModel.Instructions,
		})
	}

	return &pb.GetPerfilModelsResponse{Models: protoModels}, nil
}
