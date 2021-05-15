package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("kek check cheburek")

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get request")
		w.Write([]byte("Hello!"))
	})
	http.ListenAndServe(":8080", nil)
}
