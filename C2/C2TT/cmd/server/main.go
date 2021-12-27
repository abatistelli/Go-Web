package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abatistelli/go-web/C2/C2TT/cmd/server/handler"
	"github.com/abatistelli/go-web/C2/C2TT/internal/productos"
	"github.com/abatistelli/go-web/C2/C2TT/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archio .env")
	}

	token := os.Getenv("TOKEN")

	fmt.Println("TOKEN: ", token)

	db := store.New(store.FileType, "products.json")
	repo := productos.NewRepository(db)
	service := productos.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")

	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())
	pr.PATCH("/:id", p.UpdateNameOrPrice())
	r.Run()
}
