package handlers

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/mhl787156/seahorse_server/auth"
	"github.com/mhl787156/seahorse_server/db"
	"github.com/mhl787156/seahorse_server/models"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

// checkAdmin is the initial check to see ensure that there exists at least one admin in the system
// if no admin is found, a new user is inserted into the database of username 'admin'
// password is not set, the user must set a password for this user.
func CheckAdmin() {
	mdb := db.MongoSession
	_, found, err2 := mdb.SearchUsers(bson.M{"admin": true}, 1)
	if !found || err2 != nil {
		mdb.WriteUser(models.User{
			ID:        uuid.NewV4().String(),
			Admin:     true,
			DateAdded: time.Now().Unix(),
			Username:  "admin",
		})
	}
}

// loginUser is POST request handler for the login of a User (Employee)
// Post body contains LoginDetails.
// First we check to see if the user exists, if not, we pass back a 404
// If the user exists, we then generate a new token, check the password against the stored hash,
// Then store the jwt secret key
// We must then generate a new jwt token, save it and pass the token back to the request
func loginUser(c *gin.Context) {
	var logindetails models.LoginDetails
	err := c.BindJSON(&logindetails)

	if err != nil {
		c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred in binding\"}")
		return
	}

	mdb := db.MongoSession

	// Does the user exist
	users, found, err1 := mdb.SearchUsers(bson.M{"username": logindetails.Username}, 1)

	if err1 != nil {
		c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred in searching\"}")
		return
	}

	if !found {
		c.String(400, "{\"code\": 1002, \"message\": \"User does not exist\"}")
		return
	}

	user := users[0]
	// Has the user set his password
	usersecure, found, _ := mdb.GetUserSecure(user.ID)

	if !found {
		c.String(401, "{\"code\": 1002, \"message\": \"User needs to set password\"}")
		return
	}

	// Does the password match
	if bcrypt.CompareHashAndPassword(usersecure.Password, []byte(logindetails.Password)) != nil {
		c.String(422, "{\"code\": 1003, \"message\": \"Password does not match\"}")
		return
	}

	// Generate JWT Token
	token, key, err2 := auth.GenerateToken(user.ID)

	if err2 != nil {
		c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred in token generation\"}")
		return
	}

	// Save the secretkey for the jwt
	err4 := mdb.SetUserSecretKey(user.ID, bson.M{"secretkey": []byte(key)})

	if err4 != nil {
		c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred in setting keys\"}")
		return
	}

	// If everything works, return the token
	c.JSON(200, gin.H{"user": user.ID, "token": token})
}

// setUserPassword is POST handler for a user setting their password for the first time
// we check if user exists, if not we return 404, code 1002
// if user exists we pass the bcrypt hashed password to save in the database
// if the password is already set, we return 403, code 1004 and the password is not set.
func setUserPassword(c *gin.Context) {
	var securedetails models.LoginDetails
	err1 := c.BindJSON(&securedetails)

	if err1 != nil {
		c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred in binding\"}")
		return
	}

	mdb := db.MongoSession
	//Does user exist
	users, found, err2 := mdb.SearchUsers(bson.M{"username": securedetails.Username}, 1)

	if err2 != nil {
		c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred in Searching the Database\"}")
		return
	}

	if !found {
		c.String(400, "{\"code\": 1002, \"message\": \"User does not exist\"}")
		return
	}

	//Hash password
	hashedpwd, err3 := bcrypt.GenerateFromPassword([]byte(securedetails.Password), bcrypt.DefaultCost)
	if err3 != nil {
		c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred in Hashing\"}")
		return
	}

	set, err4 := mdb.SetUserSecure(users[0].ID, hashedpwd)

	if err4 != nil {
		c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred Writing to DB\"}")
		return
	}

	if !set {
		c.String(403, "{\"code\": 1004, \"message\": \"User password already set, ask admin to reset\"}")
		return
	}

	c.String(200, "Password set")
}

// newUser is a POST request handler for adding New Users to the system (Employees)
// Post body contains user details
// Can only be called by someone with admin rights
func newUser(c *gin.Context) {
	var newuser models.User
	err1 := c.BindJSON(&newuser)

	if err1 != nil {
		c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred in binding\"}")
		return
	}

	newuser.ID = uuid.NewV4().String()
	newuser.DateAdded = time.Now().Unix()

	mdb := db.MongoSession
	err2 := mdb.WriteUser(newuser)

	if err2 != nil {
		c.String(500, "{\"code\": -1, \"message\": \"An unexpected error occurred in writing\"}")
	} else {
		c.JSON(200, "User Added")
	}
}
