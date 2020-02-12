package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/200106-uta-go/goloardo-project-2/pkg/gethoroscope"
)

func main() {

	// Regiter handler
	http.HandleFunc("/daily", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		DH := gethoroscope.GetDailyHoroscope("libra", "today")
		fmt.Fprintf(w, "%v", DH)
	})

	// Start server & Setup channels
	fmt.Println("Server is serving at port 8080...")
	errorChan := make(chan error, 2)
	go func() {
		errorChan <- http.ListenAndServe(":8080", nil)
	}()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)
	for {
		select {
		case err := <-errorChan:
			if err != nil {
				log.Fatalln(err)
			}

		case sig := <-signalChan:
			fmt.Println("\nShutting down due to", sig)
			os.Exit(0)
		}
	}
}
