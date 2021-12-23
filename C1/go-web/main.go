/*

Ejercicio 2 - Hola{nombre}

1) Crea dentro de la carpeta go-web un archivo llamado main.go

2) Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y
diga Hola seguido por tu nombre.

3)Pegale al endpoint para corroborar que la respuesta sea la correcta.


Ejercicio 3 - Listar Entidad
Ya habiendo creado y probado nuestra API que nos saluda,
generamos una ruta que devuelve un listado de la temática elegida.

1) Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
2) Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
3)Genera un handler para el endpoint llamado “GetAll”.
4)Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.

*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id              int     `json:"id"`
	Nombre          string  `json:"nombre"`
	Color           string  `json:"color"`
	Precio          float64 `json:"precio"`
	Stock           int     `json:"stock"`
	Codigo          string  `json:"codigo"`
	Publicado       bool    `json:"publicado"`
	FechaDeCreacion string  `json:"fechaDeCreacion"`
}

func GetAll(c *gin.Context) {

	productosJSON, errArchivo := os.ReadFile("./productos.json")
	if errArchivo != nil {
		panic("Archivo no encontrado o dañado")
	}

	var p []Producto

	if errUnmarshal := json.Unmarshal(productosJSON, &p); errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}

	id := c.Query("id")
	nombre := c.Query("nombre")
	color := c.Query("color")
	precio := c.Query("precio")
	stock := c.Query("stock")
	publicado := c.Query("publicado")
	fechaDeCreacion := c.Query("fechaDeCreacion")

	var sliceFiltrado []Producto

	c.JSON(200, p)

}

func filtrarSlice()

func main() {

	///// EJERCICIO 2 TM
	r := gin.Default()
	r.GET("/welcome/:name", func(c *gin.Context) {
		nombre := c.Param("name")
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hola %s", nombre),
		})
	})
	/////////////////////////////////

	//// Ejercicio 3 TM
	r.GET("/productos", GetAll)

	/// Ejercicio 1 TT
	r.Run()
}
