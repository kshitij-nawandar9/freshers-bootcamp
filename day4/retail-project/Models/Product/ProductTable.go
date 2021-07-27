package Product

import "github.com/jinzhu/gorm"

type TableStruct struct {
	gorm.Model
	//Id       uint   `json:"id"`
	Name     string `json:"name"`
	Price    uint   `json:"price"`
	Quantity uint   `json:"quantity"`
}


func (b *TableStruct) TableName() string {
	return "products"
}
