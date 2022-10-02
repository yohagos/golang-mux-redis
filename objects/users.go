package objects

import (
	"errors"
	"log"
)

type User struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

var (
	AllUsers []User
)

func UserInit() {
	AllUsers = append(AllUsers, User{
		Name:  "Max",
		Phone: "123123123",
		Email: "Max@test.com",
	})
	AllUsers = append(AllUsers, User{
		Name:  "Eva",
		Phone: "456456456",
		Email: "Eva@test.com",
	})
	AllUsers = append(AllUsers, User{
		Name:  "Peter",
		Phone: "789789789",
		Email: "Peter@test.com",
	})
}

func GetUser(searchAfter string) (User, error){
	for _, v := range AllUsers {
		if v.Name == searchAfter {
			log.Println(v)
			return v, nil
		}
	}
	return User{}, errors.New("not found")
}

func (u *User) GetUerName() string {
	return u.Name
}

func (u *User) GetUserEmail() string {
	return u.Email
}

func (u *User) GetUserPhone() string {
	return u.Phone
}