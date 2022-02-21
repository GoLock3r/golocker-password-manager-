// This code was written by Hunter Holton for the Project GoLock3r intended purposes of this code are for the password
//generation for the coming password manager integration. All of the code is subject to change based on user needs
// Code produces passwords containing letters of both upper and lowercase chunks of 4 seperated by dashes
// User has the option to have password contain special charaters including !@#$%^&*
// Inclusivity of special characters depends on the websites that is used for what can be included
// further changes to what characters can be included will grow depending on what websites allow

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// user input for go help https://www.geeksforgeeks.org/how-to-take-input-from-the-user-in-golang/
	println("Hello and welcom to GoLock3r's simple password generator! \nThis password generator is the first iteration of GoLock3r's password generator it \nallows you to create a password of length 8 to length 16. \nIt also alows you to choose if you want specaisl character (more to come). \n\n\n\n\n ")
	println("Enter the length of the password you would like! input needs to be a number between 8 and 16")
	var num = 0
	fmt.Scanln(&num)
	var yN = ""
	var allow = false
	println("Do you want your password to include special charaters and numbers")
	fmt.Scanln(&yN)
	if yN == "Yes" || yN == "Y" || yN == "yes" || yN == "y" {
		allow = true
	} else {
		allow = false
	}
	for true { // this for loop runs off the true boolean it is meant to act as a while loop of the same idea as while true
		if num > 16 || num < 8 {
			println("The number you entered doesnt meet our requierments please enter a number between 8 and 16")
			fmt.Scanln(&num)
		}
		if num <= 16 && num >= 8 { // gaurentees that correct values sit between 8 and 16 this can be changed in the future depending on need for stronger passwords
			break // breaks out of the defined infinite loop
		}
	}
	var pw = passwordgen(num, allow)

	println("\n" + pw) // outputs password that was created at random
	println("Thanks for using our password generator Version 0.1.1")
}

func passwordgen(n int, allow bool) string {
	var password = ""
	var letters = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var lettersAndSpecialAndNumbers = [48]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "=", "+", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	rand.Seed(time.Now().UTC().UnixNano()) // didnt know that i had to seed found an i googled issues found this https://stackoverflow.com/questions/12321133/how-to-properly-seed-random-number-generator

	if allow != true {
		for i := 0; i <= n; i++ {
			var uplow = 0
			var temp = ""

			uplow = rand.Int()
			if i%4 == 0 && i != n && i != 0 {

				if uplow%2 == 0 {
					password += letters[rand.Intn(26)] //showed me how to use rand.Int https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
					password += "-"
				}
				if uplow%2 != 0 {
					temp = letters[rand.Intn(26)]
					temp = strings.ToUpper(temp)
					password += temp
					password += "-"
				}

			}
			if i%4 != 0 || i == n {
				if uplow%2 == 0 {
					password += letters[rand.Intn(26)] //showed me how to use rand.Int https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
				}
				if uplow%2 != 0 {
					temp = letters[rand.Intn(26)]
					temp = strings.ToUpper(temp)
					password += temp
				}
			}
		}
		return password
	} else {
		for i := 0; i <= n; i++ {
			var uplow = 0
			var temp = ""

			uplow = rand.Int()
			if i%4 == 0 && i != n && i != 0 {

				if uplow%2 == 0 {
					password += lettersAndSpecialAndNumbers[rand.Intn(48)] //showed me how to use rand.Int https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
					password += "-"
				}
				if uplow%2 != 0 {
					temp = lettersAndSpecialAndNumbers[rand.Intn(48)]
					temp = strings.ToUpper(temp)
					password += temp
					password += "-"
				}

			}
			if i%4 != 0 || i == n {
				if uplow%2 == 0 {
					password += lettersAndSpecialAndNumbers[rand.Intn(48)] //showed me how to use rand.Int https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
				}
				if uplow%2 != 0 {
					temp = lettersAndSpecialAndNumbers[rand.Intn(48)]
					temp = strings.ToUpper(temp)
					password += temp
				}
			}
		}
		return password

	}
}
