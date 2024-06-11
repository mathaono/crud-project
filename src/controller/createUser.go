package controller

import (
	"github.com/gin-gonic/gin"
	errormodels "github.com/mathaono/crud-project/src/configuration/errorModels"
)

func CreateUser(c *gin.Context) {

	err := errormodels.BadRequestError("Valor inválido")
	c.JSON(int(err.Code), err)
}
