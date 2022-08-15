package main

import (
	"fmt"

	"github.com/ricardas/example-go/libs/functions/config"
	"github.com/ricardas/example-go/libs/functions/date"
	"github.com/ricardas/example-go/libs/user"
)

func main() {
	user1 := user.User{Name: "Maria", Id: 1}
	user2 := user.User{Name: "Alexa", Id: 2}

	user1.SubChargedAt = date.CurrentTime()
	user2.SubChargedAt = date.CurrentTime()

	fmt.Printf("Env user: %s\n", config.GetEnvUser())
}
