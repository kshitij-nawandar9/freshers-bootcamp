//main.go
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Config"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Models/Customer"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Models/Order"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Models/Product"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Routes"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Setup"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig("retailDB")))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Product.TableStruct{})
	Config.DB.AutoMigrate(&Order.TableStruct{})
	Config.DB.AutoMigrate(&Customer.TableStruct{})

	Setup.AlertWorkers()
	//for i:=1;i< Config.NumberOfWorkers;i++{
	//	go Order.Worker()
	//}
	r := Routes.SetupRouter()
	//running
	r.Run()

}