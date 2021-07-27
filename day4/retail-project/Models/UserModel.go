////Models/UserModel.go
package Models
//
//type Product struct {
//	Id      uint   `json:"id"`
//	Name    string `json:"name"`
//	Price   uint `json:"price"`
//	Quantity   uint `json:"quantity"`
//}
//
//type Order struct {
//	//gorm.Model
//	Id      uint   `json:"id",gorm:"primaryKey"`
//	ProductID    uint `json:"product_id"`
//	CustomerID    uint `json:"customer_id"`
//	Quantity   uint `json:"quantity"`
//	Status string `json:"status"`
//
//}
//
//type Customer struct {
//	//gorm.Model
//	Id    uint `json:"id",gorm:"primaryKey"`
//	Name string `json:"name"`
//	Mobile uint64 `json:"mobile"`
//	Orders []Order `gorm:"foreignKey:Id"`
//	//Orders []Order `gorm:"foreignKey:Id,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
//}
//
//
////func (b *Product) TableName() string {
////	return "user"
////}