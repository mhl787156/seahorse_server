package handlers

import (
	// "encoding/json"
	// "fmt"

	"github.com/gin-gonic/gin"
	// "github.com/mhl787156/seahorse_server/db"
	"github.com/mhl787156/seahorse_server/models"
)

func getProjectContent(c *gin.Context) {
	// rdb := db.ReactSession
	// pid := c.Param("pid")
	// proj, err := rdb.GetProjectContent(pid)

	// if err != nil {
	// 	c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	// } else {
	// 	c.JSON(201, proj)
	// }

}

func modifyProjContent(c *gin.Context, projC *models.ProjectContent) *string {
	// rdb := db.ReactSession
	// oldPID := projC.Id
	// c.BindJSON(projC)
	// projC.Id = oldPID
	// modified, err := rdb.ModifyProjectContent(projC)
	// if err != nil {
	// 	c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	// 	return nil
	// } else if !modified {
	// 	c.String(418, "{\"code\": 0, \"message\": \"User is a teapot.\"}")
	// 	return nil
	// } else {
	// 	c.Status(201)
	// 	bs, _ := json.Marshal(projC)
	// 	s := string(bs)
	// 	return &s
	// }
	return nil
}

func writeProjectContent(c *gin.Context) {
	// rdb := db.ReactSession
	// pid := c.Param("pid")
	// proj, err := rdb.GetProjectContent(pid)
	// if err != nil {
	// 	c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	// } else {
	// 	fmt.Println(proj)
	// 	s := modifyProjContent(c, proj)
	// }
}
