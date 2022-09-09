package model

type RetrieveOutputRequest struct {
	Documents FieldRetrieveOutputsRequest `json:"documento"`
}

type FieldRetrieveOutputRequest struct {
	OrganizationId string `json:"idOrganizacion"`
	SupplierTIN    string `json:"RucProveedor"`
	BuyerTIN       string `json:"RucComprador"`
	KeyERP         string `json:"LlaveERP"`
	UserERP        string `json:"UserERP"`
}

type FieldRetrieveOutputsRequest []FieldRetrieveOutputRequest

type RetrieveInputRequest struct {
	Documents FieldRetrieveInputsRequest `json:"documents"`
}

type FieldRetrieveInputRequest struct {
	OrganizationId string `json:"organizationId"`
	SupplierTIN    string `json:"supplierTIN"`
	BuyerTIN       string `json:"buyerTIN"`
	KeyERP         string `json:"keyERP"`
	UserERP        string `json:"userERP"`
}

type FieldRetrieveInputsRequest []FieldRetrieveInputRequest

// --------- RESPONSE ------------ //

type RetrieveOutputResponse struct {
	Data FieldRetrieveOutputsResponse `json:"data"`
}

type FieldRetrieveOutputResponse struct {
	OrganizationId  string `json:"organizationId"`
	SupplierTIN     string `json:"supplierTIN"`
	BuyerTIN        string `json:"buyerTIN"`
	DocumentId      string `json:"documentId"`
	CodeError       string `json:"codeError"`
	DataError       string `json:"dataError"`
	DocumentState   string `json:"documentState"`
	DocumentVersion string `json:"documentVersion"`
	DateTimeState   string `json:"dateTimeState"`
}

type FieldRetrieveOutputsResponse []FieldRetrieveOutputResponse

type RetrieveInputResponse struct {
	Data FieldRetrieveInputsResponse `json:"data"`
}

type FieldRetrieveInputResponse struct {
	OrganizationId  string `json:"idOrganizacion"`
	SupplierTIN     string `json:"RucProveedor"`
	BuyerTIN        string `json:"RucComprador"`
	DocumentId      string `json:"DocumentoID"`
	CodeError       string `json:"CodigoError"`
	DataError       string `json:"DataError"`
	DocumentState   string `json:"EstadoDocumento"`
	DocumentVersion string `json:"VersionDocumento"`
	DateTimeState   string `json:"FechaHoraEstado"`
}

type FieldRetrieveInputsResponse []FieldRetrieveInputResponse
