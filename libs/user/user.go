package user

import "fmt"

type User struct {
	Id           int
	Name         string
	SubChargedAt string
}

func (u User) Show() {
	fmt.Printf("User: %s, id: %d\n", u.Name, u.Id)
}
