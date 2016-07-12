package auth

import (
	// "errors"
	// "fmt"
	// "time"

	"github.com/gin-gonic/gin"
	"github.com/mhl787156/seahorse_server/db"
)

const sessionCookieName string = "apx_session"

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
	// sessionKey := NewSession(uid, rdb)
	// fmt.Println("Made new session.")
	// c.SetCookie(sessionCookieName, sessionKey.SessionKey, sessionDuration, "/", "apx.twintailsare.moe", false, false)
}

//Authenticates an existing user session and retrieves the
//user ID it was assigned to
func AuthSession(c *gin.Context, rdb *db.DbConn) (string, bool, error) {
	// //Get cookie from the gin router
	// sessionKey, err := c.Cookie(sessionCookieName)
	// if err != nil {
	// 	return "", false, err
	// 	//Cookie not found, apparently...
	// }

	// //Lookup session in database
	// session, found, err := rdb.GetSession(sessionKey)
	// if err != nil {
	// 	fmt.Println("couldnt run query")
	// 	return "", false, err
	// }
	// if !found {
	// 	fmt.Println("no matching cookies")
	// 	return "", false, nil
	// } else {
	// 	curTime := time.Now().UTC().Unix()
	// 	expireTime := session.Expires
	// 	if expireTime < curTime {
	// 		return "", false, nil
	// 	} else {
	// 		return session.UID, true, nil
	// 	}
	// }
	return "", false, nil
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
