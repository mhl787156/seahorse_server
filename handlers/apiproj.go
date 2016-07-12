package handlers

// "github.com/mhl787156/seahorse_server/auth"
// "github.com/mhl787156/seahorse_server/db"

// "fmt"

// func isProjPublic(pid string) bool {
// 	// rdb := db.ReactSession
// 	// _, found, err := rdb.GetProject(pid)

// 	// if err != nil {
// 	// 	return false
// 	// } else if !found {
// 	// 	return false
// 	// } else {
// 	// 	return true
// 	// }
// 	return false
// }

// func postNewProject(c *gin.Context) {
// 	// rdb := db.ReactSession
// 	// uid, found, err := auth.AuthSession(c, rdb)
// 	// if !found && err != nil {
// 	// 	c.String(401, "{\"code\": 1001, \"message\": \"No session key was provided\"}")
// 	// } else if !found {
// 	// 	c.String(403, "{\"code\": 1000, \"message\": \"Could not find that session\"}")
// 	// } else if err != nil {
// 	// 	c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
// 	// } else {
// 	// 	me, found, err := rdb.GetUser(uid)
// 	// 	if err != nil {
// 	// 		c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
// 	// 	} else if !found {
// 	// 		c.String(404, "{\"code\": 1002, \"message\": \"User does not exist\"}")
// 	// 	} else {
// 	// 		proj, err := rdb.WriteProject(me)
// 	// 		if err != nil {
// 	// 			c.String(500, `{"code": -1, "message": "Could not insert a new project"}`)
// 	// 		} else {
// 	// 			c.JSON(201, proj)
// 	// 		}
// 	// 	}
// 	// }
// }

// func getProject(c *gin.Context) {
// 	// rdb := db.ReactSession
// 	// pid := c.Param("pid")
// 	// proj, found, err := rdb.GetProject(pid)

// 	// if !found && err == nil {
// 	// 	c.String(404, "{\"code\": 2000, \"message\": \"Could not find that project\"}")
// 	// } else if err != nil {
// 	// 	c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
// 	// } else {
// 	// 	if proj.Public {
// 	// 		c.JSON(200, proj)
// 	// 	} else {
// 	// 		uid, _, err := auth.AuthSession(c, rdb)
// 	// 		if err != nil {
// 	// 			c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
// 	// 		} else if userOwnsProj(uid, proj) {
// 	// 			c.JSON(201, proj)
// 	// 		} else {
// 	// 			c.String(403, `{"code": 2001, "message": "You need to be an owner to view that project"}`)
// 	// 		}
// 	// 	}
// 	// }

// }

// func userOwnsProj(uid string, proj *models.Project) bool {
// 	for _, u := range proj.Owners {
// 		if u == uid {
// 			return true
// 		}
// 	}
// 	return false
// }

// func postProjHandler(c *gin.Context, mod string) {
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
// 	//   pid := c.Param("pid")
// 	//   proj, found, err := rdb.GetProject(pid)

// 	//   if !found && err == nil {
// 	//     c.String(404, "{\"code\": 2000, \"message\": \"Could not find that project\"}")
// 	//   } else if err != nil {
// 	//     c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
// 	//   } else if (!userOwnsProj(curUser, proj)) {
// 	//     c.String(403, "{\"code\": 2001, \"message\": \"You do not own that project\"}")
// 	//   } else if (mod == "/meta"){
// 	//       modifyProjMeta(c, proj)
// 	//   } else if (mod == "/addowner"){
// 	//     addProjOwner(c, proj)
// 	//   } else if (mod == "/leave") {
// 	//     leaveProj(c, curUser)
// 	//   } else {
// 	//     fmt.Println(mod)
// 	//   }
// 	// }
// }

// func modifyProjMeta(c *gin.Context, proj *models.Project) {
// 	// rdb := db.ReactSession
// 	// oldPID := proj.Id
// 	// oldOwners := proj.Owners
// 	// c.BindJSON(proj)
// 	// proj.Id = oldPID
// 	// proj.Owners = oldOwners

// 	// modified, err := rdb.ModifyProject(proj)
// 	// if (err != nil) {
// 	//   c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
// 	// } else if (!modified) {
// 	//   c.String(418, "{\"code\": 0, \"message\": \"User is a teapot.\"}")
// 	// } else {
// 	//   c.Status(201)
// 	// }
// }

// func addProjOwner(c *gin.Context, proj *models.Project) {
// 	// rdb := db.ReactSession
// 	// req := map[string]string{}

// 	// c.BindJSON(&req)

// 	// newOwner := req["uid"]

// 	// newOwn, found, err := rdb.GetUser(newOwner)
// 	// if (err != nil) {
// 	//   c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
// 	//   return
// 	// } else if (!found) {
// 	//   c.String(404, "{\"code\": 1002, \"message\": \"User does not exist\"}")
// 	//   return
// 	// }

// 	// proj.Owners = append(proj.Owners, newOwner)
// 	// newOwn.Projects = append(newOwn.Projects, proj.Id)
// 	// modified, err := rdb.ModifyProject(proj)
// 	// if (err != nil) {
// 	//   c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
// 	//   return
// 	// } else if (!modified) {
// 	//   c.String(418, "{\"code\": 0, \"message\": \"User is a teapot.\"}")
// 	//   return
// 	// } else {
// 	//   modified, err := rdb.ModifyUser(newOwn)
// 	//   if (err != nil) {
// 	//     c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
// 	//   } else if (!modified) {
// 	//     c.String(418, "{\"code\": 0, \"message\": \"User is a teapot.\"}")
// 	//   } else {
// 	//     c.Status(201)
// 	//   }
// 	// }
// }

// func leaveProj(c *gin.Context, uid string) {
// 	// rdb := db.ReactSession

// 	// proj := c.Param("pid")
// 	// modified, err := rdb.RemoveUserFromProject(proj, uid)
// 	// if (err != nil) {
// 	//   c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred\"}")
// 	//   return
// 	// } else if (!modified) {
// 	//   c.String(403, "{\"code\": 2001, \"message\": \"User is not a member of that project.\"}")
// 	//   return
// 	// } else {
// 	//   c.Status(201)
// 	// }
// }
