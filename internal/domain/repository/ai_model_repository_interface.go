package repository

import (
	"agent-model-ms/internal/domain/model"
)

type AIModelRepositoryInterface interface {
	CreateModel(model model.AIModel) error
	GetModel(model_id string) (model.AIModel, error)
	GetModelList() ([]model.AIModel, error)
}
