package Order

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Config"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Models/Customer"
	"strconv"
	"sync"
	"time"
)

var timeStampMap = make(map[uint]int64)
var orderChannel = make(chan *TableStruct,Config.BufferLength)

func Worker (){
	var mutex = &sync.Mutex{}
	for{
		//fmt.Println("worker ready, jobs left : ",len(orderChannel))
		select {
		case order:=<-orderChannel:
			//fmt.Println("check3")
			mutex.Lock()
			customerID:=strconv.FormatUint(uint64(order.CustomerID), 10)
			customer := &Customer.TableStruct{}
			err := Customer.GetCustomerByID(customer, customerID)
			if err != nil {
				return
			}
			LastOrderByCustomer:=customer.UpdatedAt
			waitingTime:=time.Now().Sub(LastOrderByCustomer).Seconds()
			if waitingTime<Config.SleepingTime{
				orderChannel<-order
				//fmt.Println("pushing order to channel",waitingTime,Config.SleepingTime,order.Id)
			}else{
				order.Status="executed"
				UpdateOrder(order)
				customer.UpdatedAt=time.Now()
				Customer.UpdateCustomer(customer)
			}
			mutex.Unlock()
		//default:
		//	time.Sleep(time.Second)
		}
	}
}
func AddOrder(order *TableStruct) (err error) {
	//fmt.Println("check1")
	if err = Config.DB.Table("orders").Create(order).Error; err != nil {
		return err
	}
	//go func(order *TableStruct){
	//	orderChannel<-order
	//}(order)
	//ManageOrder()
	orderChannel<-order
	//fmt.Println("check2")
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
	//fmt.Println(order)
	Config.DB.Table("orders").Save(order)
	return nil
}

func DeleteOrder(order *TableStruct, id string) (err error) {
	Config.DB.Table("orders").Where("id = ?", id).Delete(order)
	return nil
}
