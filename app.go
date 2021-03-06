package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	// "github.com/mhl787156/seahorse_server/auth"
	"github.com/mhl787156/seahorse_server/conf"
	"github.com/mhl787156/seahorse_server/db"
	"github.com/mhl787156/seahorse_server/handlers"
	"github.com/mhl787156/seahorse_server/middlewares"
)

var siteRoot string

func pageLoc(rel string) string {
	return siteRoot + rel
}

func main() {
	//Parse server config files
	conf := conf.GetConfig()
	siteRoot = conf.HtmlDir

	//Connect to the database and start gin router
	db.Connect(conf.DBIPAddress)
	router := gin.Default()

	//Start function to clear expired sessions
	go db.MongoSession.CullSessions()

	// Check that there is at least one admin user
	handlers.CheckAdmin()

	// Middlewares
	// router.Use(middlewares.Connect)
	// router.Use(middlewares.ErrorHandler)
	router.Use(middlewares.AuthMiddleware)
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders:  "Origin, Authorization, Content-Type, Accept",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	// //Serve static angular files
	router.StaticFile("/", pageLoc("/dist/dev/index.html"))

	router.StaticFS("/dist/dev", http.Dir(pageLoc("/dist/dev")))
	router.StaticFS("/app", http.Dir(pageLoc("/dist/dev/app")))
	router.StaticFS("/node_modules", http.Dir(pageLoc("/node_modules")))

	// //API Endpoints
	handlers.APIHandlers(router)

	// //Redirect for logged-in users
	// router.GET("/loggedin", func(c *gin.Context) {
	// 	c.Redirect(303, "/dashboard")
	// })

	//Login page which installs the session cookie
	// router.GET("/fbauth", auth.AuthHandler)

	router.NoRoute(func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, pageLoc("/dist/dev/index.html"))
	})

	//Run the server
	router.Run(fmt.Sprintf(":%d", conf.Port))

}
