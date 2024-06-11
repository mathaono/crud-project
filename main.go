package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mathaono/crud-project/src/controller/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Inicialização do roteador (router)
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	//Rodando a aplicação na porta 8080. http://localhost:8080/
	//Criando e validando a variável err que armazena o retorno do tipo "error" da função Run()
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
