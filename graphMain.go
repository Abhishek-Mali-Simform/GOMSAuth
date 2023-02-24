package main

import (
	"MSTRY4/lib"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Go Graph App-Only Tutorial")
	fmt.Println()

	//Load .env files
	err := godotenv.Load()
	lib.CheckError("Error Loading .env ", err)
	graphHelper := lib.NewGraphHelper()
	initializeGraph(graphHelper)
	fmt.Println("Success Login")
	user, err := graphHelper.GetUser()
	lib.CheckError("Error Getting User: ", err)
	fmt.Println(*user.GetDisplayName())
	token, err := graphHelper.GetUserToken()
	lib.CheckError("Error Getting Token: ", err)
	fmt.Println(*token)
}

func initializeGraph(graphHelper *lib.GraphHelper) {
	err := graphHelper.InitializeGraphForAppAuth()
	lib.CheckError("Error initializing Graph for app auth: ", err)
}
