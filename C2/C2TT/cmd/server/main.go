package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abatistelli/go-web/C2/C2TT/cmd/server/handler"
	"github.com/abatistelli/go-web/C2/C2TT/docs"
	"github.com/abatistelli/go-web/C2/C2TT/internal/productos"
	"github.com/abatistelli/go-web/C2/C2TT/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwangger "github.com/swaggo/gin-swagger"
)

// @title ListaDeProductos API
// @version 1.0
// @description Esta API permite manipular una lista de productos
//
// @contact.name Products Support
// @contact.url http://www.productos.com
//

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

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwangger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/products")

	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())
	pr.PATCH("/:id", p.UpdateNameOrPrice())
	r.Run()
}
