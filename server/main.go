package main

import (
	"bufio"
	"fmt"
	"golock3r/server/authtool"
	"golock3r/server/db"
	"golock3r/server/logger"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	// Flow: Authenticate -> Show entries and perform queries -> Create entries with generated passwords -> Logout and iterate

	loggers := logger.CreateLoggers("logs.txt") // Instantiate Loggers

	db.Loggers = loggers
	// db.WriteEntry(db.CreateEntry("a", "b", "c", "d", "e"))

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

			db.Connect(username)

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
					fmt.Println("\nGoodbye!\n")
					iterate = false

				case "1":
					fmt.Println("\nViewing all entries\n")
					db.ReadAll()

				case "2":
					fmt.Println("\nHere's where you'd enter a title to search by\n")

				case "3":
					var title, url, username, other string = "", "", "", ""
					reader := bufio.NewReader(os.Stdin)

					fmt.Println("\nCreating an entry...")
					fmt.Print("Entry Title: ")
					title, _ = reader.ReadString('\n')
					title = strings.TrimSpace(title)

					fmt.Print("Website URL: ")
					url, _ = reader.ReadString('\n')
					url = strings.TrimSpace(url)

					fmt.Print("Notes: ")
					other, _ = reader.ReadString('\n')
					other = strings.TrimSpace(other)

					fmt.Print("Username: ")
					username, _ = reader.ReadString('\n')
					username = strings.TrimSpace(username)
					fmt.Print("Password: ")
					entry_password, _ := terminal.ReadPassword(0)

					db.WriteEntry(db.CreateEntry(url, title, username, string(entry_password), other))

					fmt.Println("\nWrote entry to database!\n")

				case "4":
					fmt.Println("\nHere's where you'd delete an entry by title\n")

				default:
					fmt.Println("\nInvalid input. Try again.\n")
				}
			}

		} else {
			fmt.Println("\nInvalid username / password")
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
