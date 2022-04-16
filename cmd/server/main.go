package main

import (
	"context"
	"fmt"

	"github.com/JosephO1992/go-rest-api-v2/inertnal/db"
)

// Run - is going to be responsible for the instantiation of the server
// and the start of our application.
func Run() error {
	fmt.Println("Starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to the database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("successfullt connected and pinged the database")
	return nil
}

func main() {
	fmt.Println("Go Rest API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
