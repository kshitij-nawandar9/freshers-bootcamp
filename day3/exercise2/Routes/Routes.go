//Routes/Routes.go
package Routes
import (
	"github.com/gin-gonic/gin"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day3/exercise2/Controllers"
)
//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", Controllers.GetUsers)
		grp1.POST("user", Controllers.CreateUser)
		grp1.GET("user/:id", Controllers.GetUserByID)
		grp1.PUT("user/:id", Controllers.UpdateUser)
		grp1.DELETE("user/:id", Controllers.DeleteUser)
	}
	return r
}