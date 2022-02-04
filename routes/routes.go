package routes

import "net/http"

// Routes - маршрутизатор
func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)

	return mux
}
