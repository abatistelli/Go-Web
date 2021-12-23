package main

import (
	"github.com/abatistelli/go-web/C2/C2TT/cmd/server/handler"
	"github.com/abatistelli/go-web/C2/C2TT/internal/productos"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := productos.NewRepository()
	service := productos.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")

	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())

	r.Run()
}
