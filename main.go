package main

import (
    "encoding/json"
    "net/http"
)

type Bus struct {
    Route      string `json:"route"`
    Source     string `json:"source"`
    Destination string `json:"destination"`
}


func main() {
    http.HandleFunc("/getBuses", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
	
		source := r.FormValue("source")
		destination := r.FormValue("destination")
	
		// Handle the case when source and destination are empty (e.g., when you first load the page).
		if source == "" || destination == "" {
			// Return an empty response or an appropriate message.
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	
		// Your existing logic to find buses based on source and destination.
		buses := []Bus{
			{Route: "101", Source: source, Destination: destination},
			{Route: "102", Source: source, Destination: destination},
		}
	
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(buses)
	})
	
    http.Handle("/", http.FileServer(http.Dir("static")))
    http.ListenAndServe(":8080", nil)
}
