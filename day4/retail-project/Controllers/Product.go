package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Models"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Models/Product"
	"net/http"
)

func AddProduct(c *gin.Context) {
	var product Product.TableStruct
	err := c.BindJSON(&product)
	if err != nil {
		return
	}
	errAddProduct := Product.AddProduct(&product)
	if errAddProduct != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func UpdateProduct(c *gin.Context) {
	var product Product.TableStruct
	id := c.Params.ByName("id")
	err := Product.GetProductByID(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
	}
	errBindJSON := c.BindJSON(&product)
	if errBindJSON != nil {
		return
	}
	errUpdateProduct := Product.UpdateProduct(&product)
	if errUpdateProduct != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func GetProduct(c *gin.Context) {
	var product Product.TableStruct
	id := c.Params.ByName("id")
	err := Product.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func GetAllProducts(c *gin.Context) {
	var product []Product.TableStruct
	err := Product.GetAllProducts(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func DeleteProduct(c *gin.Context) {
	var product Product.TableStruct
	id := c.Params.ByName("id")
	err := Product.DeleteProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
