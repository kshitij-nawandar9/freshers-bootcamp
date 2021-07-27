package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Models/Customer"
	"net/http"
)


func AddCustomer(c *gin.Context) {
	var customer Customer.TableStruct
	c.BindJSON(&customer)
	err := Customer.AddCustomer(&customer)
	if err != nil {
		fmt.Println(err.Error())
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
	err := Customer.GetCustomerByID(&customer, id)
	if err != nil {
		c.JSON(http.StatusNotFound, customer)
	}
	c.BindJSON(&customer)
	err = Customer.UpdateCustomer(&customer, id)
	if err != nil {
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