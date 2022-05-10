package main

import (
	"golock3r/server/authtool"
	"golock3r/server/crypt"
	"golock3r/server/db"
	"golock3r/server/logger"
	"golock3r/server/web"
)

func main() {
	// Flow: Authenticate -> Show entries and perform queries -> Create entries with generated passwords -> Logout and iterate

	loggers := logger.CreateLoggers("logs.txt") // Instantiate Loggers
	db.Loggers = loggers
	crypt.Loggers = loggers
	authtool.Loggers = loggers
	web.Loggers = loggers

	authtool.LoginFile = "logins.txt"

	web.Run()
	
}
