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

func (a AcceptanceService) Create(ctx tracingModel.Context, input model.AcceptanceServiceInput, headers map[string]string) (model.Message, error) {
	message, err := a.service.Create(ctx, input, headers)
	if err != nil {
		return model.Message{}, fmt.Errorf("acceptanceservice.service.Create(): %w", err)
	}

	return message, nil
}

func (a AcceptanceService) Update(ctx tracingModel.Context, input model.AcceptanceServiceInput, headers map[string]string) (model.Message, error) {
	message, err := a.service.Update(ctx, input, headers)
	if err != nil {
		return model.Message{}, fmt.Errorf("acceptanceservice.service.Update(): %w", err)
	}

	return message, nil
}

func (a AcceptanceService) Delete(ctx tracingModel.Context, input model.RequestDeleteInput, headers map[string]string) (model.Message, error) {
	message, err := a.service.Delete(ctx, input, headers)
	if err != nil {
		return model.Message{}, fmt.Errorf("acceptanceservice.service.Delete(): %w", err)
	}

	return message, nil
}

func (a AcceptanceService) Retrieve(ctx tracingModel.Context, input model.RetrieveInputRequest, headers map[string]string) (model.RetrieveOutputResponse, error) {
	ms, err := a.service.Retrieve(ctx, input, headers)
	if err != nil {
		return model.RetrieveOutputResponse{}, fmt.Errorf("acceptanceservice.service.Retrieve(): %w", err)
	}

	return ms, nil
}
