//main.go
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Config"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Models"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Routes"
)
var err error
func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig("retailDB3")))
	//Config.DB_orders, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig("ordersDB2")))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Product{})
	Config.DB.AutoMigrate(&Models.Order{})
	Config.DB.AutoMigrate(&Models.Customer{})
	//defer Config.DB_orders.Close()
	//Config.DB_orders.AutoMigrate(&Models.Order{})

	r := Routes.SetupRouter()
	//running
	r.Run()
}