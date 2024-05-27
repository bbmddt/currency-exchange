package main

import (
	"currency_exchange/handler"
	"currency_exchange/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	filePath := "exchange_rates.json"

	exchangeService, err := service.NewCurrencyExchangeService(filePath)
	if err != nil {
		panic(err)
	}

	handler := handler.NewHandler(exchangeService)

	r := mux.NewRouter()
	r.HandleFunc("/convert", handler.Convert).Methods("GET")
	r.HandleFunc("/test", handler.Test).Methods("GET")

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)

}
