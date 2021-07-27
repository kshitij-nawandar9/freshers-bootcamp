package Customer

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Config"
	"sync"
	"time"
)
var mutex = &sync.Mutex{}

func AddCustomer(customer *TableStruct) (err error) {
	customer.CreatedAt=time.Now()
	mutex.Lock()
	if err = Config.DB.Table("customers").Create(customer).Error; err != nil {
		return err
	}
	mutex.Unlock()
	return nil
}

func GetCustomerByID(customer *TableStruct, id string) (err error) {
	mutex.Lock()
	if err = Config.DB.Table("customers").Where("id = ?", id).First(customer).Error; err != nil {
		return err
	}
	mutex.Unlock()
	return nil
}

func GetAllCustomers(customer *[]TableStruct) (err error) {
	mutex.Lock()
	if err = Config.DB.Table("customers").Find(customer).Error; err != nil {
		return err
	}
	mutex.Unlock()
	return nil
}
func UpdateCustomer(customer *TableStruct) (err error) {
	mutex.Lock()
	Config.DB.Table("customers").Save(customer)
	mutex.Unlock()
	return nil
}

func DeleteCustomer(customer *TableStruct, id string) (err error) {
	mutex.Lock()
	Config.DB.Table("customers").Where("id = ?", id).Delete(customer)
	mutex.Unlock()
	return nil
}
