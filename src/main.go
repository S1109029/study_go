package main

import (
	"fmt"
	"net/http"
)

func helloHander(w http.ResponseWriter, r *http.Request) {

	a := []int{}
	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Println(len(a), cap(a))
	}
	fmt.Fprintf(w, "<h1>Hello World!</h1>")
}

func main() {
	http.HandleFunc("/", helloHander)
	http.ListenAndServe(":8080", nil)
}
