package productos

import (
	"fmt"

	"github.com/abatistelli/go-web/C2/C2TT/pkg/store"
)

type Producto struct {
	Id              int     `json:"id"`
	Nombre          string  `json:"nombre"`
	Color           string  `json:"color"`
	Precio          float64 `json:"precio"`
	Stock           int     `json:"stock" `
	Codigo          string  `json:"codigo"`
	Publicado       bool    `json:"publicado"`
	FechaDeCreacion string  `json:"fechaDeCreacion"`
}

var ps []Producto

//var lastID int

type Repository interface {
	GetAll() ([]Producto, error)
	Store(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Producto, error)
	LastID() (int, error)
	Update(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Producto, error)
	Delete(id int) error
	UpdateNameOrPrice(id int, name string, price float64) (Producto, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Producto, error) {
	var ps []Producto
	r.db.Read(&ps)
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	var ps []Producto
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}

	if len(ps) == 0 {
		return 0, nil
	}
	return ps[len(ps)-1].Id, nil
}

func (r *repository) Store(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Producto, error) {
	//p := Producto{id, nombre, color, precio, stock, codigo, publicado, fechaDeCreacion}
	//ps = append(ps, p)
	//lastID = p.Id
	//return p, nil

	var ps []Producto
	r.db.Read(&ps)
	p := Producto{id, nombre, color, precio, stock, codigo, publicado, fechaDeCreacion}
	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return Producto{}, err
	}

	return p, nil
}

func (r *repository) Update(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Producto, error) {
	p := Producto{Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, FechaDeCreacion: fechaDeCreacion}
	updated := false

	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			ps[i] = p
			updated = true
		}
	}

	if !updated {
		return Producto{}, fmt.Errorf("Producto %d no encontrado", id)
	}

	return p, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range ps {
		if ps[i].Id == id {
			index = 1
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("Producto %d no encontrado", id)
	}
	ps = append(ps[:index], ps[index+1:]...)
	return nil
}

func (r *repository) UpdateNameOrPrice(id int, name string, price float64) (Producto, error) {
	var p Producto
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			if name != "" {
				ps[i].Nombre = name
			}
			if price != 0 {
				ps[i].Precio = price
			}
			p = ps[i]
			updated = true
		}
	}

	if !updated {
		return Producto{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	return p, nil
}
