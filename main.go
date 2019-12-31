package main

import "fmt"

//Manager struct
type Manager struct {
	managerID  int
	managerPSW string
}

//User struct
type User struct {
	userID    int
	firstname string
	lastname  string
	address   string
	email     string
	phone     string
	psw       string
	According
}

//According struct
type According struct {
	accordingID int
	saving      float64
	checking    float64
}

func (m Manager) createUser() User {
	var u User
	fmt.Print("Enter user first name: ")
	fmt.Scan(&u.firstname)
	fmt.Print("Enter user last name ")
	fmt.Scan(&u.lastname)
	fmt.Print("Enter user address ")
	fmt.Scan(&u.address)
	fmt.Print("Enter user email ")
	fmt.Scan(&u.email)
	fmt.Print("Enter user phone ")
	fmt.Scan(&u.phone)
	fmt.Print("Create user psw ")
	fmt.Scan(&u.psw)
	return u
}

func (m Manager) createAccording() *According {
	var a According
	a.accordingID = 1111
	a.saving = 1222
	a.checking = 9000
	return &a
}

func main() {

	var m Manager
	var u User

	u = m.createUser()
	fmt.Println(u.psw)
	m.createAccording()
	fmt.Println(u.According.accordingID)
}
