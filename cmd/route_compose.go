package cmd

import (
	"agent-model-ms/internal/application/handler"
	"agent-model-ms/internal/application/service"
	"agent-model-ms/internal/infrastructure/repository"
)

type RouterCompose struct {
}

func NewRouterCompose() *RouterCompose {
	return &RouterCompose{}
}

func (manager RouterCompose) HandlerAIAgentConfiguration() *handler.AIAgentHandler {
	agentRepo := repository.NewAIAgentRepository()
	agentService := service.NewAIAgentService(agentRepo)
	return handler.NewAIAgentHandler(agentService)
}

func (manager RouterCompose) HandlerAIModelConfiguration() *handler.AIModelHandler {
	modelRepo := repository.NewAIModelRepository()
	modelService := service.NewAIModelService(modelRepo)
	return handler.NewAIModelHandler(modelService)
}
