package acceptanceservice

import (
	"fmt"

	tracingModel "dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-tracing.git/model"

	"pass-trougth/model"
)

type AcceptanceService struct {
	service Service
}

func New(service Service) AcceptanceService {
	return AcceptanceService{service: service}
}

func (a AcceptanceService) CreateOrUpdate(ctx tracingModel.Context, input model.AcceptanceService, headers map[string]string) (model.Message, error) {
	message, err := a.service.CreateOrUpdate(ctx, input, headers)
	if err != nil {
		return model.Message{}, fmt.Errorf("acceptanceservice.service.CreateOrUpdate(): %w", err)
	}

	return message, nil
}

func (a AcceptanceService) Delete(ctx tracingModel.Context, input model.RequestDelete, headers map[string]string) (model.Message, error) {
	message, err := a.service.Delete(ctx, input, headers)
	if err != nil {
		return model.Message{}, fmt.Errorf("acceptanceservice.service.Delete(): %w", err)
	}

	return message, nil
}

func (a AcceptanceService) Retrieve(ctx tracingModel.Context, input model.RetrieveRequest, headers map[string]string) (model.RetrieveResponse, error) {
	ms, err := a.service.Retrieve(ctx, input, headers)
	if err != nil {
		return model.RetrieveResponse{}, fmt.Errorf("acceptanceservice.service.Retrieve(): %w", err)
	}

	return ms, nil
}
