package db

import "github.com/mhl787156/seahorse_server/models"

//GetUUID generates a new Unique Identifier to use as an ID
func (c *DbConn) GetUUID() string {

	// query, err := gorethink.UUID().Run(c.Session)
	// if err != nil {
	// 	panic("Database broke again")
	// } else {
	// 	var uuid string
	// 	err = query.One(&uuid)
	// 	if err != nil {
	// 		panic("Could not generate a uuid")
	// 	} else {
	// 		return uuid
	// 	}
	// }
	return ""
}

// WriteUser inserts a new User into the database.
// User is inserted into the users collection
func (c *DbConn) WriteUser(user models.User) error {
	collection := c.DB.C(models.CollectionUsers)
	return collection.Insert(user)
}

// GetUserByID looks in the database by uuid.
// Returns the user, a bool to check whether any user was found,
// or whether an Error was returned
func (c *DbConn) GetUserByID(id string) (*models.User, bool, error) {
	collection := c.DB.C(models.CollectionUsers)
	query := collection.FindId(id)

	//Check that a result was found
	if query == nil {
		//No results were found
		return nil, false, nil
	}
	foundUser := models.User{}
	err := query.One(&foundUser)
	if err != nil {
		return nil, false, err
	}
	return &foundUser, true, nil
}

// SearchUsers looks in the database by queryString and returns a maximum of limit number of entries.
// The query string will most likely be a form of bson.M{"k1":"v1", "k2":"v2"...}.
// Function returns the list of users found, a check bool to see if the query was successful and any error.
func (c *DbConn) SearchUsers(queryStr interface{}, limit int) ([]models.User, bool, error) {
	if limit == 0 {
		return nil, false, nil
	}
	collection := c.DB.C(models.CollectionUsers)
	query := collection.Find(queryStr).Limit(limit)

	//Check that a result was found
	if query == nil {
		//No results were found
		return nil, false, nil
	}
	var err error
	var foundUser []models.User
	query.All(&foundUser)
	if err != nil {
		return nil, false, err
	}
	return foundUser, true, nil
}

func (c *DbConn) ModifyUser(user *models.User) (bool, error) {
	// res, err := UserTable.Get(user.Id).Update(*user).RunWrite(c.Session)
	// if err != nil {
	// 	return false, err
	// } else if res.Replaced == 0 {
	// 	return false, nil
	// } else {
	// 	return true, err
	// }
	return false, nil
}

// GetUserOrders looks in the database by uuid.
// Returns a users orders, a bool to check whether any orders were found for this user,
// or whether an Error was returned\
// If no UserOrder Struct is found, a new one is created
func (c *DbConn) GetUserOrders(id string) (*models.UserOrders, bool, error) {
	collection := c.DB.C(models.CollectionUserOrders)
	query := collection.FindId(id)

	//Check that a result was found
	if query == nil {
		//No results were found, new one created for this user
		var list []string
		err := collection.Insert(models.UserOrders{ID: id, OrderIds: list})

		return nil, false, err
	}
	foundUserOrders := models.UserOrders{}
	err := query.One(&foundUserOrders)
	if err != nil {
		return nil, false, err
	}
	return &foundUserOrders, true, nil
}

// AddUserOrders appends an orderid to the users list of orders
func (c *DbConn) AddUserOrders(id string, orderid string) (bool, error) {
	return false, nil
}

// GetUserSecure gets the Secure Details of a user with given id id
// Returns the usersecure, a boolean of if it exists or the error
func (c *DbConn) GetUserSecure(id string) (*models.UserSecure, bool, error) {
	collection := c.DB.C(models.CollectionUserSecure)
	query := collection.FindId(id)

	//Check that a result was found
	if query == nil {
		//No results were found
		return nil, false, nil
	}
	foundUserSecure := models.UserSecure{}
	err := query.One(&foundUserSecure)
	if err != nil {
		return nil, false, err
	}
	return &foundUserSecure, true, nil
}

// SetUserSecure adds a secure user document to the database
// Sets the password field for the user
// If a usersecure already exists for this user, the password is not changed
func (c *DbConn) SetUserSecure(id string, pword []byte) (bool, error) {
	collection := c.DB.C(models.CollectionUserSecure)
	query := collection.FindId(id)
	n, err := query.Count()
	if n == 0 {
		//No remaining password exists for this user, can add a password
		err := collection.Insert(models.UserSecure{ID: id, Password: pword, SecretKey: nil})
		return true, err
	}
	return false, err
}

// SetUserSecretKey sets the User secure for a given known user if the user exists.
// SecretKey interface will msot likely be a bson.M{"secretkey: "asd"}
// Returns an ErrNotFound if user does not exist
func (c *DbConn) SetUserSecretKey(userID string, secretkey interface{}) error {
	collection := c.DB.C(models.CollectionUserSecure)
	return collection.UpdateId(userID, secretkey)
}
