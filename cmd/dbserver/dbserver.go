package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/200106-uta-go/goloardo-project-2/config"
	"github.com/200106-uta-go/goloardo-project-2/pkg/dbutil"
	"github.com/dgraph-io/badger"
)

// Opts defines settings for the database structure
type Opts struct {
	Directory string `json:"DataDirectory"`
}

func main() {
	// Open a badger database with the defined directory from 'opts'
	db, err := badger.Open(badger.DefaultOptions("./app/badger"))
	if err != nil {
		panic(err)
	}
	// Defer the closing of our database so that we can acess it later on.
	defer db.Close()

	config.SendNotify("db")
	fmt.Println("Hosting badger database server on port: 8081")

	// Create a multiplexer to host mutliple endpoints in one struct
	mux1 := http.NewServeMux()
	// Index is defaulted to respond with the port.
	mux1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "port: %s", "8081")
	})
	// Calls a read command to the database.
	mux1.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		value := dbutil.DbRead(db, key)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, strings.Trim(value, "\""))
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
		errorChan <- http.ListenAndServe(":8081", mux1)
	}()

	// Used to capture the sigint(ctr+c) to print before exiting
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	for {
		select {
		case err := <-errorChan:
			if err != nil {
				config.SendDestroy("db")
				panic(err)
			}

		case sig := <-signalChan:
			fmt.Println("\nShutting down due to", sig)
			config.SendDestroy("db")
			os.Exit(0)
		}
	}
}
