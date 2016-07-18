package middlewares

// Package middlewares contains gin middlewares
// Usage: router.Use(middlewares.Connect)

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	needLoginURLs = []string{
		"/home/",
		"/customers/",
		"/orders/",
		"/admin/",
		"/me/",
	}
)

func inLoginUrls(path string) bool {
	for _, p := range needLoginURLs {
		if p == path {
			return true
		}
	}
	return false
}

/*AuthMiddleware should redirect user to login page if not logged in*/
func AuthMiddleware(c *gin.Context) {
	fmt.Println("Request Url path: " + c.Request.URL.Path)
	// if inLoginUrls(c.Request.URL.Path) {
	if true {
		if isLoggedIn(c) {
			c.Next()
		} else {
			c.Redirect(302, "/login")
			c.Abort()
		}
	}
	if c.Request.URL.Path == "/login" || c.Request.URL.Path == "/" {
		if isLoggedIn(c) {
			c.Redirect(302, "/home")
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func isLoggedIn(c *gin.Context) bool {
	// rdb := db.ReactSession
	// _, found, err := auth.AuthSession(c, rdb)
	// if !found {
	//   return false
	// } else if err != nil {
	//   panic("wat")
	// } else {
	//   return true
	// }
	return false
}

// ErrorHandler is a middleware to handle errors encountered during requests
func ErrorHandler(c *gin.Context) {
	c.Next()

	// TODO: Handle it in a better way
	if len(c.Errors) > 0 {
		c.HTML(http.StatusBadRequest, "400", gin.H{
			"errors": c.Errors,
		})
	}
}

/*CORS adds a Cords header to the request so that it will get accepted*/
func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	c.Next()
}
