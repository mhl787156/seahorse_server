package models

/*Order Model*/
type Order struct {
	ID    int `json:"id" bson:"_id,omitempty"`
	CustomerID int	`json:"customerid" bson:"customerId"`
	EmployeeID int	`json:employeeid" bson:"employeeid"`
	OrderDate  string `json:"orderdate" bson:"orderdate"`
	RefNo int    `json:"refno" bson:"refno"`

	OrderProducts []OrderProduct `json:"orderproducts" bson:"orderproducts"`

	TotalWeight float32 `json:"totalweight" bson:"totalweight"`
	DeliveryDate  string `json:"deliverydate" bson:"deliverydate"` 
	ShippingMethod int `json:"shippingmethod" bson:"shippingmethod"`
	ShippingCharge int `json:"shippingchange" bson:"shippingchange"`
	
	DeliveryNote string `json:"deliverynote" bson:"deliverynote"`
	CollectionNote string `json:"collectionnote" bson:"collectionnote"`

	TotalPrice  float32 `json:"totalprice" bson:"totalprice"`
	AmountPaid float32 `json:"amountpaid" bson:"amountpaid"`

	Status	int `json:"status" bson:"status"`
}

/*OrderProduct model, wrapper for a product*/
type OrderProduct struct {
	Product  Product `json:"product" bson:"product"`
	Quantity int     `json:"quantity" bson:"quantity"`
	Discount float32 `json:"discount" bson:"discount"`
	Weight   float32 `json:"weight" bson:"weight"`
	Price    float32 `json:"price" bson:"price"`
}
