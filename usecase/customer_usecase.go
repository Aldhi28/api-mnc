package usecase

import (
	"fmt"
	"login-go/model"
	"login-go/model/dto"
	"login-go/repository"
)

type CustomerUseCase interface {
	RegisterNewCustomer(payload model.Customer) error
	FindAllCustomerList() ([]model.Customer, error)
	FindCustomerById(id string) (model.Customer, error)
	UpdateCustomer(payload model.Customer) error
	DeleteCustomer(id string) error
	FindAllCustomer(requesPaging dto.PaginationParam, byNameEmpl string) ([]model.Customer, dto.Paging, error)
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

func (c *customerUseCase) RegisterNewCustomer(payload model.Customer) error {
	if payload.Name == "" || payload.PhoneNumber == "" || payload.Address == "" {
		return fmt.Errorf("Name, Phone Number, Address is required")
	}
	err := c.repo.Create(payload)
	if err != nil {
		return fmt.Errorf("Failed to create Customer : %s", err.Error())
	}
	return nil
}

func (c *customerUseCase) FindAllCustomerList() ([]model.Customer, error) {
	return c.repo.List()
}

func (c *customerUseCase) FindCustomerById(id string) (model.Customer, error) {
	return c.repo.Get(id)
}

func (c *customerUseCase) FindAllCustomer(requestPaging dto.PaginationParam, byNameEmpl string) ([]model.Customer, dto.Paging, error) {
	return c.repo.Paging(requestPaging, byNameEmpl)
}

func (c *customerUseCase) UpdateCustomer(payload model.Customer) error {
	_, err := c.FindCustomerById(payload.Id)
	if err != nil {
		return err
	}
	return c.repo.Update(payload)
}

func (c *customerUseCase) DeleteCustomer(id string) error {
	_, err := c.FindCustomerById(id)
	if err != nil {
		return err
	}
	return c.repo.Delete(id)
}


func NewCustomerUseCase(repository repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		repo: repository,
	}
}
