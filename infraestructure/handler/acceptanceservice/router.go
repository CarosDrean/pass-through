package acceptanceservice

import (
	"net/http"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-common.git/apiutils/response"
	"github.com/gin-gonic/gin"

	"pass-trougth/domain/acceptanceservice"
	"pass-trougth/infraestructure/handler/dataresponse"
	acceptanceserviceService "pass-trougth/infraestructure/service/acceptanceservice"
	"pass-trougth/model"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler(specification)

	publicRoutes(specification.Api, handler)
}

func buildHandler(specification model.RouterSpecification) handler {
	useCase := acceptanceservice.New(acceptanceserviceService.New(specification.Config))

	return newHandler(useCase, buildResponseApi())
}

func publicRoutes(api *gin.Engine, h handler, middlewares ...gin.HandlerFunc) {
	routes := api.Group("v1/pass-through", middlewares...)

	routes.POST("", h.createOrUpdate)
	routes.PUT("", h.createOrUpdate)
	routes.DELETE("", h.delete)
	routes.POST("/retrieve", h.retrieve)
}

func buildResponseApi() response.ApiResponse {
	responseAPI := response.New(dataresponse.New())
	responseAPI.OverrideParamFailedResponses(response.TypeCode{
		StatusHttp:   http.StatusBadRequest,
		Code:         model.BadRequestCode,
		Description:  model.BadRequestDescription,
		IsNeedSecond: false,
	})
	responseAPI.OverrideUnexpectedErrorResponses(response.TypeCode{
		StatusHttp:  http.StatusInternalServerError,
		Code:        model.InernalServerErrorCode,
		Description: model.InternalServerErrorDescription,
	})
	responseAPI.OverrideBindFailedResponses(response.TypeCode{
		StatusHttp:  http.StatusBadRequest,
		Code:        model.BadRequestCode,
		Description: model.BadRequestDescription,
	})

	return responseAPI
}
