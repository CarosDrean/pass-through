package model

type RetrieveRequest struct {
	Documents FieldsRetrieveRequest `json:"documents" json-out:"documento"`
}

type FieldRetrieveRequest struct {
	OrganizationId string `json:"organizationId" json-out:"idOrganizacion"`
	SupplierTIN    string `json:"supplierTIN" json-out:"RucProveedor"`
	BuyerTIN       string `json:"buyerTIN" json-out:"RucComprador"`
	KeyERP         string `json:"keyERP" json-out:"LlaveERP"`
	UserERP        string `json:"userERP" json-out:"UserERP"`
}

type FieldsRetrieveRequest []FieldRetrieveRequest

// --------- RESPONSE ------------ //

type RetrieveResponse struct {
	Data FieldsRetrieveResponse `json:"data" json-out:"data"`
}

type FieldRetrieveResponse struct {
	OrganizationId  string `json:"organizationId" json-out:"idOrganizacion"`
	SupplierTIN     string `json:"supplierTIN" json-out:"RucProveedor"`
	BuyerTIN        string `json:"buyerTIN" json-out:"RucComprador"`
	DocumentId      string `json:"documentId" json-out:"DocumentoID"`
	CodeError       string `json:"codeError" json-out:"CodigoError"`
	DataError       string `json:"dataError" json-out:"DataError"`
	DocumentState   string `json:"documentState" json-out:"EstadoDocumento"`
	DocumentVersion string `json:"documentVersion" json-out:"VersionDocumento"`
	DateTimeState   string `json:"dateTimeState" json-out:"FechaHoraEstado"`
}

type FieldsRetrieveResponse []FieldRetrieveResponse
