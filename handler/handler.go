package handler

import (
	"currency_exchange/service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	exchangeService *service.CurrencyExchangeService
}

// NewHandler returns a new Handler object
func NewHandler(exchangeService *service.CurrencyExchangeService) *Handler {
	return &Handler{exchangeService: exchangeService}
}

// Convert handles the conversion request
func (h *Handler) Convert(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query().Get("source")
	target := r.URL.Query().Get("target")
	amount := r.URL.Query().Get("amount")

	convertedAmount, err := h.exchangeService.Convert(source, target, amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"msg":    "success",
		"amount": convertedAmount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Test func print test ok with Handler
func (h *Handler) Test(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"msg": "success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
