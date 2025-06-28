package main

import (
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
	// fmt.Fprintf(w, "echo... %s", n)

	number, err := strconv.Atoi(n)
	if err != nil {
		fmt.Fprintf(w, "Failed to parse %s into an integer", n)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if number%3 == 0 && number%5 == 0 {
		fmt.Fprintf(w, "%s", "FizzBuzz")
	} else if number%3 == 0 {
		fmt.Fprintf(w, "%s", "Fizz")
	} else if number%5 == 0 {
		fmt.Fprintf(w, "%s", "Buzz")
	} else {
		fmt.Fprintf(w, "%s", n)
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/fizzbuzz", fizzBuzzHandler)

	fmt.Println("Server listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
