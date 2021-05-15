package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("kek check cheburek")

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get request")
		_, err := w.Write([]byte("Hello!"))
		if err != nil {
			fmt.Println(err)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
