package ordersdto


type Request_Orders struct {
	Date         string `json:"date"`
	Status       string `json:"status"`
	Product_Id   int
	Customers_Id int
}
