package main

import (
	"fmt"
)

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a value receiver.
func (u user) notify() {
	fmt.Printf("Sending User Email To %s  <%s>\n",
		u.name,
		u.email)
}

// changeEmail implements a method with a pointer receiver.
func (u *user) changeEmail(email string) {
	u.email = email
}

// main is the entry point for the application.
func main() {
	// Values of type user can be used to call methods
	// declared with a value receiver.
	mp := user{"MP", "manju@email.com"}
	mp.notify()
	// Pointers of type user can also be used to call methods
	// declared with a value receiver.
	Dell := &user{"Dell", "dell@email.com"}
	Dell.notify()

	mp.changeEmail("manju.mpsn@gmail.com")
	mp.notify()

	Dell.changeEmail("Dell@dell.com")
	Dell.notify()
}
