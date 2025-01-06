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
	// Create flags
	seedFlag := flag.Bool("seed", false, "Seed the database with sample data")
	migrateFlag := flag.Bool("migrate", false, "Migrate the database")
	flag.Parse() // Parse the flags once

	// Run seeder or migrator if flags are set
	if *seedFlag {
		runSeeder()
		os.Exit(0) // Exit after seeding
	}
	if *migrateFlag {
		runMigrater()
		os.Exit(0) // Exit after migration
	}

	// Start the server
	runServer()
}

func runSeeder() {
	database.Seed()
	log.Println("Database seeding completed successfully.")
}

func runMigrater() {
	database.Migration()
	log.Println("Database migration completed successfully.")
}

func runServer() {
	fmt.Println("hello, I am GO")
	server := app.NewServer()

	fmt.Println("Server address: ", server.Addr)
	err := server.ListenAndServe()

	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
