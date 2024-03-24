package controllers

import "github.com/gin-gonic/gin"

// LoginController interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService LoginService
	jwtService   JWTService
}

func LoginHandler(loginService LoginService,
	jWtService JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jWtService,
	}
}

func (controllers *loginController) Login(ctx *gin.Context) string {
	username := ctx.Param("username")

	return controllers.jwtService.GenerateToken(username, true)
}
