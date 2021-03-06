package handlers

import "github.com/gin-gonic/gin"

/*URI - Base URI*/
const URI string = "/api/"

/*APIHandlers URI Routing*/
func APIHandlers(router *gin.Engine) {

	//Handlers for '/me/*' endpoints
	router.GET(URI+"me", getMeHandler)

	//Handlers for '/customer/*'endpoints
	router.GET(URI+"customer/get/:id", getCustomer)
	router.GET(URI+"customer/list", getCustomerListBy)
	router.GET(URI+"customer/fulllist", getCustomerFullList)
	router.GET(URI+"customer/new", newCustomer)
	// router.GET(URI+"customer/search/:query", searchCustomer)

	router.POST(URI+"customer/update", modifyCustomer)
	// router.POST(URI+"customer/:id/:order")

	//Handlers for '/order/*'endpoints

	//Handlers for '/admin/* endpoints
	router.POST(URI+"admin/newuser", newUser)

	//Handlers for '/product/* endpoints

	//Handler for '/login endpoints
	router.POST(URI+"login", loginUser)
	router.POST(URI+"setpassword", setUserPassword)

	// //Handlers for '/user/*' endpoints
	// router.GET("/api/user/:uid", getUser)
	// router.GET("/api/search/user/:query", searchUser)
	// router.GET("/api/user/:uid/*action", func(c *gin.Context) {
	// 	if c.Param("action") == "projects" {
	// 		getPublicProj(c)
	// 	}
	// })

	// //Handlers for '/project/*' endpoints
	// router.POST("/api/project", postNewProject)
	// router.GET("/api/project/:pid", getProject)
	// router.POST("/api/project/:pid/*action", func(c *gin.Context) {
	// 	postProjHandler(c, c.Param("action"))
	// })
	// router.GET("/api/projectcontent/:pid", getProjectContent)
	// router.POST("/api/projectcontent/:pid", writeProjectContent)

	// //Handlers for '/snippet/*' endpoints
	// router.POST("/api/snippet", postNewSnippet)
	// router.GET("/api/snippet/:sid", getSnippet)
	// router.POST("/api/snippet/:sid/*action", func(c *gin.Context) {
	// 	postSnippetHandler(c, c.Param("action"))
	// })
	// router.GET("/api/snippetcontent/:sid", getSnippetContent)
	// router.POST("/api/snippetcontent/:sid", writeSnippetContent)
	// router.DELETE("/api/snippet/:sid", delSnippet)
}
