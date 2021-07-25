//main.go
package main
import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day3/first-api/Config"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day3/first-api/Models"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day3/first-api/Routes"
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