package service

import (
	"agent-model-ms/internal/domain/model"
	"agent-model-ms/internal/domain/repository"
)

type AIModelService struct {
	repo repository.AIModelRepositoryInterface
}

func NewAIModelService(repoInterface repository.AIModelRepositoryInterface) *AIModelService {
	return &AIModelService{
		repo: repoInterface,
	}
}

func (serv AIModelService) CreateAgent(model model.AIModel) error {
	return serv.repo.CreateModel(model)
}

func (serv AIModelService) GetModels() ([]model.AIModel, error) {
	return serv.repo.GetModelList()
}
