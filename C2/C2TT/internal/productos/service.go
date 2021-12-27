package productos

type Service interface {
	GetAll() ([]Producto, error)
	Store(nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Producto, error)
	Update(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Producto, error)
	Delete(id int) error
	UpdateNameOrPrice(id int, name string, price float64) (Producto, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Producto, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) Store(nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Producto, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Producto{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, nombre, color, precio, stock, codigo, publicado, fechaDeCreacion)
	if err != nil {
		return Producto{}, err
	}

	return producto, nil
}

func (s *service) Update(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaDeCreacion string) (Producto, error) {
	return s.repository.Update(id, nombre, color, precio, stock, codigo, publicado, fechaDeCreacion)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateNameOrPrice(id int, name string, price float64) (Producto, error) {
	return s.repository.UpdateNameOrPrice(id, name, price)
}
