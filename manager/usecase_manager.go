package manager

import "login-go/usecase"

type UseCaseManager interface {
	ProductUseCase() usecase.ProductUseCase
	CustomerUseCase() usecase.CustomerUseCase
	BillUseCase() usecase.BillUseCase
	UserUseCase() usecase.UserUseCase
	AuthUseCase() usecase.AuthUsecase
}

type useCaseManager struct {
	repoManager RepoManager
}

func (u *useCaseManager) ProductUseCase() usecase.ProductUseCase {
	return usecase.NewProductUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) CustomerUseCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repoManager.CustomerRepo())
}

func (u *useCaseManager) BillUseCase() usecase.BillUseCase {
	return usecase.NewBillUseCase(u.repoManager.BillRepo(), u.CustomerUseCase(), u.ProductUseCase())
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManager.userRepo())
}

func (u *useCaseManager) AuthUseCase() usecase.AuthUsecase {
	return usecase.NewAuthUseCase(u.UserUseCase())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repo,
	}
}
