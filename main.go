package main

import (
	"log"
	"net/http"
	"os"

	"github.com/stephenwithav/weather/client"
	"github.com/stephenwithav/weather/middleware"
)

func main() {
	appId := os.Getenv("OPENWEATHER_APPID")
	weatherClient := client.New(appId)
	mux := http.NewServeMux()
	mux.HandleFunc("/api", weatherClient.Retrieve)
	if err := http.ListenAndServe(":8080", middleware.ProtectMiddleware(mux)); err != nil {
		log.Fatal(err)
	}
}
