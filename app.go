package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/mhl787156/seahorse_server/auth"
	"github.com/mhl787156/seahorse_server/conf"
	"github.com/mhl787156/seahorse_server/db"
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
	go db.MongoSession.CullSessions()

	// Middlewares
	// router.Use(middlewares.Connect)
	router.Use(middlewares.ErrorHandler)
	// router.Use(middlewares.AuthMiddleware)
	router.Use(middlewares.CORS)

	// //Serve static angular files
	router.StaticFile("/", pageLoc("/dist/dev/index.html"))
	// router.StaticFile("/home", pageLoc("/dist/dev/index.html"))
	// router.StaticFile("/admin", pageLoc("/dist/dev/index.html"))
	// router.StaticFile("/login", pageLoc("/dist/dev/index.html"))

	router.GET("/product/*d", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, pageLoc("/dist/dev/index.html"))
	})

	router.GET("/order/*d", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, pageLoc("/dist/dev/index.html"))
	})

	router.GET("/login/*d", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, pageLoc("/dist/dev/index.html"))
	})

	router.GET("/admin/*d", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, pageLoc("/dist/dev/index.html"))
	})

	router.GET("/home/*d", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, pageLoc("/dist/dev/index.html"))
	})

	router.GET("/customer/*d", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, pageLoc("/dist/dev/index.html"))
	})
	// http.Handle("/", http.FileServer(http.Dir(pageLoc("/dist/dev"))))

	// // router.StaticFile("/systemjs.config.js", pageLoc("/systemjs.config.js"))
	router.StaticFS("/dist/dev", http.Dir(pageLoc("/dist/dev")))
	router.StaticFS("/app", http.Dir(pageLoc("/dist/dev/app")))
	router.StaticFS("/node_modules", http.Dir(pageLoc("/node_modules")))

	// //API Endpoints
	// handlers.APIHandlers(router)

	// //Redirect for logged-in users
	// router.GET("/loggedin", func(c *gin.Context) {
	// 	c.Redirect(303, "/dashboard")
	// })

	//Login page which installs the session cookie
	// router.GET("/fbauth", auth.AuthHandler)

	//Redirect for landing page
	//router.GET("/", handlers.LandingHandler)

	router.NoRoute(func(c *gin.Context) {
		fmt.Printf("%v\n", "No Route!")
		http.ServeFile(c.Writer, c.Request, pageLoc("/dist/dev/index.html"))
	})

	//Run the server
	router.Run(fmt.Sprintf(":%d", conf.Port))

}
