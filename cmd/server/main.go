package main

import (
	"context"
	"fmt"

	"github.com/JosephO1992/go-rest-api-v2/inertnal/comment"
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
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)

	cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID:     "71c5d074-b6cf-11ec-b909-0242ac120002",
			Slug:   "manual-test",
			Author: "Joe",
			Body:   "Hello World!",
		},
	)
	fmt.Println(cmtService.GetComment(context.Background(),
		"71c5d074-b6cf-11ec-b909-0242ac120002",
	))

	return nil
}

func main() {
	fmt.Println("Go Rest API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
