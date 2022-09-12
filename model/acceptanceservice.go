package model

import "time"

//------------ INPUT ------------//

type AcceptanceServiceInput struct {
	DocumentHAS DocumentHasInput `json:"documentHAS"`
}

type DocumentHasInput struct {
	AddressSupplier         string                       `json:"addressSupplier"`
	ReasonSocialSupplier    string                       `json:"reasonSocialSupplier"`
	ServiceAcceptanceSheets ServiceAcceptanceSheetsInput `json:"serviceAcceptanceSheets"`
	SupplierTIN             string                       `json:"supplierTIN"`
}

type ServiceAcceptanceSheetInput struct {
	AcceptanceDate    time.Time          `json:"acceptanceDate"`
	AcceptedBy        string             `json:"acceptedBy"`
	AccountingDateHAS time.Time          `json:"accountingDateHAS"`
	ApprovedBy        string             `json:"approvedBy"`
	AttachedFiles     AttachedFilesInput `json:"attachedFiles"`
	AttentionAreaCode string             `json:"attentionAreaCode"`
	BroadcastDate     time.Time          `json:"broadcastDate"`
	ClientTIN         string             `json:"clientTIN"`
	CurrencyHAS       string             `json:"currencyHAS"`
	DocumentMaterial  int64              `json:"documentMaterial"`
	FiscalYear        int                `json:"fiscalYear"`
	ItemHAS           ItemsHasInput      `json:"itemHAS"`
	NumberHAS         int                `json:"numberHAS"`
	Observation       string             `json:"observation"`
	ServiceEndDate    time.Time          `json:"serviceEndDate"`
	ServiceStartDate  time.Time          `json:"serviceStartDate"`
	ServiceType       string             `json:"serviceType"`
	Society           int                `json:"society"`
	SupplierCodeERP   int64              `json:"supplierCodeERP"`
	UserCreationHAS   string             `json:"userCreationHAS"`
}

type ServiceAcceptanceSheetsInput []ServiceAcceptanceSheetInput

