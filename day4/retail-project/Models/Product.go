//Models/User.go
package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Config"
)


func AddProduct(product *Product) (err error) {
	if err = Config.DB.Table("products").Create(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductByID(product *Product, id string) (err error) {
	if err = Config.DB.Table("products").Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(product *Product, id string) (err error) {
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}

func GetAllProducts(product *[]Product) (err error) {
	if err = Config.DB.Table("products").Find(product).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProduct(product *Product, id string) (err error) {
	Config.DB.Table("products").Where("id = ?", id).Delete(product)
	return nil
}
