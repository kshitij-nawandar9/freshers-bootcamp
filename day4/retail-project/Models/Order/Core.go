package Order

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Config"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Models/Customer"
	"strconv"
	"time"
)

var timeStampMap = make(map[uint]int64)
var orderChannel = make(chan *TableStruct)
var sleepingTime = time.Duration(15*1000000000).Seconds()
//var sleepingTime = time.Duration(1).Seconds()*30

func ExecuteOrder(order *TableStruct,customer *Customer.TableStruct,waitingTime int){
	fmt.Println("check4")
	time.Sleep(time.Second*time.Duration(waitingTime))
	order.Status="executed"
	UpdateOrder(order)
	customer.UpdatedAt=time.Now()
	Customer.UpdateCustomer(customer)
}
func ManageOrder(){
	fmt.Println("check2")
	order :=<-orderChannel
	customerID:=strconv.FormatUint(uint64(order.CustomerID), 10)
	fmt.Println(customerID)
	customer := &Customer.TableStruct{}
	err := Customer.GetCustomerByID(customer, customerID)
	if err != nil {
		return
	}
	LastOrderByCustomer:=customer.UpdatedAt
	waitingTime:=time.Now().Sub(LastOrderByCustomer).Seconds()
	fmt.Println("check3",waitingTime,sleepingTime)
	if waitingTime<sleepingTime{
		go ExecuteOrder(order,customer,int(sleepingTime-waitingTime))
		//orderChannel<-order
	}else{
		fmt.Println("check5")
		order.Status="executed"
		UpdateOrder(order)
		customer.UpdatedAt=time.Now()
		Customer.UpdateCustomer(customer)
	}
	fmt.Println("check6")
	fmt.Println("order executed at : ",customer.UpdatedAt)
}
func AddOrder(order *TableStruct) (err error) {
	fmt.Println("check1")
	if err = Config.DB.Table("orders").Create(order).Error; err != nil {
		return err
	}
	go func(order *TableStruct){
		orderChannel<-order
	}(order)
	ManageOrder()
	//orderChannel<-order

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
