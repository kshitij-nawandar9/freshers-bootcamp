//Models/User.go
package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Config"
)

func AddOrder(order *Order) (err error) {
	if err = Config.DB_orders.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderByID(order *Order, id string) (err error) {
	if err = Config.DB_orders.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}

func GetAllOrders(order *[]Order) (err error) {
	if err = Config.DB_orders.Find(order).Error; err != nil {
		return err
	}
	return nil
}
func UpdateOrder(order *Order, id string) (err error) {
	fmt.Println(order)
	Config.DB_orders.Save(order)
	return nil
}

func DeleteOrder(order *Order, id string) (err error) {
	Config.DB_orders.Where("id = ?", id).Delete(order)
	return nil
}
