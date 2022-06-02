package products

type Service interface {
	GetAll() ([]Product, error)
	GetOne(id int64) (Product, error)
	Save(name, color string, price float64, stock int, code string, published bool) (Product, error)
	Update(id int64, name, color string, price float64, stock int, code string, published bool) (Product, error)
	Delete(id int64) error
}

type service struct {
	repository Repository
}

func (s service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s service) GetOne(id int64) (Product, error) {
	p, err := s.repository.GetOne(id)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func (s service) Save(name, color string, price float64, stock int, code string, published bool) (Product, error) {
	p, err := s.repository.Save(name, color, price, stock, code, published)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func (s service) Update(id int64, name, color string, price float64, stock int, code string, published bool) (Product, error) {
	p, err := s.repository.Update(id, name, color, price, stock, code, published)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func (s service) Delete(id int64) error {
	err := s.repository.Delete(id)
	return err
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
