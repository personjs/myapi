// config/config.go
package config

import (
	"fmt"
	"os"
)

func LoadConfig() {
	fmt.Println("Loading configuration...")
	os.Setenv("DB_CONNECTION", "postgres://user:password@localhost:5432/dbname")
}
