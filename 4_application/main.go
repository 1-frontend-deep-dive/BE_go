package goApplication

import (
	"fmt"
	"net/http"
)

func RungAppMain() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "hello world")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf(fmt.Sprintf("Bytes written: %d", n))
	})

	http.ListenAndServe(":8080", nil)
}