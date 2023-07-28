package main

import (
	"fmt"
	"go-games-api/initializers"
)

func init() {
	initializers.LoadEnvironment()
	initializers.DatabaseConnect()
}

func main() {
	fmt.Println("main")
}
