package controllers

import "github.com/gin-gonic/gin"

// "fmt"

// "os/user"

// "fmt"

// "github.com/gin-gonic/gin"

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

// hash text to some random string
func hash(text string) string {
	hashed_text := ""
	for _, c := range text {
		hashed_text += string(c + 3)
	}

	return hashed_text
}

func (controllers *loginController) Login(ctx *gin.Context) string {
	// var credential LoginCredentials
	// err := ctx.ShouldBind(&credential)
	// if err != nil {
	// 	return "no data found"
	// }
	username := ctx.Param("username")
	// role := ctx.Param("role")
	// fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> username:", username)
	// isUserAuthenticated := controllers.loginService.LoginUser(username)
	// if isUserAuthenticated {
	// 	// return controllers.jWtService.GenerateToken(username, role, true)
	return controllers.jWtService.GenerateToken(username, true)

	// }
	// return ""
	// return hash(username)
}
