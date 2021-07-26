//Models/User.go
package Models
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day3/exercise2/Config"
	"strconv"
	"time"
)
//GetAllUsers Fetch all user data
func GetAllUsers(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}
//CreateUser ... Insert New data
func CreateUser(user *User) (err error) {
	date,e:=time.Parse(time.RFC3339,user.DOB)
	if e!=nil{
		panic("Can't parse time format")
	}
	dateUnix:=date.Unix()
	fmt.Println(dateUnix)
	user.DOB=strconv.FormatInt(dateUnix,10)
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}
//UpdateUser ... Update user
func UpdateUser(user *User, id string) (err error) {
	date,e:=time.Parse(time.RFC3339,user.DOB)
	if e!=nil{
		panic("Can't parse time format")
	}
	dateUnix:=date.Unix()
	fmt.Println(dateUnix)
	user.DOB=strconv.FormatInt(dateUnix,10)

	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}
//DeleteUser ... Delete user
func DeleteUser(user *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}