// This code was written by Hunter Holton for the Project GoLock3r intended purposes of this code are for the password
//generation for the coming password manager integration. All of the code is subject to change based on user needs
// Code produces passwords containing letters of both upper and lowercase chunks of 4 seperated by dashes
// User has the option to have password contain special charaters including !@#$%^&*
// Inclusivity of special characters depends on the websites that is used for what can be included
// further changes to what characters can be included will grow depending on what websites allow

package generator

import (
	// "fmt"
	"golock3r/server/logger"
	"math/rand"
	"strings"
	"time"
)

var Loggers *logger.Loggers
//Generates a password based on  the number of characters you want plus the amount of dashes not specified by user as well as asks user if they would like
//to use special characters or not in the password
func GenPassword(n int, allow bool) string { // n is length of password without dashes expect password length to add anywhere from 0-3 dashes depending on length
	var password = ""
	var letters = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var lettersAndSpecialAndNumbers = [48]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "=", "+", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	rand.Seed(time.Now().UTC().UnixNano()) // didnt know that i had to seed found an i googled issues found this https://stackoverflow.com/questions/12321133/how-to-properly-seed-random-number-generator

	if n <= 16 && n > 0 {
		if !allow {
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
	} else {
		Loggers.LogWarning.Println("please enter a number between 1 and 16")
		return ""
	}
}
