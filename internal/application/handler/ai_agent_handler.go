package handler

import (
	"agent-model-ms/internal/application/service"
	"agent-model-ms/internal/domain/model"
	pb "agent-model-ms/proto"
	"context"
	"fmt"
)

// AIAgentHandler implementa o serviço gRPC
type AIAgentHandler struct {
	pb.UnimplementedAIAgentServiceServer
	service *service.AIAgentService
}

func NewAIAgentHandler(s *service.AIAgentService) *AIAgentHandler {
	return &AIAgentHandler{
		service: s,
	}
}
func (h *AIAgentHandler) GetAgentModelByPhone(ctx context.Context, req *pb.AgentRequest) (*pb.AgentResponse, error) {
	phoneNumber := req.GetPhoneNumber()
	agentId, err := h.service.GetAgentFromPhone(phoneNumber)
	if err != nil {
		return nil, fmt.Errorf("falha ao obter agente: %w", err)
	}

	return &pb.AgentResponse{Status: "success", AgentId: uint64(agentId)}, nil

}

func (h *AIAgentHandler) CreateAgent(ctx context.Context, req *pb.CreateAgentRequest) (*pb.CreateAgentResponse, error) {
	agent := model.AIAgent{
		Name:               req.GetName(),
		ModelId:            uint(req.GetModelId()),
		CompanyName:        req.GetCompanyName(),
		CompanyDescription: req.GetCompanyDescription(),
		BehaviourIa:        req.GetBehaviourIa(),
		CompanyUrl:         req.GetCompanyUrl(),
		UUID_user:          req.GetUuidUser(),
	}

	err := h.service.CreateBotAI(agent)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar agente: %w", err)
	}

	// Você precisa de uma forma de obter o ID do agente criado.
	// O seu serviço 'CreateBotAI' precisaria retornar o ID do novo agente.
	// Por enquanto, vamos assumir que ele retorna um ID.
	return &pb.CreateAgentResponse{Id: "id-do-agente-criado", Status: "success"}, nil
}

func (h *AIAgentHandler) UpdateAgent(ctx context.Context, req *pb.UpdateAgentRequest) (*pb.UpdateAgentResponse, error) {

	agent := model.AIAgent{
		Id:                 req.GetId(),
		Name:               req.GetName(),
		ModelId:            uint(req.GetModelId()),
		CompanyName:        req.GetCompanyName(),
		CompanyDescription: req.GetCompanyDescription(),
		BehaviourIa:        req.GetBehaviourIa(),
		CompanyUrl:         req.GetCompanyUrl(),
		UUID_user:          req.GetUuidUser(),
	}

	err := h.service.UpdateBotAI(agent)
	if err != nil {
		return nil, fmt.Errorf("falha ao atualizar agente: %w", err)
	}

	return &pb.UpdateAgentResponse{Status: "success"}, nil
}

func (h *AIAgentHandler) GetAgent(ctx context.Context, req *pb.GetAgentRequest) (*pb.GetAgentResponse, error) {
	agentID := req.GetId()
	getInstructions := req.GetInstructions()

	agent, err := h.service.GetBot(uint(agentID))
	if err != nil {
		return nil, fmt.Errorf("agente não encontrado: %w", err)
	}

	instructions := agent.Instructions

	if getInstructions {
		instructions = ""
	}

	return &pb.GetAgentResponse{
		Agent: &pb.Agent{
			Id:                 agent.Id,
			Name:               agent.Name,
			ModelId:            uint32(agent.ModelId),
			CompanyName:        agent.CompanyName,
			CompanyDescription: agent.CompanyDescription,
			BehaviourIa:        agent.BehaviourIa,
			CompanyUrl:         agent.CompanyUrl,
			UuidUser:           agent.UUID_user,
			Instructions:       instructions,
		},
	}, nil
}

func (h *AIAgentHandler) ListAgents(ctx context.Context, req *pb.ListAgentsRequest) (*pb.ListAgentsResponse, error) {
	agents, err := h.service.List(req.GetUuidUser())
	if err != nil {
		return nil, fmt.Errorf("falha ao listar agentes: %w", err)
	}

	var protoAgents []*pb.Agent
	for _, agent := range agents {
		protoAgents = append(protoAgents, &pb.Agent{
			Id:                 agent.Id,
			Name:               agent.Name,
			ModelId:            uint32(agent.ModelId),
			CompanyName:        agent.CompanyName,
			CompanyDescription: agent.CompanyDescription,
			BehaviourIa:        agent.BehaviourIa,
			CompanyUrl:         agent.CompanyUrl,
			UuidUser:           agent.UUID_user,
		})
	}

	return &pb.ListAgentsResponse{Agents: protoAgents}, nil
}

func (h *AIAgentHandler) DeleteAgent(ctx context.Context, req *pb.DeleteAgentRequest) (*pb.DeleteAgentResponse, error) {
	err := h.service.DeleteAgent(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("falha ao deletar agente: %w", err)
	}

	return &pb.DeleteAgentResponse{Status: "success"}, nil
}
