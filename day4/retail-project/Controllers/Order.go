package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Order"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Models/Order"
	"net/http"
)


func AddOrder(c *gin.Context) {
	var order Order.TableStruct
	c.BindJSON(&order)
	//order.Status="executed"
	err := Order.AddOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func GetOrder(c *gin.Context) {
	var order Order.TableStruct
	id := c.Params.ByName("id")
	err := Order.GetOrderByID(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func GetAllOrders(c *gin.Context) {
	var order []Order.TableStruct
	err := Order.GetAllOrders(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}


func UpdateOrder(c *gin.Context) {
	var order Order.TableStruct
	id := c.Params.ByName("id")
	err := Order.GetOrderByID(&order, id)
	if err != nil {
		c.JSON(http.StatusNotFound, order)
	}
	c.BindJSON(&order)
	err = Order.UpdateOrder(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func DeleteOrder(c *gin.Context) {
	var order Order.TableStruct
	id := c.Params.ByName("id")
	err := Order.DeleteOrder(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}