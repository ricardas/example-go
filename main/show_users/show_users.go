package main

import "github.com/ricardas/example-go/libs/user"

func main() {
	user1 := user.User{Name: "Maria", Id: 1}
	user2 := user.User{Name: "Maria", Id: 1}

	user1.Show()
	user2.Show()
}
