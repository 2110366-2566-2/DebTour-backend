package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

// func RegisterRoutes() http.Handler {
// 	// Define the routes here
// 	//create new router using gin
// 	router := gin.Default()
// 	router.GET("/auth/:provider/callback", GetAuthCallBackFunction)
// 	return router
// }

// func (s *Server) getAuthCallBackFunction(w http.ResponseWriter, r *http.Request) {
// 	// Do this in gin framework
// 	provider := chi.URLParam(r, "provider")

// 	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
// 	//provider := r.URL.Query().Get("provider")
// 	// Get the URL of the user's profile
// 	user, err := gothic.CompleteUserAuth(w, r)
// 	if err != nil {
// 		fmt.Fprintln(w, r)
// 		return
// 	}

// 	fmt.Println(user)

// }

func GetAuthCallBackFunction(c *gin.Context) {
	//Get the value of the URL parameter "provider"
	provider := c.Param("provider")

	//Clone the request with a new context
	r := c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
	// Get the URL of the user's profile
	user, err := gothic.CompleteUserAuth(c.Writer, r)
	if err != nil {
		fmt.Fprintln(c.Writer, err)
		return
	}
	fmt.Println(user)

	http.Redirect(c.Writer, r, "http://localhost:9000", http.StatusFound)
}
