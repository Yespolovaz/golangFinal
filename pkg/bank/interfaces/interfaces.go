package interfaces

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Name     string
	Surname  string
}

type ResponseUser struct {
	Id       uint
	Username string
	Email    string
	Name     string
	Surname  string
	Balance  int
	Accounts []ResponseAccount
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance int
	UserId  uint
}

type ResponseAccount struct {
	Id      uint
	Name    string
	Balance int
}

type Validation struct {
	Value string
	Valid string
}

type ErrorResponse struct {
	Message string
}

type Transaction struct {
	gorm.Model
	From   uint
	To     uint
	Amount int
}

type ResponseTransaction struct {
	ID     uint
	From   uint
	To     uint
	Amount int
}
