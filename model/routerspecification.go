package model

import (
	"github.com/gin-gonic/gin"
)

type RouterSpecification struct {
	Api    *gin.Engine
	Logger Logger
	Config ServiceConfig
}
