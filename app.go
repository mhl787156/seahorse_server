package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhl787156/seahorse_server/auth"
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
	db.Connect()
	router := gin.Default()

	//Start function to clear expired sessions
	// go db.ReactSession.CullSessions()

	// Middlewares
	//router.Use(middlewares.Connect)
	router.Use(middlewares.ErrorHandler)
	router.Use(middlewares.AuthMiddleware)
	router.Use(middlewares.CORS)

	//Serve static angular files
	router.StaticFile("/home", pageLoc("/index.html"))
	router.StaticFile("/login", pageLoc("/index.html"))
	router.StaticFile("/explore", pageLoc("/index.html"))

	router.GET("/dashboard/*d", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, pageLoc("/index.html"))
	})
	router.GET("/project/*d", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, pageLoc("/index.html"))
	})
	router.GET("/profile/*d", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, pageLoc("/index.html"))
	})

	router.StaticFile("/systemjs.config.js", pageLoc("/systemjs.config.js"))
	router.StaticFS("/built/app", http.Dir(pageLoc("/built/app/")))
	router.StaticFS("/app", http.Dir(pageLoc("/app/")))
	router.StaticFS("/node_modules", http.Dir(pageLoc("/node_modules")))
	router.StaticFile("/js/fblogin.js", pageLoc("/js/fblogin.js"))

	//API Endpoints
	handlers.ApiHandlers(router)

	//Redirect for logged-in users
	router.GET("/loggedin", func(c *gin.Context) {
		c.Redirect(303, "/dashboard")
	})

	//Login page which installs the session cookie
	router.GET("/fbauth", auth.AuthHandler)

	//Redirect for landing page
	//router.GET("/", handlers.LandingHandler)

	router.NoRoute(func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, pageLoc("/index.html"))
	})

	//Run the server
	router.Run(fmt.Sprintf(":%d", conf.Port))

}