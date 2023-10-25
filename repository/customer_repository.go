package repository

import (
	"database/sql"
	"fmt"

	"login-go/model"
	"login-go/model/dto"
	"login-go/utils/common"
	"login-go/utils/constant"
)

type CustomerRepository interface {
	BaseRepository[model.Customer]
	BaseRepositoryPaging[model.Customer]
}

type customerRepository struct {
	db *sql.DB
}

func (c *customerRepository) Create(payload model.Customer) error {
	_, err := c.db.Exec(constant.CUSTOMER_INSERT, payload.Id, payload.Name, payload.PhoneNumber, payload.Address)
	if err != nil {
		return err
	}
	return nil
}

func (c *customerRepository) List() ([]model.Customer, error) {
	rows, err := c.db.Query(constant.CUSTOMER_LIST)
	if err != nil {
		return nil, err
	}
	var customers []model.Customer

	for rows.Next() {
		var customer model.Customer
		err = rows.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address)
		customers = append(customers, customer)
	}
	return nil, nil
}

func (c *customerRepository) Get(id string) (model.Customer, error) {
	var customer model.Customer
	err := c.db.QueryRow(constant.CUSTOMER_GET, id).Scan(
		&customer.Id,
		&customer.Name,
		&customer.PhoneNumber,
		&customer.Address,
	)
	if err != nil {
		return model.Customer{}, fmt.Errorf("Error Get Customer : %s ", err.Error())
	}
	return customer, nil
}

func (c *customerRepository) Update(payload model.Customer) error {
	_, err := c.db.Exec(constant.CUSTOMER_UPDATE, payload.Name, payload.PhoneNumber, payload.Address, payload.Id)
	if err != nil {
		return fmt.Errorf(" Error Update Customer : %s ", err.Error())
	}
	return nil
}

func (c *customerRepository) Delete(id string) error {
	_, err := c.db.Exec(constant.CUSTOMER_DELETE, id)
	if err != nil {
		return fmt.Errorf("repo : Error Delete Customer : %s ", err.Error())
	}
	return nil
}

func (c *customerRepository) Paging(requestPaging dto.PaginationParam, query ...string) ([]model.Customer, dto.Paging, error) {
	var paginationQuery dto.PaginationQuery
	paginationQuery = common.GetPaginationParams(requestPaging)
	querySelect := "SELECT id, name, phone_number, address FROM customer LIMIT $1 OFFSET $2"
	rows, err := c.db.Query(querySelect, paginationQuery.Take, paginationQuery.Skip)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.PhoneNumber, &customer.Address)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		customers = append(customers, customer)
	}

	// count total rows
	var totalRows int
	row := c.db.QueryRow("SELECT COUNT(*) FROM customer")
	err = row.Scan(&totalRows)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	return customers, common.Paginate(paginationQuery.Page, paginationQuery.Take, totalRows), nil
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

