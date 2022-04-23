package main

import (
	"fmt"

	"github.com/JosephO1992/go-rest-api-v2/inertnal/comment"
	"github.com/JosephO1992/go-rest-api-v2/inertnal/db"
	transportHttp "github.com/JosephO1992/go-rest-api-v2/inertnal/transport/http"
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

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go Rest API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
