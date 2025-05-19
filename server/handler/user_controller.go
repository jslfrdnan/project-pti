package handler

import (
	"golang-tutorial/contract"
	"golang-tutorial/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service contract.UserService
}

func (u *UserController) getPrefix() string {
	return "/user"
}

func (u *UserController) initService(service *contract.Service) {
	u.service = service.User
}

func (u *UserController) initRoute(app *gin.RouterGroup) {
	app.GET("/user/:id", u.GetUser)
	app.POST("/register", u.Register)
	app.POST("/login", u.Login)
}

func (u *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	response, err := u.service.GetUser(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (u *UserController) Register(ctx *gin.Context) {
	var payload dto.UserRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := u.service.Register(&payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (u *UserController) Login(ctx *gin.Context) {
	var payload dto.UserRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := u.service.Login(&payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(response.StatusCode, response)
}
