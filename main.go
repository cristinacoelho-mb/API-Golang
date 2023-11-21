package main

import (
	"API/controllers"
	"API/database"
	"API/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

// @externalDocs.description  OpenAPI
// in hearder
// @name Authorization

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()

	// Swageer 2.0 Meta Information

	docs.SwaggerInfo.Title = "Example API Bank"
	docs.SwaggerInfo.Description = "This is a sample API Bank server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Welcome To Sample program swagger"})
	})

	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/account")
		{
			accounts.POST("/create", controllers.NewAccount)
			accounts.PATCH("/update/:id", controllers.EditaAccount)
			accounts.DELETE("/delete/:id", controllers.DeletaAccount)

		}
	}
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
