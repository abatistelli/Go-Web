package main

import (
	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id              int     `json:"id"`
	Nombre          string  `json:"nombre" binding:"required"`
	Color           string  `json:"color" binding:"required"`
	Precio          float64 `json:"precio" binding:"required"`
	Stock           int     `json:"stock"  binding:"required"`
	Codigo          string  `json:"codigo" binding:"required"`
	Publicado       bool    `json:"publicado"`
	FechaDeCreacion string  `json:"fechaDeCreacion" binding:"required"`
}

var productos []Producto

func Guardar(id *int) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token != "0000" {
			c.JSON(401, gin.H{
				"error": "no tiene permisos para realizar la petición solicitada",
			})
			return
		}

		var req Producto

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		req.Id = *id
		*id++

		productos = append(productos, req)
		c.JSON(200, productos)
	}
}

func main() {
	r := gin.Default()
	lastID := 0
	var id *int
	id = &lastID

	r.POST("/productos", Guardar(id))

	r.Run()
}
