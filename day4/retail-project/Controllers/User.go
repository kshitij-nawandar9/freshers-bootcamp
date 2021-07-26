//Controllers/User.go
package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Models"
	"net/http"
)

func AddProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := Models.AddProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func UpdateProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
	}
	c.BindJSON(&product)
	err = Models.UpdateProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func GetProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func GetAllProducts(c *gin.Context) {
	var product []Models.Product
	err := Models.GetAllProducts(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func DeleteProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
func AddOrder(c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)
	order.Status="executed"
	err := Models.AddOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func GetOrder(c *gin.Context) {
	var order Models.Order
	id := c.Params.ByName("id")
	err := Models.GetOrderByID(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func GetAllOrders(c *gin.Context) {
	var order []Models.Order
	err := Models.GetAllOrders(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}


func UpdateOrder(c *gin.Context) {
	var order Models.Order
	id := c.Params.ByName("id")
	err := Models.GetOrderByID(&order, id)
	if err != nil {
		c.JSON(http.StatusNotFound, order)
	}
	c.BindJSON(&order)
	err = Models.UpdateOrder(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func DeleteOrder(c *gin.Context) {
	var order Models.Order
	id := c.Params.ByName("id")
	err := Models.DeleteOrder(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
////GetUsers ... Get all users
//func GetUsers(c *gin.Context) {
//	var user []Models.User
//	err := Models.GetAllUsers(&user)
//	if err != nil {
//		c.AbortWithStatus(http.StatusNotFound)
//	} else {
//		c.JSON(http.StatusOK, user)
//	}
//}
////CreateUser ... Create User
//func CreateUser(c *gin.Context) {
//	var user Models.User
//	c.BindJSON(&user)
//	err := Models.CreateUser(&user)
//	if err != nil {
//		fmt.Println(err.Error())
//		c.AbortWithStatus(http.StatusNotFound)
//	} else {
//		c.JSON(http.StatusOK, user)
//	}
//}
////GetUserByID ... Get the user by id
//func GetUserByID(c *gin.Context) {
//	id := c.Params.ByName("id")
//	var user Models.User
//	err := Models.GetUserByID(&user, id)
//	if err != nil {
//		c.AbortWithStatus(http.StatusNotFound)
//	} else {
//		c.JSON(http.StatusOK, user)
//	}
//}
////UpdateUser ... Update the user information
//func UpdateUser(c *gin.Context) {
//	var user Models.User
//	id := c.Params.ByName("id")
//	err := Models.GetUserByID(&user, id)
//	if err != nil {
//		c.JSON(http.StatusNotFound, user)
//	}
//	c.BindJSON(&user)
//	err = Models.UpdateUser(&user, id)
//	if err != nil {
//		c.AbortWithStatus(http.StatusNotFound)
//	} else {
//		c.JSON(http.StatusOK, user)
//	}
//}
////DeleteUser ... Delete the user
//func DeleteUser(c *gin.Context) {
//	var user Models.User
//	id := c.Params.ByName("id")
//	err := Models.DeleteUser(&user, id)
//	if err != nil {
//		c.AbortWithStatus(http.StatusNotFound)
//	} else {
//		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
//	}
//}