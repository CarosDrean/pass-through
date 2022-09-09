package model

const (
	InernalServerErrorCode = "SKN0005"
	BadRequestCode         = "SKN0007"
)

const (
	InternalServerErrorDescription = "Ocurrió un error inesperado. Por favor contactarse con el Soporte Tecnico."
	BadRequestDescription          = "Parámetros obligatorios."
)

type ErrorApi struct {
	Code             string           `json:"code"`
	Description      string           `json:"description"`
	ErrorType        string           `json:"errorType"`
	ExceptionDetails ExceptionDetails `json:"exceptionDetails"`
}

type ExceptionDetail struct {
	Component   string `json:"component"`
	Description string `json:"description"`
	Endpoint    string `json:"endpoint"`
}

type ExceptionDetails []ExceptionDetail
