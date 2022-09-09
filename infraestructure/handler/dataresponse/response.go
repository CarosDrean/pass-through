package dataresponse

import (
	"net/http"

	"dev.azure.com/ciaalicorp/CIA-Funciones/cia-library-domain-alicorp.git/customerrror"

	"pass-trougth/model"
)

type DataResponse struct{}

func New() DataResponse {
	return DataResponse{}
}

func (d DataResponse) GetErrorResponse(err *customerror.Error) (int, any) {
	if !err.HasCode() {
		err.SetCode(model.InernalServerErrorCode)
	}

	if !err.HasAPIMessage() {
		err.SetErrorAsAPIMessage()
	}

	if !err.HasStatusHTTP() {
		err.SetStatusHTTP(http.StatusInternalServerError)
	}

	errorResponse := model.ErrorApi{
		Code:        string(err.Code()),
		Description: string(err.APIMessage()),
	}

	if err.HasFields() {
		for _, field := range err.Fields {
			errorResponse.ExceptionDetails = append(errorResponse.ExceptionDetails, model.ExceptionDetail{
				Component:   field.Field,
				Description: field.Description,
				Endpoint:    "",
			})
		}
	}

	return err.StatusHTTP(), errorResponse
}

func (d DataResponse) GetErrorResponseUnexpected(err *customerror.Error) any {
	return model.ErrorApi{Code: string(err.Code()), Description: string(err.APIMessage())}
}
