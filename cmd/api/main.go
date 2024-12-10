package main

import (
	"flag"
	"fmt"
	"go-blog-api/database"
	"go-blog-api/internal/app"
	"log"
	"os"
)

func main() {
	runSeeder()

	fmt.Println("hello, I am GO")
	server := app.NewServer()

	fmt.Println("Server address: ", server.Addr)
	err := server.ListenAndServe()

	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}

func runSeeder() {
	// Create a flag to trigger database seeding
	seedFlag := flag.Bool("seed", false, "Seed the database with sample data")
	flag.Parse()

	// If the -seed flag is passed, run the Seed function
	if *seedFlag {
		database.Seed()
		log.Println("Database seeding completed successfully.")

		os.Exit(0)
		return
	}

}
