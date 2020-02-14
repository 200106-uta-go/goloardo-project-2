package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("web")))

	http.HandleFunc("/getdaily", func(w http.ResponseWriter, r *http.Request) {

	})

	println("Server is running on port 80")
	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal(err)
	}

}
