package db

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/mhl787156/seahorse_server/models"
)

// WriteCustomer inserts a new Customer into the database.
// Customer is inserted into the customers collection
func (c *DbConn) WriteCustomer(customer models.Customer) error {
	collection := c.DB.C(models.CollectionCustomers)
	return collection.Insert(customer)
}

type result []struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Name     string `json:"name" bson:"name"`
	PostCode string `json:"postcode" bson:"postcode"`
}

// GetCustomerList retrieves a list of all customer Ids
// Returns Customer, or an error
func (c *DbConn) GetCustomerList(n int, name string, postcode string) ([]models.ListReturn, error) {

	var result []models.ListReturn

	query := bson.M{"$or": []interface{}{
		bson.M{"firstname": bson.M{"$regex": name}},
		bson.M{"surname": bson.M{"$regex": name}},
		bson.M{"postcode": bson.M{"$regex": postcode}},
	}}

	sel := bson.M{"id": 1, "firstname": 1, "surname": 1, "postcode": 1}

	if name == "" || postcode == "" {
		query = nil
	}

	collection := c.DB.C(models.CollectionCustomers)
	err := collection.Find(query).Select(sel).All(&result)

	fmt.Println(name, postcode)

	if err != nil {
		return nil, err
	}
	for _, item := range result {
		fmt.Println(item)
	}
	return result, nil
}

// GetCustomerByID looks in the database by uuid.
// Returns the Customer, a bool to check whether any Customer was found,
// or whether an Error was returned
func (c *DbConn) GetCustomerByID(id string) (*models.Customer, bool, error) {
	collection := c.DB.C(models.CollectionCustomers)
	query := collection.FindId(id)

	//Check that a result was found
	if n, _ := query.Count(); n == 0 {
		//No results were found
		return nil, false, nil
	}
	foundCustomer := models.Customer{}
	err := query.One(&foundCustomer)
	if err != nil {
		return nil, false, err
	}
	return &foundCustomer, true, nil
}

// SearchCustomers looks in the database by queryString and returns a maximum of limit number of entries.
// The query string will most likely be a form of bson.M{"k1":"v1", "k2":"v2"...}.
// Function returns the list of Customers found, a check bool to see if the query was successful and any error.
func (c *DbConn) SearchCustomers(queryStr interface{}, limit int) ([]models.Customer, bool, error) {
	if limit == 0 {
		return nil, false, nil
	}
	collection := c.DB.C(models.CollectionCustomers)
	query := collection.Find(queryStr).Limit(limit)

	//Check that a result was found
	if n, _ := query.Count(); n == 0 {
		//No results were found
		return nil, false, nil
	}
	var err error
	var foundCustomer []models.Customer
	query.All(&foundCustomer)
	if err != nil {
		return nil, false, err
	}
	return foundCustomer, true, nil
}

func (c *DbConn) ModifyCustomer(customer *models.Customer) error {
	collection := c.DB.C(models.CollectionCustomers)
	return collection.UpdateId(customer.ID, customer)
}

// AddCustomerOrders appends an orderid to the Customers list of orders
func (c *DbConn) AddCustomerOrders(id string, orderid string) (bool, error) {
	return false, nil
}
