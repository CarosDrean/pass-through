package acceptanceservice

import (
	tracingModel "dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-tracing.git/model"

	"pass-trougth/model"
)

type UseCase interface {
	CreateOrUpdate(ctx tracingModel.Context, input model.AcceptanceService, headers map[string]string) (model.Message, error)
	Delete(ctx tracingModel.Context, input model.RequestDelete, headers map[string]string) (model.Message, error)
	Retrieve(ctx tracingModel.Context, input model.RetrieveRequest, headers map[string]string) (model.RetrieveResponse, error)
}

type Service interface {
	CreateOrUpdate(ctx tracingModel.Context, input model.AcceptanceService, headers map[string]string) (model.Message, error)
	Delete(ctx tracingModel.Context, input model.RequestDelete, headers map[string]string) (model.Message, error)
	Retrieve(ctx tracingModel.Context, input model.RetrieveRequest, headers map[string]string) (model.RetrieveResponse, error)
}
