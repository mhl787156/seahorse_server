package handlers

import (
	"github.com/gin-gonic/gin"
	// "github.com/mhl787156/seahorse_server/db"
	// "github.com/mhl787156/seahorse_server/models"
)

func getCustomer(c *gin.Context) {
	// 	 rdb := db.MongoSession

	// 	// targetId := c.Param("id")
	// 	// target, found, err := rdb.GetUser(targetId)
	// 	// if err != nil {
	// 	// 	c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	// 	// } else if !found {
	// 	// 	c.String(404, "{\"code\": 1002, \"message\": \"User does not exist\"}")
	// 	// } else if target.Private {
	// 	// 	c.String(403, "{\"code\": 1004, \"message\": \"User has set their profile to private\"}")
	// 	// } else {
	// 	// 	c.Header("Access-Control-Allow-Origin", "*")
	// 	// 	c.JSON(200, target)
	// 	// }
}

// func postCustomer(c *gin.Context, mod string) {
// 	// rdb := db.ReactSession

// 	// //Get the cookie
// 	// curUser, found, err := auth.AuthSession(c, rdb)
// 	// if (!found && err != nil) {
// 	//   c.String(401, "{\"code\": 1001, \"message\": \"No session key was provided\"}")
// 	// } else if (!found) {
// 	//   c.String(403, "{\"code\": 1000, \"message\": \"Could not find that session\"}")
// 	// } else if (err != nil) {
// 	//   c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
// 	// } else {
// 	//   me, found, err := rdb.GetUser(curUser)
// 	//   if (err != nil){
// 	//     c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
// 	//   } else if (!found || me == nil) {
// 	//     c.String(404, "{\"code\": 1002, \"message\": \"User does not exist\"}")
// 	//   } else {
// 	//     if (mod == "general"){
// 	//       modifyUser(c, me)
// 	//     } else if (mod == "addfriend") {
// 	//       addFriend(c, me)
// 	//     } else if (mod == "removefriend") {
// 	//       removeFriend(c, me)
// 	//     }
// 	//   }
// 	// }
// }

// func modifyCustomer(c *gin.Context, customer *models.Customer) {
// 	// rdb := db.ReactSession
// 	// oldUID := me.Id
// 	// c.BindJSON(me)
// 	// me.Id = oldUID
// 	// modified, err := rdb.ModifyUser(me)
// 	// if (err != nil) {
// 	//   c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
// 	// } else if (!modified) {
// 	//   c.String(418, "{\"code\": 0, \"message\": \"User is a teapot.\"}")
// 	// } else {
// 	//   c.Status(201)
// 	// }

// }

func searchCustomer(c *gin.Context) {
	// 	rdb := db.MongoSession

	// 	// query := fmt.Sprintf("(?i)%s", c.Param("query"))
	// 	// users, err := rdb.SearchUsers(query, 30)
	// 	// if err != nil {
	// 	// 	c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	// 	// } else {
	// 	// 	c.JSON(200, users)
	// 	// }
}
