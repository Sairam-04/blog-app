package user

import (
	"fmt"
	"net/http"
)

type User struct{}

func (u *User) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Registering a User..")
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logging in a User..")
}
