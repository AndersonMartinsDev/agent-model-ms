package repository

import (
	"agent-model-ms/internal/domain/model"
)

type AIAgentRepositoryInterface interface {
	Createagent(agent model.AIAgent) error
	UpdateAgent(agent model.AIAgent) error
	GetAgent(agent_id uint) (model.AIAgent, error)
	List(user_uuid string) ([]model.AIAgent, error)
	Delete(agent_id string) error
	GetAgentAndPerfil(agent_id uint64) (model.AIAgent, error)
	GetAgentFromPhone(phone string) (uint, error)
}
