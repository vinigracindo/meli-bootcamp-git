package products

import (
	"fmt"
	"time"
)

var productsInMemory []Product = []Product{}

type Repository interface {
	GetAll() ([]Product, error)
	GetOne(id int64) (Product, error)
	Save(name, color string, price float64, stock int, code string, published bool) (Product, error)
	Update(id int64, name, color string, price float64, stock int, code string, published bool) (Product, error)
	Delete(id int64) error
}

type repository struct {
	products []Product
}

func (r *repository) getLastID() int64 {
	lenOfProducts := len(r.products)
	if lenOfProducts != 0 {
		return r.products[lenOfProducts-1].ID
	} else {
		return 0
	}
}

func (r *repository) GetAll() ([]Product, error) {
	return r.products, nil
}

func (r *repository) GetOne(id int64) (Product, error) {
	for _, product := range r.products {
		if product.ID == id {
			return product, nil
		}
	}
	return Product{}, fmt.Errorf("product with id %d not found", id)
}

func (r *repository) Save(name, color string, price float64, stock int, code string, published bool) (Product, error) {
	p := Product{
		Name:      name,
		Color:     color,
		Price:     price,
		Stock:     stock,
		Code:      code,
		Published: published,
	}
	p.CreatedAt = time.Now()
	p.ID = r.getLastID() + 1
	r.products = append(r.products, p)
	return p, nil
}

func (r *repository) Update(id int64, name, color string, price float64, stock int, code string, published bool) (Product, error) {
	p := Product{Name: name, Color: color, Price: price, Stock: stock, Code: code, Published: published}
	updated := false
	for index, product := range r.products {
		if product.ID == id {
			p.ID = id
			p.CreatedAt = product.CreatedAt
			r.products[index] = p
			updated = true
			break
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("produto %d não encontrado", id)
	}
	return p, nil
}

func (r *repository) Delete(id int64) error {
	deleted := false
	for index, product := range r.products {
		if product.ID == id {
			r.products = append(r.products[:index], r.products[index+1:]...)
			deleted = true
			break
		}
	}
	if !deleted {
		return fmt.Errorf("produto %d nao encontrado", id)
	}
	return nil
}

func NewRepository() Repository {
	return &repository{
		products: productsInMemory,
	}
}
