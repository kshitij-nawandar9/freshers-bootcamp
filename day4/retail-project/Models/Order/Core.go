package Order

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Config"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Models/Customer"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Models/Product"
	"strconv"
	"sync"
	"time"
)

//var SleepingTime = time.Duration(15*1000000000).Seconds()
//var BufferLength int =100
var orderChannel = make(chan *TableStruct,Config.BufferLength)
var mutex = &sync.Mutex{}

func Worker (){
	//workers which continuously listen to channels and execute orders as soon as the waiting period is over
	for{
		select {
		case order:=<-orderChannel:
			//Calculate the time executed since the last order by the customer, and if it is less than
			//the sleeping time , push the order back into the channel
			customerID:=strconv.FormatUint(uint64(order.CustomerID), 10)
			customer := &Customer.TableStruct{}
			err := Customer.GetCustomerByID(customer, customerID)
			if err != nil {
				return
			}
			LastOrderByCustomer:=customer.UpdatedAt
			waitingTime:=time.Now().Sub(LastOrderByCustomer).Seconds()
			if waitingTime< Config.SleepingTime {
				orderChannel<-order
			}else{
				//check if the order quantity is available, if not then update order status as failed
				productID :=strconv.FormatUint(uint64(order.ProductID),10)
				product := &Product.TableStruct{}
				err:=Product.GetProductByID(product,productID)
				if err!=nil{
					return
				}
				if product.Quantity<order.Quantity{
					order.Status="Failed"
				}else{
					order.Status="executed"
					product.Quantity-=order.Quantity
					Product.UpdateProduct(product)
				}
				UpdateOrder(order)
				customer.UpdatedAt=time.Now()
				Customer.UpdateCustomer(customer)
			}
		}
	}
}
func AddOrder(order *TableStruct) (err error) {
	mutex.Lock()
	if err = Config.DB.Table("orders").Create(order).Error; err != nil {
		return err
	}
	mutex.Unlock()
	orderChannel<-order
	return nil
}

func GetOrderByID(order *TableStruct, id string) (err error) {
	mutex.Lock()
	if err = Config.DB.Table("orders").Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	mutex.Unlock()
	return nil
}

func GetAllOrders(order *[]TableStruct) (err error) {
	mutex.Lock()
	if err = Config.DB.Table("orders").Find(order).Error; err != nil {
		return err
	}
	mutex.Unlock()
	return nil
}
func UpdateOrder(order *TableStruct) (err error) {
	//fmt.Println(order)
	mutex.Lock()
	Config.DB.Table("orders").Save(order)
	mutex.Unlock()
	return nil
}

func DeleteOrder(order *TableStruct, id string) (err error) {
	mutex.Lock()
	Config.DB.Table("orders").Where("id = ?", id).Delete(order)
	mutex.Unlock()
	return nil
}
