package dataresponse

import "pass-trougth/model"

func GetAcceptedDefault() model.Accepted {
	return model.Accepted{
		Description: "Petición en proceso",
		Message:     "Aceptado",
	}
}
