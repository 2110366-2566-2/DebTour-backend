package controllers

import (
	// "fmt"

	"github.com/gin-gonic/gin"
)

// login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService LoginService
	jWtService   JWTService
}

func LoginHandler(loginService LoginService,
	jWtService JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controllers *loginController) Login(ctx *gin.Context) string {
	var credential LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	username := ctx.Param("username")
	role := ctx.Param("role")
	isUserAuthenticated := controllers.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controllers.jWtService.GenerateToken(username, role, true)

	}
	return ""
}
