package controllers

import (
	"flag"
	"fmt"

	// "io/ioutil"
	"net/http"
	"os"
	"path"

	// "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zalando/gin-oauth2/google"
	goauth "google.golang.org/api/oauth2/v2"
)

var redirectURL, credFile string // new
var secret []byte
var sessionName string
var stateKey string
var sessionID string
var scopes []string

func InitOauth() string {
	bin := path.Base(os.Args[0])
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage of %s
================
`, bin)
		flag.PrintDefaults()
	}
	flag.StringVar(&redirectURL, "redirect", "http://localhost:9000/api/v1/auth/", "URL to be redirected to after authorization.")
	flag.StringVar(&credFile, "cred-file", "./test-clientid.google.json", "Credential JSON file")
	scopes = []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
		// You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
	}
	flag.Parse()
	stateKey = "state"
	sessionID = "ginoauth_google_session"
	secret = []byte("secret")
	sessionName = "goquestsession"
	google.Setup(redirectURL, credFile, scopes, secret)
	return sessionName
}

func UserInfoHandler(ctx *gin.Context) { // new
	var (
		res goauth.Userinfo
		ok  bool
	)

	val := ctx.MustGet("user")
	if res, ok = val.(goauth.Userinfo); !ok {
		res = goauth.Userinfo{Name: "no user"}
	}

	// ctx.JSON(http.StatusOK, gin.H{"Hello": "from private", "user": res.Email})

	output := make(map[string]interface{})
	output["username"] = res.Id
	output["email"] = res.Email
	output["firstname"] = res.GivenName
	output["lastname"] = res.FamilyName
	output["picture"] = res.Picture

	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": output})
}
