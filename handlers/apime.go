package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mhl787156/seahorse_server/models"
)

func getMeHandler(c *gin.Context) {
	// rdb := db.ReactSession

	// //Get the cookie
	// curUser, found, err := auth.AuthSession(c, rdb)
	// if (!found && err != nil) {
	//   c.String(401, "{\"code\": 1001, \"message\": \"No session key was provided\"}")
	// } else if (!found) {
	//   c.String(403, "{\"code\": 1000, \"message\": \"Could not find that session\"}")
	// } else if (err != nil) {
	//   c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	// } else {
	//   me, found, err := rdb.GetUser(curUser)
	//   if (err != nil){
	//     c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	//   } else if (!found) {
	//     c.String(404, "{\"code\": 1002, \"message\": \"User does not exist\"}")
	//   } else {
	//     c.JSON(200, me)
	//   }
	// }
}

func postMeHandler(c *gin.Context, mod string) {
	// rdb := db.ReactSession

	// //Get the cookie
	// curUser, found, err := auth.AuthSession(c, rdb)
	// if (!found && err != nil) {
	//   c.String(401, "{\"code\": 1001, \"message\": \"No session key was provided\"}")
	// } else if (!found) {
	//   c.String(403, "{\"code\": 1000, \"message\": \"Could not find that session\"}")
	// } else if (err != nil) {
	//   c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	// } else {
	//   me, found, err := rdb.GetUser(curUser)
	//   if (err != nil){
	//     c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	//   } else if (!found || me == nil) {
	//     c.String(404, "{\"code\": 1002, \"message\": \"User does not exist\"}")
	//   } else {
	//     if (mod == "general"){
	//       modifyUser(c, me)
	//     } else if (mod == "addfriend") {
	//       addFriend(c, me)
	//     } else if (mod == "removefriend") {
	//       removeFriend(c, me)
	//     }
	//   }
	// }
}

func removeFriend(c *gin.Context, me *models.User) {
	// rdb := db.ReactSession
	// req := map[string]string{}

	// c.BindJSON(&req)
	// newNemesis := req["uid"]

	// for i, uid := range me.Following {
	//   if (uid == newNemesis) {
	//     me.Following = append(me.Following[:i], me.Following[i+1:]...)
	//     modified, err := rdb.ModifyUser(me)
	//     if (err != nil) {
	//       c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	//       return
	//     } else if (!modified) {
	//       c.String(418, "{\"code\": 0, \"message\": \"User is a teapot.\"}")
	//       return
	//     } else {
	//       c.Status(201)
	//       return
	//     }
	//   }
	// }
	// c.String(404, "{\"code\": 1003, \"message\": \"They never were your friend...\"}")
}

func addFriend(c *gin.Context, me *models.User) {
	// rdb := db.ReactSession
	// req := map[string]string{}

	// c.BindJSON(&req)

	// newFriend := req["uid"]
	// for _, uid := range me.Following {
	//   if (uid == newFriend) {
	//     c.Status(200)
	//     return
	//   }
	// }

	// _, found, err := rdb.GetUser(newFriend)
	// if (err != nil) {
	//   c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	// } else if (!found) {
	//   c.String(404, "{\"code\": 1002, \"message\": \"User does not exist\"}")
	// }
	// me.Following = append(me.Following, newFriend)
	// modified, err := rdb.ModifyUser(me)
	// if (err != nil) {
	//   c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	// } else if (!modified) {
	//   c.String(418, "{\"code\": 0, \"message\": \"User is a teapot.\"}")
	// } else {
	//   c.Status(201)
	// }
}

func modifyUser(c *gin.Context, me *models.User) {
	// rdb := db.ReactSession
	// oldUID := me.Id
	// c.BindJSON(me)
	// me.Id = oldUID
	// modified, err := rdb.ModifyUser(me)
	// if (err != nil) {
	//   c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	// } else if (!modified) {
	//   c.String(418, "{\"code\": 0, \"message\": \"User is a teapot.\"}")
	// } else {
	//   c.Status(201)
	// }

}

func getFriendHandler(c *gin.Context) {
	// rdb := db.ReactSession

	// //Get the cookie
	// curUser, found, err := auth.AuthSession(c, rdb)
	// if (!found && err != nil) {
	//   c.String(401, "{\"code\": 1001, \"message\": \"No session key was provided\"}")
	// } else if (!found) {
	//   c.String(403, "{\"code\": 1000, \"message\": \"Could not find that session\"}")
	// } else if (err != nil) {
	//   c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	// } else {
	//   me, found, err := rdb.GetUser(curUser)
	//   if (err != nil){
	//     c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	//   } else if (!found) {
	//     c.String(404, "{\"code\": 1002, \"message\": \"User does not exist\"}")
	//   } else {
	//     c.JSON(200, me.Following)
	//   }
	// }
}

func getProjHandler(c *gin.Context) {
	// rdb := db.ReactSession

	// //Get the cookie
	// curUser, found, err := auth.AuthSession(c, rdb)
	// if (!found && err != nil) {
	//   c.String(401, "{\"code\": 1001, \"message\": \"No session key was provided\"}")
	// } else if (!found) {
	//   c.String(403, "{\"code\": 1000, \"message\": \"Could not find that session\"}")
	// } else if (err != nil) {
	//   c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	// } else {
	//   me, found, err := rdb.GetUser(curUser)
	//   if (err != nil){
	//     c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	//   } else if (!found) {
	//     c.String(404, "{\"code\": 1002, \"message\": \"User does not exist\"}")
	//   } else {
	//     c.JSON(200, me.Projects)
	//   }
	// }
}

func getSnippetHandler(c *gin.Context) {
	// rdb := db.ReactSession

	// //Get the cookie
	// curUser, found, err := auth.AuthSession(c, rdb)
	// if (!found && err != nil) {
	//   c.String(401, "{\"code\": 1001, \"message\": \"No session key was provided\"}")
	// } else if (!found) {
	//   c.String(403, "{\"code\": 1000, \"message\": \"Could not find that session\"}")
	// } else if (err != nil) {
	//   c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	// } else {
	//   snippets, err := rdb.GetSnippets(curUser)
	//   if err != nil {
	//     fmt.Println(err)
	//     c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	//   } else {
	//     c.JSON(200, snippets)
	//   }
	// }
}

func getProjMetaHandler(c *gin.Context) {
	// rdb := db.ReactSession

	// //Get the cookie
	// curUser, found, err := auth.AuthSession(c, rdb)
	// if (!found && err != nil) {
	//   c.String(401, "{\"code\": 1001, \"message\": \"No session key was provided\"}")
	// } else if (!found) {
	//   c.String(403, "{\"code\": 1000, \"message\": \"Could not find that session\"}")
	// } else if (err != nil) {
	//   c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	// } else {
	//   me, found, err := rdb.GetUser(curUser)
	//   if (err != nil){
	//     c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
	//   } else if (!found) {
	//     c.String(404, "{\"code\": 1002, \"message\": \"User does not exist\"}")
	//   } else {
	//     res := []models.Project{}
	//     for _, pid := range me.Projects {
	//       proj, found, err := rdb.GetProject(pid)
	//       if (!found || err != nil) {
	//         panic("wat")
	//       }
	//       res = append(res, *proj)
	//     }
	//     c.JSON(200, res)
	//   }
	// }
}
