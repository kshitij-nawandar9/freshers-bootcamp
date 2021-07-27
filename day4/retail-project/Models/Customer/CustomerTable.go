package Customer

import "github.com/jinzhu/gorm"

type TableStruct struct {
	gorm.Model
	//Id    uint `json:"id",gorm:"primaryKey"`
	Name   string `json:"name"`
	Mobile uint64 `json:"mobile"`
}


func (b *TableStruct) TableName() string {
	return "customers"
}
