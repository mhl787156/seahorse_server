package middlewares

// Package middlewares contains gin middlewares
// Usage: router.Use(middlewares.Connect)

import (
	"net/http"
	"strings"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mhl787156/seahorse_server/auth"
	"github.com/mhl787156/seahorse_server/db"
)

/*AuthMiddleware should redirect user to login page if not logged in*/
func AuthMiddleware(c *gin.Context) {
	if !strings.Contains(c.Request.URL.Path, "api") || c.Request.URL.Path == "/api/login" || c.Request.URL.Path == "/api/setpassword" {
		return
	}

	if c.Request.Method == "OPTIONS" {
		return
	}

	res := c.Request.Header.Get("Authorization")
	if res == "" {
		c.String(404, "{\"code\": 1003, \"message\": \"Authorization token not sent\"}")
		c.Abort()
		return
	}

	authtoken := strings.Split(res, " ")

	if authtoken[0] != "Bearer" || len(authtoken) != 3 || authtoken[1] == "" || authtoken[2] == "" {
		c.String(404, "{\"code\": 1004, \"message\": \"Authorization token not bearer format\"}")
		c.Abort()
		return
	}

	uid := authtoken[1]
	if uid == "" || uid == "null" {
		c.String(404, "{\"code\": 1005, \"message\": \"No UID provied in token\"}")
		c.Abort()
		return
	}
	fmt.Println("Authenticated API Request by:", uid)

	token := authtoken[2]

	if !authenticateToken(uid, token) {
		c.String(404, "{\"code\": 1006, \"message\": \"Token not accepted\"}")
		c.Abort()
		return
	}

	c.Next()
}

func authenticateToken(uid string, token string) bool {
	mdb := db.MongoSession
	usersecure, found, err := mdb.GetUserSecure(uid)
	if !found {
		return false
	}
	if err != nil {
		panic("Database Error")
	}

	founduid, authd := auth.Authenticate(token, usersecure.SecretKey)

	if founduid != uid {
		return false
	}
	return authd
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

}
