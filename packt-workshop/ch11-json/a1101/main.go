package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type address struct {
	Street  string
	City    string
	State   string
	Zipcode int
}

type item struct {
	Name        string `json:"itemname"`
	Description string `json:"desc,omitempty"`
	Quantity    int    `json:"qty"`
	Price       int
}

type order struct {
	TotalPrice  int
	IsPaid      bool `json:"paid"`
	Fragile     bool `json:",omitempty"`
	OrderDetail []item
}

type customer struct {
	UserName        string
	Token, Password string  `json:"-"`
	ShipTo          address `json:"shipto"`
	PurchaseOrder   order   `json:"order"`
}

func main() {
	jsonData := []byte(` 
	{ 
	  "username" :"blackhat", 
	  "shipto": 
	  { 
		"street": "Sulphur Springs Rd", 
		"city": "Park City", 
		"state": "VA", 
		"zipcode": 12345 
	  }, 
	  "order": 
	  { 
		"paid":false, 
		"orderdetail" : 
		[{ 
		  "itemname":"A Guide to the World of zeros and ones", 
		  "desc": "book", 
		  "qty": 3, 
		  "price": 50 
		}] 
	  } 
	} 
	`)

	if !json.Valid(jsonData) {
		log.Fatal("JSON not valid")

	}
	var customer customer
	json.Unmarshal(jsonData, &customer)
	fmt.Printf("%#v\n", &customer)

	toAdd := []item{
		{
			Name:        "Shadow of the Wind",
			Quantity:    1,
			Price:       15,
			Description: "Nice book..",
		}, {
			Name:        "Little pinky",
			Quantity:    1,
			Price:       15,
			Description: "Nice book",
		},
	}
	customer.PurchaseOrder.OrderDetail = append(customer.PurchaseOrder.OrderDetail, toAdd...)
	fmt.Println(customer.PurchaseOrder.OrderDetail)

	//customer.totalPrice()
	customer.PurchaseOrder.TotalPrice = customer.totalPrice()

	c, err := json.MarshalIndent(&customer, "", " ")
	if err != nil {
		fmt.Println("Error encoding to JSON:", err)
	}
	fmt.Printf("%s\n", c)
}

func (c *customer) totalPrice() int {
	totalPrice := 0
	for _, v := range c.PurchaseOrder.OrderDetail {
		totalPrice += (v.Quantity * v.Price)
	}
	return totalPrice
}
