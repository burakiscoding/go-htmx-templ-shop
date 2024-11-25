package stores

import (
	"errors"
	"slices"
)

type Store struct {
}

func NewStore() *Store {
	products = append(products,
		newProduct("Sugar 1kg", 50, 5),
		newProduct("Water 0.5lt", 50, 0.5),
		newProduct("Rice 1kg", 30, 4))

	return &Store{}
}

type Product struct {
	Id        int
	Name      string
	Quantity  int
	UnitPrice float32
}

var id int = 0
var products []Product

func newId() int {
	id++
	return id
}

func newProduct(name string, quantity int, unitPrice float32) Product {
	return Product{Id: newId(), Name: name, Quantity: quantity, UnitPrice: unitPrice}
}

func (s *Store) GetAll() ([]Product, error) {
	return products, nil
}

func (s *Store) GetById(id int) (Product, error) {
	idx := slices.IndexFunc(products, func(product Product) bool {
		return product.Id == id
	})
	if idx != -1 {
		return products[idx], nil
	}
	return Product{}, errors.New("not found")
}

func (s *Store) Update(id int, name string, quantity int, unitPrice float32) (Product, error) {
	idx := slices.IndexFunc(products, func(product Product) bool {
		return product.Id == id
	})
	if idx != -1 {
		products[idx].Name = name
		products[idx].Quantity = quantity
		products[idx].UnitPrice = unitPrice
		return products[idx], nil
	}
	return Product{}, errors.New("not found")
}

func (s *Store) Delete(id int) error {
	idx := slices.IndexFunc(products, func(product Product) bool {
		return product.Id == id
	})
	if idx != -1 {
		products = append(products[:idx], products[idx+1:]...)
		return nil
	}

	return errors.New("not found")
}

func (s *Store) Add(name string, quantity int, unitPrice float32) (Product, error) {
	product := newProduct(name, quantity, unitPrice)
	products = append(products, product)
	return product, nil
}
