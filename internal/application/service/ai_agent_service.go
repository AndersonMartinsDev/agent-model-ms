package service

import (
	"agent-model-ms/internal/domain/model"
	"agent-model-ms/internal/domain/repository"
)

type AIAgentService struct {
	repo repository.AIAgentRepositoryInterface
}

func NewAIAgentService(repoInterface repository.AIAgentRepositoryInterface) *AIAgentService {
	return &AIAgentService{
		repo: repoInterface,
	}
}

func (serv AIAgentService) CreateBotAI(agent model.AIAgent) error {
	return serv.repo.Createagent(agent)
}
func (serv AIAgentService) UpdateBotAI(agent model.AIAgent) error {
	return serv.repo.UpdateAgent(agent)
}
func (serv AIAgentService) GetBot(agent_id uint) (model.AIAgent, error) {
	return serv.repo.GetAgent(agent_id)
}
func (serv AIAgentService) List(user_uuid string) ([]model.AIAgent, error) {
	return serv.repo.List(user_uuid)
}
func (serv AIAgentService) DeleteAgent(agent_id string) error {
	return serv.repo.Delete(agent_id)
}
func (serv AIAgentService) GetAgentAndPerfil(agent_id uint64) (model.AIAgent, error) {
	return serv.repo.GetAgentAndPerfil(agent_id)
}

func (serv AIAgentService) GetAgentFromPhone(phone string) (uint, error) {
	return serv.repo.GetAgentFromPhone(phone)
}
