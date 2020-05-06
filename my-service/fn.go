package hello

import (
	"fmt"
	"net/http"
)

// Helloevent : It accept events like bucket, pubsub, bigquery.
func Helloevent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello event function called", r)
	fmt.Fprintf(w, "Hello Event, Successfully called")
}

// Hellorequest : It accept http call with QueryParams
func Hellorequest(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		id = "someone"
	}
	fmt.Println("Hello http function called", r)
	fmt.Fprintf(w, "Hello Event, Successfully called for Id %s", id)
}
