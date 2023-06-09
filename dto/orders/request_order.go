package ordersdto


type Request_Orders struct {
	Date         string `json:"date"`
	Status       string `json:"status"`
	// Products     models.Products
	Product_Id   int
	Customers_Id int
}
