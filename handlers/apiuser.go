package handlers

import (
	// "fmt"
	// "github.com/mhl787156/seahorse_server/db"
	"github.com/gin-gonic/gin"
	//"github.com/mhl787156/seahorse_server/auth"
)

func getUser(c *gin.Context) {
	// rdb := db.ReactSession

	// targetId := c.Param("uid")
	// target, found, err := rdb.GetUser(targetId)
	// if err != nil {
	// 	c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	// } else if !found {
	// 	c.String(404, "{\"code\": 1002, \"message\": \"User does not exist\"}")
	// } else if target.Private {
	// 	c.String(403, "{\"code\": 1004, \"message\": \"User has set their profile to private\"}")
	// } else {
	// 	c.Header("Access-Control-Allow-Origin", "*")
	// 	c.JSON(200, target)
	// }
}

func searchUser(c *gin.Context) {
	// rdb := db.ReactSession

	// query := fmt.Sprintf("(?i)%s", c.Param("query"))
	// users, err := rdb.SearchUsers(query, 30)
	// if err != nil {
	// 	c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	// } else {
	// 	c.JSON(200, users)
	// }
}

func getPublicProj(c *gin.Context) {
	// rdb := db.ReactSession

	// targetId := c.Param("uid")
	// target, found, err := rdb.GetUser(targetId)
	// if err != nil {
	// 	c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	// } else if !found {
	// 	c.String(404, "{\"code\": 1002, \"message\": \"User does not exist\"}")
	// } else if target.Private {
	// 	c.String(403, "{\"code\": 1004, \"message\": \"User has set their profile to private\"}")
	// } else {
	// 	c.JSON(200, filterPublicProj(target.Projects))
	// }
}

func filterPublicProj(projs []string) []string {
	res := []string{}
	for _, p := range projs {
		if isProjPublic(p) {
			res = append(res, p)
		}
	}
	return res
}
