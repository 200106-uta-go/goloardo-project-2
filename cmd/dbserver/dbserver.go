package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/CodeZipline/project-2/pkg/dbutil"
	badger "github.com/dgraph-io/badger"
)

// Opts defines settings for the database structure
type Opts struct {
	Directory string `json:"DataDirectory"`
}

func main() {
	fmt.Println("8081")
	// Opening our configuration jsonFile.
	jsonFile, err := os.Open("./config.json")
	// If we os.Open returns an error then handle it.
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Successfully Opened config.json")
	// Defer the closing of our jsonFile so that we can parse it later on.
	defer jsonFile.Close()

	// Read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// We initialize our setting datastruct.
	var opts Opts

	// We unmarshal our byteArray which contains our
	// jsonFile's content into 'opts' which we defined above
	json.Unmarshal(byteValue, &opts)

	// Open a badger database with the defined directory from 'opts'
	db, err := badger.Open(badger.DefaultOptions(opts.Directory))
	if err != nil {
		log.Panic(err)
	}
	// Defer the closing of our database so that we can acess it later on.
	defer db.Close()

	fmt.Printf("Hosting badger database server on port: %s. \n", "8081")

	// Create a multiplexer to host mutliple endpoints in one struct
	mux1 := http.NewServeMux()
	// Index is defaulted to respond with the port.
	mux1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "port: %s", "8081")
	})
	// Calls a read command to the database.
	mux1.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Read_testing from port: %s \n", "8081")
		// Provide the query string for key="key".?key="jey"
		key := r.FormValue("key")
		value := dbutil.DbRead(db, key)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, key, value)
	})
	// Calls a write command to the database.
	mux1.HandleFunc("/write", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Write_test: %s \n", "8081")
		// Provide the query string for key="key"&value="value".
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
		errorChan <- http.ListenAndServe(":"+"8081", mux1)
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
