package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

func main() {
	runProxy()
}

func runProxy() {

	currTime := time.Now()
	path := "logs/proxy/" + currTime.Format("01-02-2006") + ".log"
	file, _ := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	log.SetOutput(file)

	//parse the url
	myURL, err := url.Parse("http://localhost:8081")

	if err != nil {
		log.Fatal(err)
	}
	//initiate the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(myURL)

	fmt.Println("Server is running on port 443")

	//listen on given ports
	err = http.ListenAndServeTLS(":443", "cert.pem", "key.pem", proxy)
	if err != nil {
		log.Fatal(err)
	}

}
