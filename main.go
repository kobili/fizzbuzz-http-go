package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from fizzbuzz")
}

func fizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	queryparams := r.URL.Query()

	n := queryparams.Get("n")

	number, err := strconv.Atoi(n)
	if err != nil {
		fmt.Fprintf(w, "Failed to parse %s into an integer", n)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var value string

	if number%3 == 0 && number%5 == 0 {
		value = "FizzBuzz"
	} else if number%3 == 0 {
		value = "Fizz"
	} else if number%5 == 0 {
		value = "Buzz"
	} else {
		value = n
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"value": value})
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/fizzbuzz", fizzBuzzHandler)

	fmt.Println("Server listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
