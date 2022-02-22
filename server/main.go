package main

import (
	"fmt"
	"golock3r/server/authtool"
	"golock3r/server/db"
	"golock3r/server/logger"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	// Flow: Authenticate -> Show entries and perform queries -> Create entries with generated passwords

	loggers := logger.CreateLoggers("logs.txt") // Instantiate Loggers

	db.Loggers = loggers
	// db.Connect("demo")

	authtool.Loggers = loggers
	authtool.LoginFile = "logins.txt"

	fmt.Println("Welcome to GoLock3r, extremely rudimentary CLI edition\n")

	var username = ""
	var input = ""

	// Uncomment to create a user account for demonstration
	// createUser("demo", "demo123")

	for {
		fmt.Print("Please enter your username: ")
		fmt.Scanln(&username)

		fmt.Print("Please enter your password: ")
		password, _ := terminal.ReadPassword(0)

		if authtool.ValidateUser(username, string(password)) {
			fmt.Println("\n\nWelcome", username)
			var iterate = true

			for iterate {

				fmt.Println("Type 1 to view all entries")
				fmt.Println("Type 2 to search for an entry by title")
				fmt.Println("Type 3 to create an entry")
				fmt.Println("Type 4 to delete an entry")
				fmt.Println("Type 0 to logout")
				fmt.Print("> ")
				fmt.Scanln(&input)

				switch input {
				case "0":
					fmt.Println("\nGoodbye!")
					iterate = false

				case "2":
					fmt.Println("\nEntry title: ")
					fmt.Scanln(&input)

				case "3":
					fmt.Println("\nInserting an entry...")
					fmt.Print("Entry Title: ")
					fmt.Scanln(&input)
					fmt.Print("Website URL: ")
					fmt.Scanln(&input)
					fmt.Print("Username: ")
					fmt.Scanln(&input)
					fmt.Print("Password: ")
					fmt.Scanln(&input)

				case "4":
					fmt.Println("\nDelete an entry: ")

				case "1":
					fmt.Println("\nViewing all entries")

				default:
					fmt.Println("\nInvalid input")
				}
			}

		} else {
			fmt.Println("\nInvalid username / password.")
		}
		fmt.Print("Continue? y/n: ")
		fmt.Scanln(&input)

		if input == "n" {
			break
		}
	}
}

func createUser(username string, password string) {
	authtool.HashUserPassword(username, password, 12)
	authtool.WriteFile()
}
