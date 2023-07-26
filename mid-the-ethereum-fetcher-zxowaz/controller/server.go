package controller

import (
	"fmt"
	"mid-the-ethereum-fetcher-zxowaz/service"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func HandleRequests() error {
	transaction := service.NewServiceTransaction()
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(commonMiddleware)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/lime/eth", transaction.GetByRlphex).Methods("GET")
	myRouter.HandleFunc("/lime/all", transaction.GetAll).Methods("GET")

	err := http.ListenAndServe(os.Getenv("API_PORT"), myRouter)
	return err
}
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
