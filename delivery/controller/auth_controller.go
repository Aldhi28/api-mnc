package controller

import (
	"login-go/model"
	"login-go/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	router *gin.Engine
	authUc usecase.AuthUsecase
}

func (a *AuthController) createHandler(c *gin.Context) {
	var payload model.UserCredential
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	token, err := a.authUc.Login(payload.UserName, payload.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success Login",
		"token":   token,
	})
}

func NewAuthController(r *gin.Engine, authUsecase usecase.AuthUsecase) {
	ctr := &AuthController{
		router: r,
		authUc: authUsecase,
	}

	routerGroup := r.Group("/api/v1")
	routerGroup.POST("/login", ctr.createHandler)
}
