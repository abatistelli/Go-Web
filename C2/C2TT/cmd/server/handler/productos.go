package handler

import (
	"fmt"
	"strconv"

	"github.com/abatistelli/go-web/C2/C2TT/internal/productos"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre          string  `json:"nombre"`
	Color           string  `json:"color"`
	Precio          float64 `json:"precio"`
	Stock           int     `json:"stock" `
	Codigo          string  `json:"codigo"`
	Publicado       bool    `json:"publicado"`
	FechaDeCreacion string  `json:"fechaDeCreacion"`
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

func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "0000" {
			ctx.JSON(401, gin.H{
				"error": "token invaldio",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
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

		if req.Nombre == "" {
			ctx.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			return
		}
		if req.Color == "" {
			ctx.JSON(400, gin.H{"error": "El color del producto es requerido"})
			return
		}
		if req.Precio == 0 {
			ctx.JSON(400, gin.H{"error": "El precio del producto es requerido"})
			return
		}
		if req.Stock == 0 {
			ctx.JSON(400, gin.H{"error": "El stock del producto es requerido"})
			return
		}
		if req.Codigo == "" {
			ctx.JSON(400, gin.H{"error": "El codigo del producto es requerido"})
			return
		}
		if req.FechaDeCreacion == "" {
			ctx.JSON(400, gin.H{"error": "La fecha de creacion del producto es requerido"})
			return
		}

		p, err := c.service.Update(int(id), req.Nombre, req.Color, req.Precio, req.Stock, req.Color, req.Publicado, req.FechaDeCreacion)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "0000" {
			ctx.JSON(401, gin.H{
				"error": "token invaldio",
			})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, gin.H{"error": "invalid ID"})
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"data": fmt.Sprintf("El producto %d ha sido eliminado", id)})
	}
}

func (c *Product) UpdateNameOrPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "0000" {
			ctx.JSON(401, gin.H{
				"error": "token invaldio",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(404, gin.H{"error": "invalid ID"})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}

		if req.Nombre == "" && req.Precio == 0 {
			ctx.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			return
		}

		if req.Precio == 0 && req.Nombre == "" {
			ctx.JSON(400, gin.H{"error": "El precio del producto es requerido"})
			return
		}

		p, err := c.service.UpdateNameOrPrice(int(id), req.Nombre, req.Precio)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}
