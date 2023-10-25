package usecase

import (
	"fmt"
	"login-go/model"
	"login-go/model/dto"
	"login-go/repository"
)

type ProductUseCase interface {
	RegisterNewProduct(payload model.Product) error
	FindAllProductList() ([]model.Product, error)
	FindProductById(id string) (model.Product, error)
	UpdateProduct(payload model.Product) error
	DeleteProduct(id string) error
	FindAllProduct(requesPaging dto.PaginationParam, byNameEmpl string) ([]model.Product, dto.Paging, error)
}

type productUseCase struct {
	repo repository.ProductRepository
}

func (p *productUseCase) RegisterNewProduct(payload model.Product) error {
	if payload.Name == "" || payload.Price == 0 || payload.Uom == "" {
		return fmt.Errorf("Name, price, Uom is required")
	}
	err := p.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("Failed to create product : %s", err.Error())
	}
	return nil
}

func (p *productUseCase) FindAllProductList() ([]model.Product, error) {
	return p.repo.List()
}

func (p *productUseCase) FindProductById(id string) (model.Product, error) {
	return p.repo.Get(id)
}

func (p *productUseCase) FindAllProduct(requestPaging dto.PaginationParam, byNameEmpl string) ([]model.Product, dto.Paging, error) {
	return p.repo.Paging(requestPaging, byNameEmpl)
}

func (p *productUseCase) UpdateProduct(payload model.Product) error {
	_, err := p.FindProductById(payload.Id)
	if err != nil {
		return err
	}
	return p.repo.Update(payload)
}

func (p *productUseCase) DeleteProduct(id string) error {
	_, err := p.FindProductById(id)
	if err != nil {
		return err
	}
	return p.repo.Delete(id)
}

func NewProductUseCase(repository repository.ProductRepository) ProductUseCase {
	return &productUseCase{
		repo: repository,
	}
}
