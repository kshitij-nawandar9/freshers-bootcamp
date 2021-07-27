package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Models/Customer"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Models/Order"
	"net/http"
)

func AddCustomer(c *gin.Context) {
	var customer Customer.TableStruct
	err := c.BindJSON(&customer)
	if err != nil {
		return
	}
	errAddCustomer := Customer.AddCustomer(&customer)
	if errAddCustomer != nil {
		fmt.Println(errAddCustomer.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}

func GetCustomer(c *gin.Context) {
	var customer Customer.TableStruct
	id := c.Params.ByName("id")
	err := Customer.GetCustomerByID(&customer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}

func GetAllCustomers(c *gin.Context) {
	var customer []Customer.TableStruct
	err := Customer.GetAllCustomers(&customer)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}


func UpdateCustomer(c *gin.Context) {
	var customer Customer.TableStruct
	id := c.Params.ByName("id")
	errGetCustomer := Customer.GetCustomerByID(&customer, id)
	if errGetCustomer != nil {
		c.JSON(http.StatusNotFound, customer)
	}
	err := c.BindJSON(&customer)
	if err != nil {
		return
	}
	errUpdateCustomer := Customer.UpdateCustomer(&customer)
	if errUpdateCustomer != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}

func DeleteCustomer(c *gin.Context) {
	var customer Customer.TableStruct
	id := c.Params.ByName("id")
	err := Customer.DeleteCustomer(&customer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

func HistoryCustomer(c *gin.Context) {
	var customer Customer.TableStruct
	id := c.Params.ByName("id")
	err := Customer.GetCustomerByID(&customer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var orders []Order.TableStruct
		err2 := Order.GetOrderByCustomer(&orders, id)
		if err2 != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, orders)
		}
	}
}
