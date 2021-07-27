package Product

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Config"
	"sync"
)

var mutex = &sync.Mutex{}

func AddProduct(product *TableStruct) (err error) {
	mutex.Lock()
	if err = Config.DB.Table("products").Create(product).Error; err != nil {
		return err
	}
	mutex.Unlock()
	return nil
}

func GetProductByID(product *TableStruct, id string) (err error) {
	mutex.Lock()
	if err = Config.DB.Table("products").Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	mutex.Unlock()
	return nil
}

func UpdateProduct(product *TableStruct) (err error) {
	mutex.Lock()
	Config.DB.Save(product)
	mutex.Unlock()
	return nil
}

func GetAllProducts(product *[]TableStruct) (err error) {
	mutex.Lock()
	if err = Config.DB.Table("products").Find(product).Error; err != nil {
		return err
	}
	mutex.Unlock()
	return nil
}

func DeleteProduct(product *TableStruct, id string) (err error) {
	mutex.Lock()
	Config.DB.Table("products").Where("id = ?", id).Delete(product)
	mutex.Unlock()
	return nil
}
