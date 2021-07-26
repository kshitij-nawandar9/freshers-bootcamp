//main.go
package main
import (
	"github.com/kshitij-nawandar9/freshers-bootcamp/day3/exercise1/Config"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day3/exercise1/Models"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day3/exercise1/Routes"
	"fmt"
	"github.com/jinzhu/gorm"
)
var err error
func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}