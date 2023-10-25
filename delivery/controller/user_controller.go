package controller

import (
	"login-go/model"
	"login-go/usecase"
	"login-go/utils/common"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	router *gin.Engine
	userUc usecase.UserUseCase
}

func (u *UserController) createHandler(c *gin.Context) {
	var user model.UserCredential
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	user.Id = common.GenerateUUID()
	findUser, err := u.userUc.RegisterNewUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	userResponse := map[string]any{
		"id":       findUser.Id,
		"username": findUser.UserName,
		"isActive": findUser.IsActive,
	}

	c.JSON(http.StatusOK, userResponse)
}

func NewUserController(r *gin.Engine, userUseCase usecase.UserUseCase) {
	controller := UserController{
		router: r,
		userUc: userUseCase,
	}

	routerGroup := r.Group("/api/v1")
	routerGroup.POST("/users", controller.createHandler)
}
