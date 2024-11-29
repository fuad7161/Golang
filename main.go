package main

import (
	"fmt"
	"time"
)

/*
If I give the function name as a capital latter then I will act like public function
but If I give a function name as a small letter then I will act like private function
*/
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func (m *User) printFirstName() string {
	return m.FirstName
}
func (m *User) setFirstName(s string) {
	m.FirstName = s
}
func main() {
	user := User{
		FirstName:   "fuadul",
		LastName:    "hasan",
		PhoneNumber: "123456",
		Age:         25,
		BirthDate:   time.Now(),
	}
	fmt.Println(user.printFirstName())
	user.setFirstName("fuad")
	fmt.Println(user.printFirstName())
}
