package main

import (
	"webserver/cmd/server/controllers"
	"webserver/internal/products"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := controllers.NewProduct(service)

	r := gin.Default()

	productGroupRouter := r.Group("/api/v1/products")
	productGroupRouter.GET("/", p.GetAll())
	productGroupRouter.GET("/:id", p.GetOne())
	productGroupRouter.POST("/", p.Store())
	productGroupRouter.DELETE("/:id", p.Delete())
	productGroupRouter.PUT("/:id", p.Update())

	r.Run()
}
