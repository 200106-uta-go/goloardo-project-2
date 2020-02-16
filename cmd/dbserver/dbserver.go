package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/CodeZipline/project-2/pkg/dbutil"
	"github.com/dgraph-io/badger"
)

func main() {

	// Open a badger database with the defined directory from 'opts'
	db, err := badger.Open(badger.DefaultOptions("./app/badger"))
	if err != nil {
		log.Panic(err)
	}
	// Defer the closing of our database so that we can acess it later on.
	defer db.Close()

	// Create a multiplexer to host mutliple endpoints in one struct
	mux1 := http.NewServeMux()

	// Index is defaulted to respond with the port.
	mux1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "port: %s", "9090")
	})

	// Calls a read command to the database.
	mux1.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		value := dbutil.DbRead(db, key)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, key, value)
	})

	// Calls a write command to the database.
	mux1.HandleFunc("/write", func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		value := r.FormValue("value")
		rKey, rVal := dbutil.DbWrite(db, key, value)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, rKey, rVal)
	})

	// ListenAnd Serve the multiplexer functions on designated port and send the error to a capture channel
	// , the capture channel is waiting for an error to handle it below in the for loop
	errorChan := make(chan error, 5)
	go func() {
		fmt.Printf("Hosting badger database server on port: %s. \n", "9090")
		errorChan <- http.ListenAndServe(":9090", mux1)
	}()

	// Used to capture the sigint(ctr+c) to print before exiting
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	for {
		select {
		case err := <-errorChan:
			if err != nil {
				log.Panic(err)
			}

		case sig := <-signalChan:
			fmt.Println("\nShutting down due to", sig)
			os.Exit(0)
		}
	}
}
