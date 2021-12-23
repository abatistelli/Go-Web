package productos

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

var ps []Producto
var lastID int

type Repository interface {
	GetAll() ([]Producto, error)
	Store(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Producto, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Producto, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Producto, error) {
	p := Producto{id, nombre, color, precio, stock, codigo, publicado, fechaDeCreacion}
	ps = append(ps, p)
	lastID = p.Id
	return p, nil
}
