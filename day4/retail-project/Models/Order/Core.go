package Order

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Config"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Models/Customer"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Models/Product"
	"strconv"
	"time"
)

//var SleepingTime = time.Duration(15*1000000000).Seconds()
//var BufferLength int =100
var orderChannel = make(chan *TableStruct,Config.BufferLength)

//workers which continuously listen to channels and execute orders as soon as the waiting period is over
func Worker (){
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
				errProductID:=Product.GetProductByID(product,productID)
				if errProductID!=nil{
					return
				}
				if product.Quantity<order.Quantity{
					order.Status="Failed"
				}else{
					order.Status="executed"
					product.Quantity-=order.Quantity
					err := Product.UpdateProduct(product)
					if err != nil {
						return
					}
				}
				err := UpdateOrder(order)
				if err != nil {
					return
				}
				customer.UpdatedAt=time.Now()
				errUpdateCustomer := Customer.UpdateCustomer(customer)
				if errUpdateCustomer != nil {
					return
				}
			}
		}
	}
}
func AddOrder(order *TableStruct) (err error) {
	order.Status="order placed"
	if err = Config.DB.Table("orders").Create(order).Error; err != nil {
		return err
	}
	orderChannel<-order
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
func GetOrderByCustomer(order *[]TableStruct, id string) (err error) {
	if err = Config.DB.Table("orders").Where("customer_id = ?", id).Find(order).Error; err != nil {
		return err
	}
	return nil
}