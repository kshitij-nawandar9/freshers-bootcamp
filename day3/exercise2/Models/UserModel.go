//Models/UserModel.go
package Models

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	LastName    string `json:"last-name"`
	DOB string `json:"dob"`
	Address string `json:"address"`
	Subject string `json:"subject"`
	Marks uint `json:"marks"`
}
func (b *User) TableName() string {
	return "user"
}