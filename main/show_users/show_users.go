package main

import "github.com/ricardas/example-go/libs/users"

func main() {
	user1 := users.User{Name: "Maria", Id: 1}
	user2 := users.User{Name: "Maria", Id: 1}

	user1.Show()
	user2.Show()
}
