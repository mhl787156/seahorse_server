// Package middlewares contains gin middlewares
// Usage: router.Use(middlewares.Connect)
package middlewares

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	fmt.Println(c.Request.URL.Path)
	if c.Request.URL.Path == "/dashboard/" || c.Request.URL.Path == "/studio/" || c.Request.URL.Path == "/profile/" {
		if isLoggedIn(c) {
			c.Next()
		} else {
			c.Redirect(302, "/login")
			c.Abort()
		}
	}
	if c.Request.URL.Path == "/login" || c.Request.URL.Path == "/" {
		if isLoggedIn(c) {
			c.Redirect(302, "/dashboard")
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

func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}
