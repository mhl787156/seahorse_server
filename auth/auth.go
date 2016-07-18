package auth

import (
	// "errors"
	// "fmt"
	// "time"

	// "github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"github.com/mhl787156/seahorse_server/db"
	// "github.com/mhl787156/seahorse_server/models"
)

const sessionCookieName string = "apx_session"

/*
*	Auth Handler called by middleware on all requests that are not login
* 	Or called directly by handlers to check validity before response is returned
*/
func AuthHandler(c *gin.Context) {
	// //Get the authentication token from the request
	// authToken := c.Query("auth_token")
	// //Retrieve the database connection
	// rdb := db.ReactSession

	// uid, err := AuthenticateUser(authToken, rdb)
	// if err != nil {
	// 	//Token is invalid
	// 	panic("Token error, should probably handle this better...")
	// }
}

func AuthenticateUser(token string, rdb *db.DbConn) (string, error) {
	// userDetails := GetUserDetails(token)
	// if !userDetails.Is_valid {
	// 	return "", errors.New("Token provided is not valid or has expired.")
	// }

	// //Lookup the user in database
	// fid := userDetails.User_id
	// user, found, err := rdb.GetFBUser(fid)
	// if err != nil {
	// 	return "", err
	// } else if found {
	// 	//If the user exists, return the user id and nil
	// 	return user.Id, nil
	// } else {
	// //Otherwise, make a new user with BuildUser, and add it to the database
	// newUser := BuildUser(userDetails, rdb)
	// rdb.WriteUser(newUser)
	// user = &newUser
	// fmt.Println("Writing new user")
	// }
	// return user.Id, nil
	return "", nil
}
