package Customer

import "github.com/jinzhu/gorm"

type TableStruct struct {
	gorm.Model
	//Id    uint `json:"id",gorm:"primaryKey"`
	Name string `json:"name"`
	Mobile uint64 `json:"mobile"`
	//Orders []Order `gorm:"foreignKey:Id"`
	//Orders []Order `gorm:"foreignKey:Id,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}


func (b *TableStruct) TableName() string {
	return "customers"
}
