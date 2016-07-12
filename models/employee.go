package models

/*Employee model*/
type Employee struct {
	ID        int    `json:"id" bson:"_id,omitempty"`
	Admin     bool   `json:"admin" bson:"admin"`
	DateAdded string `json:"date_added" bson:"date_added"`

	Firstname string `json:"firstname" bson:"firstname"`
	Surname   string `json:"surname" bson:"surname"`

	HomePhone   string `json:"home_phone" bson:"home_phone"`
	MobilePhone string `json:"mobile_phone" bson:"mobile_phone"`
	FaxNumber   string `json:"fax_number" bson:"fax_number"`

	Email string `json:"email" bson:"email"`

	OrderIds []string `json:"orderids" bson:"orderids"`
}

/*Session model*/
type Session struct {
	SessionKey string `json:"id" bson:"id"`
	UID        string `json:"uid" bson:"uid"`
	PermaLogin bool   `json:"permalogin" bson:"permalogin"`
	Expires    int64  `json:"expires_at" bson:"expires_at"`
}
