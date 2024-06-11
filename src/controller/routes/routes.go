package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mathaono/crud-project/src/controller"
)

// Criação das rotas da aplicação
// É inicializado na main.go
// O RouterGroup é passado como parâmetro no InitRoutes pq ele vai atrelar as rotas da aplicação ao router group criado abaixo
func InitRoutes(r *gin.RouterGroup) {
	// Criação do router group
	r.GET("/findUserById/:userId", controller.FindUserByID)
	r.GET("/findUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
