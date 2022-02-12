package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"golock3r/server/authtool"
)

func main() {
	var user string

	fmt.Print("Enter your username: ")
	fmt.Scanln(&user)

	fmt.Print("Enter your password: ")
	password, _ := terminal.ReadPassword(0)

	// hashUserPassword(user, string(password), 20)
	// writeFile()
	authtool.Init()
	if authtool.ValidateUser(user, string(password)) {
		fmt.Println("\nYou're in!")
	} else {
		fmt.Println("\nInvalid username / password")
	}
}
