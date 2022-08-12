package main

import (
	"fmt"

	"github.com/ricardas/example-go/libs/functions/config"
	"github.com/ricardas/example-go/libs/functions/datetime"
	"github.com/ricardas/example-go/libs/users"
)

func main() {
	user1 := users.User{Name: "Maria", Id: 1}
	user2 := users.User{Name: "Maria", Id: 1}

	user1.SubChargedAt = datetime.CurrentTime()
	user2.SubChargedAt = datetime.CurrentTime()

	fmt.Printf("Env user: %s\n", config.GetEnvUser())
}
