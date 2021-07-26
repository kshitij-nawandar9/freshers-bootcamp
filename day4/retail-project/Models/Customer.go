//Models/User.go
package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Config"
)

func AddCustomer(customer *Customer) (err error) {
	if err = Config.DB.Table("customers").Create(customer).Error; err != nil {
		return err
	}
	return nil
}

func GetCustomerByID(customer *Customer, id string) (err error) {
	if err = Config.DB.Table("customers").Where("id = ?", id).First(customer).Error; err != nil {
		return err
	}
	return nil
}

func GetAllCustomers(customer *[]Customer) (err error) {
	if err = Config.DB.Table("customers").Find(customer).Error; err != nil {
		return err
	}
	return nil
}
func UpdateCustomer(customer *Customer, id string) (err error) {
	fmt.Println(customer)
	Config.DB.Table("customers").Save(customer)
	return nil
}

func DeleteCustomer(customer *Customer, id string) (err error) {
	Config.DB.Table("customers").Where("id = ?", id).Delete(customer)
	return nil
}
