package vendor

//	Store provides an object to contain offered Items and Deals
type Store struct {
	Items []Product `json:"products"`
	Deals []Special
}

//	A Basket(n-many, n=customer) maps [ItemCodes]->{# in basket}.
type Basket map[string]int32

//	Product is a struct containing the ID, name, & cost of an offered Product
type Product struct {
	ID    string `json:"code"`
	Name  string `json:"name"`
	Price string `json:"price"` //1000 = $10.00
}

//Special is the schema associated with Specials created by The Company.
type Special struct {
	Label      string   `json:"code"`
	BuyProduct string   `json:"bp"`           //	Buy Product:	Product required for special
	BuyCount   int      `json:"bc"`           //	Buy Count:	# of item necessary to qualify
	GetProduct string   `json:"gp,omitempty"` //	Get Product:	Product received in sale; in case of modifying sale (not received product), no product.
	Discount   Discount `json:"dis"`          //	Discount:	Discount on GetP
	Limit      int      `json:"lmt,omitempty"`
}

//	if relative, expected amount in (-100 : 0) to be multiplied with original price.
//	!relative -> absolute: expected amount ADDED to original price.
type Discount struct {
	Relative  bool `json:"rel"`
	Amount    int  `json:"amt"`
	Inclusive bool `json:"all,omitempty"`
}
