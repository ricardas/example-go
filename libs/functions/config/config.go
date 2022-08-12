package config

import "os"

func GetEnvUser() string {
	return os.Getenv("USER")
}
