package main

import (
	"fmt"
	"learn-echo-fundamental/basic"
	"learn-echo-fundamental/jwt"
)

func main() {
	var choice int
	fmt.Println("=== Learn Echo ===")
	fmt.Println("1. Run Basic")
	fmt.Println("2. Run JWT")

	fmt.Print("Choose: ")
	fmt.Scan(&choice)

	switch choice {
		case 1:
			basic.RunBasic()
		case 2:
			jwt.RunJwt()
		default:
			fmt.Println("Thank you...")
			return
	}
}