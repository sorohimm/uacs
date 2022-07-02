package models

type User struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,required"`
	Username  string `json:"username,required"`
	Password  string `json:"password,required"`
}
