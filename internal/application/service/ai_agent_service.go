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

func (serv AIAgentService) GetBehaviorAgentIa() map[int32]string {
	behavioural := make(map[int32]string)
	behavioural[int32(model.Formal)] = model.Formal.String()
	behavioural[int32(model.Informal)] = model.Informal.String()
	behavioural[int32(model.Profissional)] = model.Profissional.String()
	behavioural[int32(model.Amigavel)] = model.Amigavel.String()
	behavioural[int32(model.Educado)] = model.Educado.String()
	behavioural[int32(model.Conciso)] = model.Conciso.String()
	behavioural[int32(model.Detalhado)] = model.Detalhado.String()
	behavioural[int32(model.Divertido)] = model.Divertido.String()
	behavioural[int32(model.Respeitoso)] = model.Respeitoso.String()
	return behavioural
}
