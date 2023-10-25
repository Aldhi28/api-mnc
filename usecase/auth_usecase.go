package usecase

import (
	"fmt"

	"login-go/utils/security"
)

type AuthUsecase interface {
	Login(username string, password string) (string, error)
}

type authUsecase struct {
	userUc UserUseCase
}

func (a *authUsecase) Login(username string, password string) (string, error) {
	user, err := a.userUc.FindByUsernamePassword(username, password)

	if err != nil {
		return "", fmt.Errorf("invalid username & password")
	}

	// setelah login berhasil, maka kita berikan token
	token, err := security.CreateAccessToken(user)
	if err != nil {
		return "", fmt.Errorf("Failed to Generate Token : %s", err.Error())
	}
	return token, nil
}

func NewAuthUseCase(userUseCase UserUseCase) AuthUsecase {
	return &authUsecase{
		userUc: userUseCase,
	}
}
