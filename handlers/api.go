package handlers

import (
	"github.com/gin-gonic/gin"
)

func ApiHandlers(router *gin.Engine) {
	//Handlers for '/me/*' endpoints
	router.GET("/api/me", getMeHandler)
	router.GET("/api/me/following", getFriendHandler)
	router.GET("/api/me/projects", getProjHandler)
	router.GET("/api/me/snippets", getSnippetHandler)
	router.GET("/api/me/projects/meta", getProjMetaHandler)
	router.POST("/api/me", func(c *gin.Context) {
		postMeHandler(c, "general")
	})
	router.POST("/api/me/follow", func(c *gin.Context) {
		postMeHandler(c, "addfriend")
	})
	router.POST("/api/me/unfollow", func(c *gin.Context) {
		postMeHandler(c, "removefriend")
	})

	//Handlers for '/user/*' endpoints
	router.GET("/api/user/:uid", getUser)
	router.GET("/api/search/user/:query", searchUser)
	router.GET("/api/user/:uid/*action", func(c *gin.Context) {
		if c.Param("action") == "projects" {
			getPublicProj(c)
		}
	})

	//Handlers for '/project/*' endpoints
	router.POST("/api/project", postNewProject)
	router.GET("/api/project/:pid", getProject)
	router.POST("/api/project/:pid/*action", func(c *gin.Context) {
		postProjHandler(c, c.Param("action"))
	})
	router.GET("/api/projectcontent/:pid", getProjectContent)
	router.POST("/api/projectcontent/:pid", writeProjectContent)

	//Handlers for '/snippet/*' endpoints
	router.POST("/api/snippet", postNewSnippet)
	router.GET("/api/snippet/:sid", getSnippet)
	router.POST("/api/snippet/:sid/*action", func(c *gin.Context) {
		postSnippetHandler(c, c.Param("action"))
	})
	router.GET("/api/snippetcontent/:sid", getSnippetContent)
	router.POST("/api/snippetcontent/:sid", writeSnippetContent)
	router.DELETE("/api/snippet/:sid", delSnippet)
}
