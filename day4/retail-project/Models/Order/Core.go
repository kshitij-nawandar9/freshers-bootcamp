package Order

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Config"
	"time"
)

var timeStampMap = make(map[uint]int64)

func executeOrder(order *TableStruct){
	LastExecutionTime,present:= timeStampMap[order.CustomerID]
	if present{
		if waitingTime:=20-time.Now().Unix()+LastExecutionTime; present && waitingTime>0{
			fmt.Println("Time to sleep !")
			fmt.Println(waitingTime)
			time.Sleep(time.Duration(1000000000*waitingTime))
		}
	}
	timeStampMap[order.CustomerID]=time.Now().Unix()
	order.Status="processed"
	UpdateOrder(order)

}
func AddOrder(order *TableStruct) (err error) {
	if err = Config.DB.Table("orders").Create(order).Error; err != nil {
		return err
	}

	go executeOrder(order)

	return nil
}

func GetOrderByID(order *TableStruct, id string) (err error) {
	if err = Config.DB.Table("orders").Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}

func GetAllOrders(order *[]TableStruct) (err error) {
	if err = Config.DB.Table("orders").Find(order).Error; err != nil {
		return err
	}
	return nil
}
func UpdateOrder(order *TableStruct) (err error) {
	fmt.Println(order)
	Config.DB.Table("orders").Save(order)
	return nil
}

func DeleteOrder(order *TableStruct, id string) (err error) {
	Config.DB.Table("orders").Where("id = ?", id).Delete(order)
	return nil
}
