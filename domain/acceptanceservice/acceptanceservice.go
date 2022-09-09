package acceptanceservice

import (
	tracingModel "dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-tracing.git/model"

	"pass-trougth/model"
)

type UseCase interface {
	Create(ctx tracingModel.Context, input model.AcceptanceServiceInput, headers map[string]string) (model.Message, error)
	Update(ctx tracingModel.Context, input model.AcceptanceServiceInput, headers map[string]string) (model.Message, error)
	Delete(ctx tracingModel.Context, input model.RequestDeleteInput, headers map[string]string) (model.Message, error)
	Retrieve(ctx tracingModel.Context, input model.RetrieveInputRequest, headers map[string]string) (model.RetrieveOutputResponse, error)
}

type Service interface {
	Create(ctx tracingModel.Context, input model.AcceptanceServiceInput, headers map[string]string) (model.Message, error)
	Update(ctx tracingModel.Context, input model.AcceptanceServiceInput, headers map[string]string) (model.Message, error)
	Delete(ctx tracingModel.Context, input model.RequestDeleteInput, headers map[string]string) (model.Message, error)
	Retrieve(ctx tracingModel.Context, input model.RetrieveInputRequest, headers map[string]string) (model.RetrieveOutputResponse, error)
}
