package controller

import (
	"net/http"
	"strconv"

	"login-go/model"
	"login-go/model/dto"
	"login-go/usecase"
	"login-go/utils/common"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	router    *gin.Engine
	productUc usecase.ProductUseCase
}

func (p *ProductController) createHandler(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	product.Id = common.GenerateUUID()
	err := p.productUc.RegisterNewProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Success Create New Product",
		"data":    product,
	})
}

func (p *ProductController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	name := c.Query("name")
	paginationParam := dto.PaginationParam{
		Page:  page,
		Limit: limit,
	}
	products, paging, err := p.productUc.FindAllProduct(paginationParam, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get All Data Successfully",
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data":   products,
		"paging": paging,
	})
}

func (p *ProductController) getHandler(c *gin.Context) {
	id := c.Param("id")
	product, err := p.productUc.FindProductById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success Get Employee by Id",
		"data":    product,
	})
}

func (e *ProductController) updateHandler(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := e.productUc.UpdateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Success Updated Employee",
		"data":    product,
	})
}

func (p *ProductController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := p.productUc.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success Delete",
	})
}

func NewProductController(router *gin.Engine, productUseCase usecase.ProductUseCase) {
	ctr := &ProductController{
		router:    router,
		productUc: productUseCase,
	}

	routerGroup := ctr.router.Group("/api/v1")
	routerGroup.POST("/product", ctr.createHandler)
	routerGroup.GET("/product", ctr.listHandler)
	routerGroup.GET("/product/:id", ctr.getHandler)
	routerGroup.PUT("/product", ctr.updateHandler)
	routerGroup.DELETE("/product/:id", ctr.deleteHandler)
}
