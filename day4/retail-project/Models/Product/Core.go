package Product

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Config"
)


func AddProduct(product *TableStruct) (err error) {
	if err = Config.DB.Table("products").Create(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductByID(product *TableStruct, id string) (err error) {
	if err = Config.DB.Table("products").Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(product *TableStruct) (err error) {
	Config.DB.Save(product)
	return nil
}

func GetAllProducts(product *[]TableStruct) (err error) {
	if err = Config.DB.Table("products").Find(product).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProduct(product *TableStruct, id string) (err error) {
	Config.DB.Table("products").Where("id = ?", id).Delete(product)
	return nil
}
