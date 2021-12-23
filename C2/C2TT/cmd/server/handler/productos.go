package handler

import (
	"github.com/abatistelli/go-web/C2/C2TT/internal/productos"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre          string  `json:"nombre" binding:"required"`
	Color           string  `json:"color" binding:"required"`
	Precio          float64 `json:"precio" binding:"required"`
	Stock           int     `json:"stock"  binding:"required"`
	Codigo          string  `json:"codigo" binding:"required"`
	Publicado       bool    `json:"publicado"`
	FechaDeCreacion string  `json:"fechaDeCreacion" binding:"required"`
}

type Product struct {
	service productos.Service
}

func NewProduct(p productos.Service) *Product {
	return &Product{
		service: p,
	}
}

func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "0000" {
			ctx.JSON(401, gin.H{
				"error": "token invaldio",
			})
			return
		}

		var req request

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		p, err := c.service.Store(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaDeCreacion)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, p)
		return
	}
}

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "0000" {
			ctx.JSON(401, gin.H{
				"error": "token invaldio",
			})
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, p)
		return
	}
}
