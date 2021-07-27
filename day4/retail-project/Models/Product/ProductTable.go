package Product

type TableStruct struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Price   uint `json:"price"`
	Quantity   uint `json:"quantity"`
}


func (b *TableStruct) TableName() string {
	return "products"
}
