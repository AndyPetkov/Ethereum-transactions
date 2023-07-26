package main

import (
	"fmt"
	"mid-the-ethereum-fetcher-zxowaz/controller"
	"mid-the-ethereum-fetcher-zxowaz/loader"
	"mid-the-ethereum-fetcher-zxowaz/logger"
)

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	logger.GetInstance().InfoLogger.Println("Starting the application...")
	loader.LoadTransactions()
	controller.HandleRequests()

}
