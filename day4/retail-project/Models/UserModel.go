//Models/UserModel.go
package Models
type Product struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Price   uint `json:"price"`
	Quantity   uint `json:"quantity"`
}

type Order struct {
	Id      uint   `json:"id"`
	ProductID    uint `json:"product_id"`
	CustomerID    uint `json:"customer_id"`
	Quantity   uint `json:"quantity"`
	Status string `json:"status"`
}

func (b *Product) TableName() string {
	return "user"
}