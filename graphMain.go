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

}

func initializeGraph(graphHelper *lib.GraphHelper) {
	user, err := graphHelper.InitializeGraphForAppAuth()
	fmt.Println(*user.GetDisplayName())
	lib.CheckError("Error initializing Graph for app auth: ", err)
	//fmt.Printf("Hello, %s!\n", *user.GetDisplayName())
}
