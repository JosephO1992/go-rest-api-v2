package main

import "fmt"

// Run - is going to be responsible for the instantiation of the server
// and the start of our application.
func Run() error {
	fmt.Println("Starting up our application")
	return nil
}

func main() {
	fmt.Println("Go Rest API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
