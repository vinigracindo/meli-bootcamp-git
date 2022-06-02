package controllers

import (
	"net/http"
	"strconv"
	"webserver/internal/products"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service products.Service
}

func NewProduct(p products.Service) *ProductController {
	return &ProductController{
		service: p,
	}
}

func (c *ProductController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}
func (c *ProductController) GetOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		p, err := c.service.GetOne(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *ProductController) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := struct {
			Name      string  `json:"name"`
			Color     string  `json:"color"`
			Price     float64 `json:"price"`
			Stock     int     `json:"stock"`
			Code      string  `json:"code"`
			Published bool    `json:"published"`
		}{}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		p, err := c.service.Save(request.Name, request.Color, request.Price, request.Stock, request.Code, request.Published)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *ProductController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		request := struct {
			Name      string  `json:"name"`
			Color     string  `json:"color"`
			Price     float64 `json:"price"`
			Stock     int     `json:"stock"`
			Code      string  `json:"code"`
			Published bool    `json:"published"`
		}{}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		p, err := c.service.Update(id, request.Name, request.Color, request.Price, request.Stock, request.Code, request.Published)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *ProductController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		err = c.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
	}
}
