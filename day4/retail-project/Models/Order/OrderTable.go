package Order

type TableStruct struct {
	//gorm.Model
	Id         uint   `json:"id",gorm:"primaryKey"`
	ProductID  uint   `json:"product_id"`
	CustomerID uint   `json:"customer_id"`
	Quantity   uint   `json:"quantity"`
	Status     string `json:"status"`
}

func (b *TableStruct) TableName() string {
	return "orders"
}
