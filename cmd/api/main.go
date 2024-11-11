package main

import (
	"fmt"
	"go-blog-api/internal/app"
)

func main() {
	fmt.Println("hello, I am GO")

	server := app.NewServer()

	fmt.Println("Server address: ", server.Addr)

	err := server.ListenAndServe()

	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}