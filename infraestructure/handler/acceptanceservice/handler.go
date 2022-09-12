package acceptanceservice

import (
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-common.git/apiutils/request"
	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-common.git/apiutils/response"
	tracingModel "dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-extension-rkgin-tracing.git/model"
	"github.com/gin-gonic/gin"

	"pass-trougth/domain/acceptanceservice"
	"pass-trougth/infraestructure/handler/dataresponse"
	"pass-trougth/model"
)

type handler struct {
	useCase  acceptanceservice.UseCase
	response response.ApiResponse
}

func newHandler(useCase acceptanceservice.UseCase, response response.ApiResponse) handler {
	return handler{useCase: useCase, response: response}
}

func (h handler) createOrUpdate(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	originData := c.GetHeader("origen_datos")
	typeBusiness := c.GetHeader("tipo_empresa")
	organizationID := c.GetHeader("org_id")

	ctx := tracingModel.NewContext()
	ctx.Data[tracingModel.InfoOperationID] = "create"

	headersRequired := map[string]string{
		"origen_datos": originData,
		"tipo_empresa": typeBusiness,
		"org_id":       organizationID,
	}

	if err := request.ValidateMissingParam(headersRequired); err != nil {
		c.JSON(h.response.ParamFailed(c, ctx, err))
		return
	}

	var input model.AcceptanceService
	if err := c.BindJSON(&input); err != nil {
		c.JSON(h.response.BindFailed(c, ctx, err))
		return
	}

	_, err := h.useCase.CreateOrUpdate(ctx, input, headersRequired)
	if err != nil {
		c.JSON(h.response.Error(c, ctx, "h.useCase.CreateOrUpdate()", err))
		return
	}

	c.JSON(h.response.Accepted(c, ctx, dataresponse.GetAcceptedDefault()))
}

func (h handler) delete(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	originData := c.GetHeader("origen_datos")
	typeBusiness := c.GetHeader("tipo_empresa")
	organizationID := c.GetHeader("org_id")

	ctx := tracingModel.NewContext()
	ctx.Data[tracingModel.InfoOperationID] = "delete"

	headersRequired := map[string]string{
		"origen_datos": originData,
		"tipo_empresa": typeBusiness,
		"org_id":       organizationID,
	}

	if err := request.ValidateMissingParam(headersRequired); err != nil {
		c.JSON(h.response.ParamFailed(c, ctx, err))
		return
	}

	var input model.RequestDelete
	if err := c.BindJSON(&input); err != nil {
		c.JSON(h.response.BindFailed(c, ctx, err))
		return
	}

	_, err := h.useCase.Delete(ctx, input, headersRequired)
	if err != nil {
		c.JSON(h.response.Error(c, ctx, "h.useCase.Delete()", err))
		return
	}

	c.JSON(h.response.Accepted(c, ctx, dataresponse.GetAcceptedDefault()))
}

func (h handler) retrieve(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	originData := c.GetHeader("origen_datos")
	typeBusiness := c.GetHeader("tipo_empresa")
	organizationID := c.GetHeader("org_id")

	ctx := tracingModel.NewContext()
	ctx.Data[tracingModel.InfoOperationID] = "retrieve"

	headersRequired := map[string]string{
		"origen_datos": originData,
		"tipo_empresa": typeBusiness,
		"org_id":       organizationID,
	}

	if err := request.ValidateMissingParam(headersRequired); err != nil {
		c.JSON(h.response.ParamFailed(c, ctx, err))
		return
	}

	var input model.RetrieveRequest
	if err := c.BindJSON(&input); err != nil {
		c.JSON(h.response.BindFailed(c, ctx, err))
		return
	}

	ms, err := h.useCase.Retrieve(ctx, input, headersRequired)
	if err != nil {
		c.JSON(h.response.Error(c, ctx, "h.useCase.Retrieve()", err))
		return
	}

	c.JSON(h.response.OK(c, ctx, ms))
}
