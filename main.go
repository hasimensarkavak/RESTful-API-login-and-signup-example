package main

import (
	"Login_And_SignUp_Example/database"
	"Login_And_SignUp_Example/handlers"
)

func main() {
	database.DatabaseBegin()
	handlers.HandlersRun()
}
