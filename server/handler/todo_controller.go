package handler

import (
	"golang-tutorial/contract"
	"golang-tutorial/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	service contract.TodoService
}

func (c *TodoController) getPrefix() string {
	return "/todo"
}

func (c *TodoController) initService(service *contract.Service) {
	c.service = service.Todo
}

func (c *TodoController) initRoute(app *gin.RouterGroup) {
	app.GET("/:id", c.GetTodo)
	app.POST("/create", c.CreateTodo)
	app.PUT("/:id", c.UpdateTodo)
	app.PATCH("/:id", c.UpdateTodo)
	app.DELETE("/:id", c.DeleteTodo)
}

func (c *TodoController) GetTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	response, err := c.service.GetTodo(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *TodoController) CreateTodo(ctx *gin.Context) {
	var payload dto.TodoRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.service.CreateTodo(&payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *TodoController) UpdateTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var payload dto.TodoRequest
	err = ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.service.UpdateTodo(intID, &payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)

}

func (c *TodoController) DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	response, err := c.service.DeleteTodo(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
