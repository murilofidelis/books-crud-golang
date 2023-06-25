package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()
	fmt.Println("Runnig API...")
	fmt.Println("Port: ", config.Port)
	routers := router.GenerateRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), routers))
}
