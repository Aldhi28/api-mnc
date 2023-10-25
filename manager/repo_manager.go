package manager

import "login-go/repository"

type RepoManager interface {
	BillRepo() repository.BillRepository
	CustomerRepo() repository.CustomerRepository
	ProductRepo() repository.ProductRepository
	userRepo() repository.UserRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) BillRepo() repository.BillRepository {
	return repository.NewBillRepository(r.infra.Conn())
}

func (r *repoManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.Conn())
}

func (r *repoManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infra.Conn())
}

func (r *repoManager) userRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func NewRepoManager(infraParam InfraManager) RepoManager {
	return &repoManager{
		infra: infraParam,
	}
}
