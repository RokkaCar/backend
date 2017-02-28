package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type coords struct {
	Lat float64
	Lng float64
}

func handler(w http.ResponseWriter, r *http.Request) {
	var c coords
	if r.Body == nil {
		http.Error(w, "Missing body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Printf("ip: %s, lat: %f, lng: %f\n", r.RemoteAddr, c.Lat, c.Lng)
}

func main() {
	http.HandleFunc("/coords/put", handler)
	http.ListenAndServe(":8080", nil)
}
