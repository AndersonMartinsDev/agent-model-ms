package cmd

import (
	"agent-model-ms/internal/application/handler"
	"agent-model-ms/internal/application/service"
	"agent-model-ms/internal/infrastructure/repository"
)

type RouterCompose struct {
	WebhookPrMsURL string
}

func NewRouterCompose() *RouterCompose {
	return &RouterCompose{
		WebhookPrMsURL: "localhost:50051",
	}
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
