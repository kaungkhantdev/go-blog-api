package app

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
}


func NewServer() *http.Server {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		fmt.Println("Can't parse port")
	}
	
	nServer := &Server{
		port: port,
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", nServer.port),
		Handler:      nServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server;
}