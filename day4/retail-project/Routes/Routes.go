//Routes/Routes.go
package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Controllers"
)
//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/product")
	{
		grp1.POST("", Controllers.AddProduct)
		grp1.GET("/:id", Controllers.GetProduct)
		grp1.GET("", Controllers.GetAllProducts)
		grp1.PATCH("/:id", Controllers.UpdateProduct)
		grp1.DELETE("/:id", Controllers.DeleteProduct)
	}
	grp2 := r.Group("/order")
	{
		grp2.POST("", Controllers.AddOrder)
		grp2.GET("/:id", Controllers.GetOrder)
		grp2.GET("", Controllers.GetAllOrders)
		grp2.PATCH("/:id", Controllers.UpdateOrder)
		grp2.DELETE("/:id", Controllers.DeleteOrder)
	}
	return r
}