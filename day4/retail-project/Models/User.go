//Models/User.go
package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kshitij-nawandar9/freshers-bootcamp/day4/retail-project/Config"
)


func AddProduct(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductByID(product *Product, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(product *Product, id string) (err error) {
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}

func GetAllProducts(product *[]Product) (err error) {
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProduct(product *Product, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(product)
	return nil
}

func AddOrder(order *Order) (err error) {
	if err = Config.DB_orders.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderByID(order *Order, id string) (err error) {
	if err = Config.DB_orders.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}

func GetAllOrders(order *[]Order) (err error) {
	if err = Config.DB_orders.Find(order).Error; err != nil {
		return err
	}
	return nil
}
func UpdateOrder(order *Order, id string) (err error) {
	fmt.Println(order)
	Config.DB_orders.Save(order)
	return nil
}

func DeleteOrder(order *Order, id string) (err error) {
	Config.DB_orders.Where("id = ?", id).Delete(order)
	return nil
}

////GetAllUsers Fetch all user data
//func GetAllUsers(user *[]User) (err error) {
//	if err = Config.DB.Find(user).Error; err != nil {
//		return err
//	}
//	return nil
//}
////CreateUser ... Insert New data
//func CreateUser(user *User) (err error) {
//	if err = Config.DB.Create(user).Error; err != nil {
//		return err
//	}
//	return nil
//}
////GetUserByID ... Fetch only one user by Id
//func GetUserByID(user *User, id string) (err error) {
//	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
//		return err
//	}
//	return nil
//}
////UpdateUser ... Update user
//func UpdateUser(user *User, id string) (err error) {
//	fmt.Println(user)
//	Config.DB.Save(user)
//	return nil
//}
////DeleteUser ... Delete user
//func DeleteUser(user *User, id string) (err error) {
//	Config.DB.Where("id = ?", id).Delete(user)
//	return nil
//}