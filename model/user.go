package model

type User struct {
	//Id int
	Name 	string  `json:"name"`
	Surname string  `json:"surname"`
	Phone 	string	`json:"phone"`
}