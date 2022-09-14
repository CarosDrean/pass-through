package model

import "time"

type AcceptanceService struct {
	DocumentHAS DocumentHas `json:"documentHAS" json-out:"DOCUMENTOHAS"`
}

type DocumentHas struct {
	AddressSupplier         string                  `json:"addressSupplier" json-out:"DireccionProveedor"`
	ReasonSocialSupplier    string                  `json:"reasonSocialSupplier" json-out:"RazonSocialProveedor"`
	SupplierTIN             string                  `json:"supplierTIN" json-out:"RucProveedor"`
	ServiceAcceptanceSheets ServiceAcceptanceSheets `json:"serviceAcceptanceSheets" json-out:"HojaAceptacionServicio"`
}

type ServiceAcceptanceSheet struct {
	AcceptedBy        string        `json:"acceptedBy" json-out:"AceptadoPor"`
	ApprovedBy        string        `json:"approvedBy" json-out:"AprobadoPor"`
	AttentionAreaCode string        `json:"attentionAreaCode" json-out:"CodigoAreaAtencion"`
	ClientTIN         string        `json:"clientTIN" json-out:"RucCliente"`
	CurrencyHAS       string        `json:"currencyHAS" json-out:"MonedaHAS"`
	DocumentMaterial  int64         `json:"documentMaterial" json-out:"DocumentoMaterial"`
	FiscalYear        int           `json:"fiscalYear" json-out:"AnioFiscal"`
	NumberHAS         int           `json:"numberHAS" json-out:"NumeroHAS"`
	Observation       string        `json:"observation" json-out:"Observaciones"`
	ServiceType       string        `json:"serviceType" json-out:"TipoServicio"`
	Society           int           `json:"society" json-out:"Sociedad"`
	SupplierCodeERP   int64         `json:"supplierCodeERP" json-out:"CodigoERPProveedor"`
	UserCreationHAS   string        `json:"userCreationHAS" json-out:"UsuarioCreacionHAS"`
	AcceptanceDate    time.Time     `json:"acceptanceDate" json-out:"FechaAceptacion"`
	AccountingDateHAS time.Time     `json:"accountingDateHAS" json-out:"FechaContabilizacionHAS"`
	BroadcastDate     time.Time     `json:"broadcastDate" json-out:"FechaEmision"`
	ServiceEndDate    time.Time     `json:"serviceEndDate" json-out:"FechaFinServicio"`
	ServiceStartDate  time.Time     `json:"serviceStartDate" json-out:"FechaInicioServicio"`
	ItemHAS           ItemsHas      `json:"itemHAS" json-out:"ItemHAS"`
	AttachedFiles     AttachedFiles `json:"attachedFiles" json-out:"ArchivoAdjunto"`
}

type ServiceAcceptanceSheets []ServiceAcceptanceSheet

type AttachedFile struct {
	URL         string `json:"URL" json-out:"URL"`
	Description string `json:"description" json-out:"Descripcion"`
	Name        string `json:"name" json-out:"Nombre"`
}

type AttachedFiles []AttachedFile

type ItemHas struct {
	AttendedAmount               float64           `json:"attendedAmount" json-out:"CantidadAtendida"`
	CodeConditionPayment         string            `json:"codeConditionPayment" json-out:"CodigoCondicionPago"`
	DocumentCurrency             string            `json:"documentCurrency" json-out:"MonedaDocumento"`
	DocumentReference            int64             `json:"documentReference" json-out:"DocumentoReferencia"`
	EndDateServiceItem           string            `json:"endDateServiceItem" json-out:"FechaFinServicioItem"`
	FinalDeliveryIndicator       string            `json:"finalDeliveryIndicator" json-out:"IndicadorEntregaFinal"`
	ItemDescription              string            `json:"itemDescription" json-out:"DescripcionItem"`
	ItemNumber                   string            `json:"itemNumber" json-out:"NumeroItem"`
	ItemNumberOC                 string            `json:"itemNumberOC" json-out:"NumeroItemOC"`
	ItemQuantity                 int               `json:"itemQuantity" json-out:"CantidadItem"`
	NumberOC                     int64             `json:"numberOC" json-out:"NumeroOC"`
	ObservationItem              string            `json:"observationItem" json-out:"ObservacionItem"`
	PaymentCondition             string            `json:"paymentCondition" json-out:"CondicionPago"`
	PriceItem                    float64           `json:"priceItem" json-out:"PrecioItem"`
	ReferenceType                int               `json:"referenceType" json-out:"TipoReferencia"`
	ServiceCode                  string            `json:"serviceCode" json-out:"CodigoServicio"`
	ServicesGroup                string            `json:"servicesGroup" json-out:"GrupoServicios"`
	SubTotalItemCurrencyDocument float64           `json:"subTotalItemCurrencyDocument" json-out:"SubTotalItemMonedaDocumento"`
	SubTotalItemLocalCurrency    float64           `json:"subTotalItemLocalCurrency" json-out:"SubTotalItemMonedaLocal"`
	UnitMeasureItem              string            `json:"unitMeasureItem" json-out:"UnidadMedidaItem"`
	SubItemHAS                   SubItemsHas       `json:"subItemHAS" json-out:"SubItemHAS"`
	MaterialDocuments            MaterialDocuments `json:"materialDocuments" json-out:"DocumentosMaterial"`
}

type ItemsHas []ItemHas

type MaterialDocument struct {
	Amount                 float64 `json:"amount" json-out:"Cantidad"`
	Exercise               int     `json:"exercise" json-out:"Ejercicio"`
	MovementClass          int     `json:"movementClass" json-out:"ClaseMovimiento"`
	NumberDocumentMaterial int64   `json:"numberDocumentMaterial" json-out:"NumeroDocumentoMaterial"`
	OperationType          string  `json:"operationType" json-out:"TipoOperacion"`
	Position               string  `json:"position" json-out:"Posicion"`
	SubTotal               float64 `json:"subTotal" json-out:"Subtotal"`
}

type MaterialDocuments []MaterialDocument

type SubItemHas struct {
	SubItemDescription string  `json:"subItemDescription" json-out:"DescripcionSubItem"`
	SubItemNumber      string  `json:"subItemNumber" json-out:"NumeroSubItem"`
	SubItemPrice       int     `json:"subItemPrice" json-out:"PrecioSubItem"`
	SubItemQuantity    float64 `json:"subItemQuantity" json-out:"CantidadSubItem"`
	SubItemServiceCode string  `json:"subItemServiceCode" json-out:"CodigoServicioSubItem"`
	SubItemUnitMeasure string  `json:"subItemUnitMeasure" json-out:"UnidadMedidaSubItem"`
	SubTotalSubItem    float64 `json:"subTotalSubItem" json-out:"SubTotalSubItem"`
}

type SubItemsHas []SubItemHas
