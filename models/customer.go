package models

const (
	// CollectionCustomers holds the name of the customers collection
	CollectionCustomers = "customers"
)

/*Customer Model*/
type Customer struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	DateAdded int64  `json:"date_added" bson:"date_added"`

	Firstname string `json:"firstname" bson:"firstname"`
	Surname   string `json:"surname" bson:"surname"`

	CompanyName string `json:"company_name" bson:"company_name"`
	Address     string `json:"address" bson:"address"`
	PostCode    string `json:"postcode" bson:"postcode"`

	HomePhone   string `json:"home_phone" bson:"home_phone"`
	MobilePhone string `json:"mobile_phone" bson:"mobile_phone"`
	FaxNumber   string `json:"fax_number" bson:"fax_number"`

	Email string `json:"email" bson:"email"`

	AdMailing bool `json:"mailing" bson:"mailing"`

	Notes string `json:"notes" bson:"notes"`

	OrderIds []string `json:"orderids" bson:"orderids"`

	Active bool `json:"active" bson:"active"`
}

/*ListReturn Model - return type when a list is requested for dynamic search*/
type ListReturn struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	FirstName string `json:"firstname" bson:"firstname"`
	SurName   string `json:"surname" bson:"surname"`
	PostCode  string `json:"postcode" bson:"postcode"`
}

/*ListFullReturn Model - return type when a list is requested for dynamic search*/
type ListFullReturn struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Date      int    `json:"date" bson: "date"`
	FirstName string `json:"firstname" bson:"firstname"`
	SurName   string `json:"surname" bson:"surname"`
	PostCode  string `json:"postcode" bson:"postcode"`
}
