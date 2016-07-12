package models

/*Product Model*/
type Product struct {
	ID          int    `json:"id" bson:"_id,omitempty"`
	Code        string `json:"pid" bson:"pid"`
	Name        string `json:"firstname" bson:"firstname"`
	Price       int    `json:"price" bson:"price"`
	Description string `json:"description" bson:"description"`
	Weight      int    `json:"wight" bson:"weight"`

	InStock bool `json:"instock" bson:"instock"`
	Stock   int  `json:"stock" bson:"stock"`

	DateAdded string `json:"date_added" bson:"date_added"`
	Supplier  string `json:"supplier" bson:"supplier"`
}
