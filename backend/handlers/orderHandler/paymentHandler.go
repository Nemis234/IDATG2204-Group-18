package orderHandler

import (
	"log"
	"net/http"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("PaymentHandler called with method: ", r.Method)
	switch r.Method {
	case http.MethodGet:
		http.Error(w, "Method not implemented", http.StatusNotImplemented)
	case http.MethodPost:
		http.Error(w, "Method not implemented", http.StatusNotImplemented)
	case http.MethodPut:
		http.Error(w, "Method not implemented", http.StatusNotImplemented)
	case http.MethodDelete:
		http.Error(w, "Method not implemented", http.StatusNotImplemented)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
