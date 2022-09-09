package handler

import (
	"pass-trougth/infraestructure/handler/acceptanceservice"
	"pass-trougth/model"
)

func InitRoutes(specification model.RouterSpecification) {
	// A
	acceptanceservice.NewRouter(specification)
}
