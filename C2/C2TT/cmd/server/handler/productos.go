package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/abatistelli/go-web/C2/C2TT/internal/productos"
	"github.com/abatistelli/go-web/C2/C2TT/pkg/web"
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

// AlmacenarProductos godoc
// @Summary Almacenar Productos
// @Tags Products, Almacenar, Guardar
// @Description almacena un producto ingresado, en el archivo .json
// @Accept	json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]

func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "TOKEN INVALIDO"))
			return
		}

		var req request

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		p, err := c.service.Store(req.Nombre, req.Color, req.Precio, req.Stock, req.Codigo, req.Publicado, req.FechaDeCreacion)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, p, ""))
		return
	}
}

// Lista de Productos godoc
// @Summary ListaDeProductos
// @Tags Products, List
// @Description obtiene la lista de los productos almacenados en un .json
// @Accept	json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "TOKEN INVALIDO"))
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, p, ""))
		return
	}
}

// Actualizar Producto godoc
// @Summary Actualiza Producto
// @Tags Products, Update
// @Description actualiza un producto por medio de su id
// @Accept	json
// @Produce json
// @Param token header string true "token"
// @Param id of product
// @Success 200 {object} web.Response
// @Router /products/:id [put]

func (c *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "TOKEN INVALIDO"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		if req.Nombre == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}
		if req.Color == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El color del producto es requerido"))
			return
		}
		if req.Precio == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio del producto es requerido"))
			return
		}
		if req.Stock == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El stock del producto es requerido"))
			return
		}
		if req.Codigo == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El codigo del producto es requerido"))
			return
		}
		if req.FechaDeCreacion == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "La fecha de creacion del producto es requerido"))
			return
		}

		p, err := c.service.Update(int(id), req.Nombre, req.Color, req.Precio, req.Stock, req.Color, req.Publicado, req.FechaDeCreacion)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// Eliminar Producto godoc
// @Summary Eliminar Producto
// @Tags Products, Delete
// @Description elimina el producto segun su id
// @Accept	json
// @Produce json
// @Param token header string true "token"
// @Param id of product
// @Success 200 {object} web.Response
// @Router /products/:id [delete]

func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "TOKEN INVALIDO"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "invalid ID"))
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("El producto %d ha sido eliminado", id), ""))
	}
}

// Acutalizar Nombre o Precio de Productos godoc
// @Summary Actualizar Nombre o Precio del Producto
// @Tags Products, UpdateNameOrPrice
// @Description actualiza le nombre, precio o ambos de un producto segun su id
// @Accept	json
// @Produce json
// @Param token header string true "token"
// @Param id of product
// @Success 200 {object} web.Response
// @Router /products/:id [patch]

func (c *Product) UpdateNameOrPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "TOKEN INVALIDO"))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "invalid ID"))
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		if req.Nombre == "" && req.Precio == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}

		if req.Precio == 0 && req.Nombre == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio del producto es requerido"))
			return
		}

		p, err := c.service.UpdateNameOrPrice(int(id), req.Nombre, req.Precio)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}