type AttachedFileInput struct {
	URL         string `json:"URL"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type AttachedFilesInput []AttachedFileInput

type ItemHasInput struct {
	AttendedAmount               float64                `json:"attendedAmount"`
	CodeConditionPayment         string                 `json:"codeConditionPayment"`
	DocumentCurrency             string                 `json:"documentCurrency"`
	DocumentReference            int64                  `json:"documentReference"`
	EndDateServiceItem           string                 `json:"endDateServiceItem"`
	FinalDeliveryIndicator       string                 `json:"finalDeliveryIndicator"`
	ItemDescription              string                 `json:"itemDescription"`
	ItemNumber                   string                 `json:"itemNumber"`
	ItemNumberOC                 string                 `json:"itemNumberOC"`
	ItemQuantity                 int                    `json:"itemQuantity"`
	MaterialDocuments            MaterialDocumentsInput `json:"materialDocuments"`
	NumberOC                     int64                  `json:"numberOC"`
	ObservationItem              string                 `json:"observationItem"`
	PaymentCondition             string                 `json:"paymentCondition"`
	PriceItem                    float64                `json:"priceItem"`
	ReferenceType                int                    `json:"referenceType"`
	ServiceCode                  string                 `json:"serviceCode"`
	ServicesGroup                string                 `json:"servicesGroup"`
	SubItemHAS                   SubItemsHasInput       `json:"subItemHAS"`
	SubTotalItemCurrencyDocument float64                `json:"subTotalItemCurrencyDocument"`
	SubTotalItemLocalCurrency    float64                `json:"subTotalItemLocalCurrency"`
	UnitMeasureItem              string                 `json:"unitMeasureItem"`
}

type ItemsHasInput []ItemHasInput

type MaterialDocumentInput struct {
	Amount                 float64 `json:"amount"`
	Exercise               int     `json:"exercise"`
	MovementClass          int     `json:"movementClass"`
	NumberDocumentMaterial int64   `json:"numberDocumentMaterial"`
	OperationType          string  `json:"operationType"`
	Position               string  `json:"position"`
	SubTotal               float64 `json:"subTotal"`
}

type MaterialDocumentsInput []MaterialDocumentInput

type SubItemHasInput struct {
	SubItemDescription string  `json:"subItemDescription"`
	SubItemNumber      string  `json:"subItemNumber"`
	SubItemPrice       int     `json:"subItemPrice"`
	SubItemQuantity    float64 `json:"subItemQuantity"`
	SubItemServiceCode string  `json:"subItemServiceCode"`
	SubItemUnitMeasure string  `json:"subItemUnitMeasure"`
	SubTotalSubItem    float64 `json:"subTotalSubItem"`
}

type SubItemsHasInput []SubItemHasInput

//------------ OUTPUT ------------//

type AcceptanceServiceOutput struct {
	DocumentHAS DocumentHasOutput `json:"documentHAS"`
}

type DocumentHasOutput struct {
	AddressSupplier         string                        `json:"addressSupplier"`
	ReasonSocialSupplier    string                        `json:"reasonSocialSupplier"`
	ServiceAcceptanceSheets ServiceAcceptanceSheetsOutput `json:"serviceAcceptanceSheets"`
	SupplierTIN             string                        `json:"supplierTIN"`
}

type ServiceAcceptanceSheetOutput struct {
	AcceptanceDate    time.Time           `json:"acceptanceDate"`
	AcceptedBy        string              `json:"acceptedBy"`
	AccountingDateHAS time.Time           `json:"accountingDateHAS"`
	ApprovedBy        string              `json:"approvedBy"`
	AttachedFiles     AttachedFilesOutput `json:"attachedFiles"`
	AttentionAreaCode string              `json:"attentionAreaCode"`
	BroadcastDate     time.Time           `json:"broadcastDate"`
	ClientTIN         string              `json:"clientTIN"`
	CurrencyHAS       string              `json:"currencyHAS"`
	DocumentMaterial  int64               `json:"documentMaterial"`
	FiscalYear        int                 `json:"fiscalYear"`
	ItemHAS           ItemsHasOutput      `json:"itemHAS"`
	NumberHAS         int                 `json:"numberHAS"`
	Observation       string              `json:"observation"`
	ServiceEndDate    time.Time           `json:"serviceEndDate"`
	ServiceStartDate  time.Time           `json:"serviceStartDate"`
	ServiceType       string              `json:"serviceType"`
	Society           int                 `json:"society"`
	SupplierCodeERP   int64               `json:"supplierCodeERP"`
	UserCreationHAS   string              `json:"userCreationHAS"`
}

type ServiceAcceptanceSheetsOutput []ServiceAcceptanceSheetOutput

type AttachedFileOutput struct {
	URL         string `json:"URL"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type AttachedFilesOutput []AttachedFileOutput

type ItemHasOutput struct {
	AttendedAmount               float64                 `json:"attendedAmount"`
	CodeConditionPayment         string                  `json:"codeConditionPayment"`
	DocumentCurrency             string                  `json:"documentCurrency"`
	DocumentReference            int64                   `json:"documentReference"`
	EndDateServiceItem           string                  `json:"endDateServiceItem"`
	FinalDeliveryIndicator       string                  `json:"finalDeliveryIndicator"`
	ItemDescription              string                  `json:"itemDescription"`
	ItemNumber                   string                  `json:"itemNumber"`
	ItemNumberOC                 string                  `json:"itemNumberOC"`
	ItemQuantity                 int                     `json:"itemQuantity"`
	MaterialDocuments            MaterialDocumentsOutput `json:"materialDocuments"`
	NumberOC                     int64                   `json:"numberOC"`
	ObservationItem              string                  `json:"observationItem"`
	PaymentCondition             string                  `json:"paymentCondition"`
	PriceItem                    float64                 `json:"priceItem"`
	ReferenceType                int                     `json:"referenceType"`
	ServiceCode                  string                  `json:"serviceCode"`
	ServicesGroup                string                  `json:"servicesGroup"`
	SubItemHAS                   SubItemsHasOutput       `json:"subItemHAS"`
	SubTotalItemCurrencyDocument float64                 `json:"subTotalItemCurrencyDocument"`
	SubTotalItemLocalCurrency    float64                 `json:"subTotalItemLocalCurrency"`
	UnitMeasureItem              string                  `json:"unitMeasureItem"`
}

type ItemsHasOutput []ItemHasOutput

type MaterialDocumentOutput struct {
	Amount                 float64 `json:"amount"`
	Exercise               int     `json:"exercise"`
	MovementClass          int     `json:"movementClass"`
	NumberDocumentMaterial int64   `json:"numberDocumentMaterial"`
	OperationType          string  `json:"operationType"`
	Position               string  `json:"position"`
	SubTotal               float64 `json:"subTotal"`
}

type MaterialDocumentsOutput []MaterialDocumentOutput

type SubItemHasOutput struct {
	SubItemDescription string  `json:"subItemDescription"`
	SubItemNumber      string  `json:"subItemNumber"`
	SubItemPrice       int     `json:"subItemPrice"`
	SubItemQuantity    float64 `json:"subItemQuantity"`
	SubItemServiceCode string  `json:"subItemServiceCode"`
	SubItemUnitMeasure string  `json:"subItemUnitMeasure"`
	SubTotalSubItem    float64 `json:"subTotalSubItem"`
}

type SubItemsHasOutput []SubItemHasOutput
