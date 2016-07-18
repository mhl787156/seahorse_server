package models

const (
	//CollectionUsers holds name of users collection
	CollectionUsers = "users"

	//CollectionUserOrder holds all of a users orders
	CollectionUserOrders = "userorders"

	//CollectionUserSecure holds all of a users orders
	CollectionUserSecure = "usersecure"
)

//User model stores the details of a user
type User struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Admin     bool   `json:"admin" bson:"admin"`
	DateAdded int64  `json:"date_added" bson:"date_added"`

	Username  string `json:"username" bson:"username"`
	Firstname string `json:"firstname" bson:"firstname"`
	Surname   string `json:"surname" bson:"surname"`

	HomePhone   string `json:"home_phone" bson:"home_phone"`
	MobilePhone string `json:"mobile_phone" bson:"mobile_phone"`
	FaxNumber   string `json:"fax_number" bson:"fax_number"`

	Email string `json:"email" bson:"email"`
}

//UserOrders model stores all the orders of a given user
type UserOrders struct {
	ID       string   `json:"id" bson:"_id,omitempty"`
	OrderIds []string `json:"orderids" bson:"orderids"`
}

//UserSecure model stores important details of a user
type UserSecure struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Password  []byte `json:"password" bson:"password"`
	SecretKey []byte `json:"secretkey" bson:"secretkey"`
}

/*LoginDetails model*/
type LoginDetails struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type Session struct {
	Blah       int
	UID        string
	SessionKey string
	Expires    int
}
