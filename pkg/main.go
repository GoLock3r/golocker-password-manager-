package main

import (
	"bufio"
	"fmt"
	"golock3r/server/authtool"
	"golock3r/server/crypt"
	"golock3r/server/db"
	"golock3r/server/logger"
	"golock3r/server/web"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	// Flow: Authenticate -> Show entries and perform queries -> Create entries with generated passwords -> Logout and iterate
	web.Run()
	loggers := logger.CreateLoggers("logs.txt") // Instantiate Loggers

	db.Loggers = loggers
	crypt.Loggers = loggers

	authtool.Loggers = loggers
	authtool.LoginFile = "logins.txt"

	fmt.Println("Welcome to GoLock3r, extremely rudimentary CLI edition")

	var username = ""
	var input = ""

	// Uncomment to create a user account for demonstration
	authtool.CreateUser("demo", "demo123")

	for {
		fmt.Print("Please enter your username: ")
		fmt.Scanln(&username)

		fmt.Print("Please enter your password: ")
		password, _ := terminal.ReadPassword(0)

		if authtool.ValidateUser(username, string(password)) {
			key := authtool.GetKey(username, string(password))

			fmt.Println("\n\nWelcome", username)

			db.Connect(username)

			var iterate = true

			for iterate {

				fmt.Println("Type 1 to view all entries")
				fmt.Println("Type 2 to search for an entry by title")
				fmt.Println("Type 3 to search for an entry by username")
				fmt.Println("Type 4 to create an entry")
				fmt.Println("Type 5 to update an entry")
				fmt.Println("Type 6 to delete an entry")
				fmt.Println("Type 0 to logout")
				fmt.Print("> ")
				fmt.Scanln(&input)

				switch input {
				case "0":
					fmt.Println("\nGoodbye!")
					iterate = false

				case "1":
					fmt.Println("\nViewing all entries")
					for _, entry := range db.ReadAll() {
						for k, v := range db.DecryptEntry(key, entry) {
							fmt.Println(k, v)
						}
						fmt.Println("")
					}

				case "2":
					var title string = ""
					reader := bufio.NewReader(os.Stdin)
					fmt.Println("\nPlease enter the title by which you'd like to search: ")
					title, _ = reader.ReadString('\n')
					title = strings.TrimSpace(title)
					db.ReadFromTitle(title)

				case "3":
					var username string = ""
					reader := bufio.NewReader(os.Stdin)
					fmt.Println("\nPlease enter the username by which you'd like to search: ")
					username, _ = reader.ReadString('\n')
					username = strings.TrimSpace(username)
					db.ReadFromUsername(username)

				case "4":
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
					str_password := string(entry_password)

					entry := map[string]string{
						"title":         title,
						"private_url":   url,
						"private_notes": other,
						"username":      username,
						"password":      str_password,
					}

					db.WriteEntry(db.EncryptEntry(key, entry))

					fmt.Println("\nWrote entry to database!")

				case "5":
					var title string = ""
					reader := bufio.NewReader(os.Stdin)
					fmt.Println("\nPlease enter the title of the entry you'd like to update: ")
					title, _ = reader.ReadString('\n')
					title = strings.TrimSpace(title)
					//db.UpdateEntry(title) going to need to update this line at some point to reflect changes hunter made to code to get ready for imput from front end

				case "6":
					var title string = ""
					reader := bufio.NewReader(os.Stdin)
					fmt.Println("\nPlease enter the title of the entry you'd like to delete: ")
					title, _ = reader.ReadString('\n')
					title = strings.TrimSpace(title)
					db.DeleteEntry(title)

				default:
					fmt.Println("\nInvalid input. Try again.")
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
