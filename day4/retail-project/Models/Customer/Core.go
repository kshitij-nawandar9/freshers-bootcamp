package Customer

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Config"
	"time"
)

func AddCustomer(customer *TableStruct) (err error) {
	customer.CreatedAt=time.Now()
	fmt.Println("customer is getting added")
	time.Sleep(30*time.Second)
	if err = Config.DB.Table("customers").Create(customer).Error; err != nil {
		return err
	}
	return nil
}

func GetCustomerByID(customer *TableStruct, id string) (err error) {
	if err = Config.DB.Table("customers").Where("id = ?", id).First(customer).Error; err != nil {
		return err
	}
	return nil
}

func GetAllCustomers(customer *[]TableStruct) (err error) {
	if err = Config.DB.Table("customers").Find(customer).Error; err != nil {
		return err
	}
	return nil
}
func UpdateCustomer(customer *TableStruct) (err error) {
	Config.DB.Table("customers").Save(customer)
	return nil
}

func DeleteCustomer(customer *TableStruct, id string) (err error) {
	Config.DB.Table("customers").Where("id = ?", id).Delete(customer)
	return nil
}
