package model

type AcceptanceServiceInput struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type AcceptanceServicesInput []AcceptanceServiceInput

type AcceptanceServiceOutput struct {
	Code string `json:"codigo"`
	Name string `json:"nombre"`
}

type AcceptanceServicesOutput []AcceptanceServiceOutput
