package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/icholy/vedis"
)

var db = flag.String("db", ":mem:", "The database file to use.")

func main() {
	flag.Parse()

	// Open a vedis database
	v, err := vedis.Open(*db)
	if err != nil {
		fmt.Println("Error opening Vedis database.", err)
		os.Exit(1)
	}
	defer v.Close()

	// Create a new logger
	l := log.New(os.Stdout, "", 0)

	http.Handle("/", storeHandler(v, l))

	addr := getAddr()

	l.Printf("Listening on http://0.0.0.0%s", addr)

	// Start the server
	http.ListenAndServe(addr, nil)
}

func storeHandler(v *vedis.Store, l *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := []byte(r.URL.Path)

		if r.Method == "GET" {
			val, err := v.KvFetch(key)

			if err != nil {
				l.Printf("Error: %s", err)
			} else {
				w.Write(val)
			}
		} else if r.Method == "POST" {
			val, err := ioutil.ReadAll(r.Body)

			if err != nil {
				l.Printf("Error: %s", err)
			} else {
				if err := v.KvStore(key, val); err != nil {
					l.Printf("Error: %s", err)
				} else if err := v.Commit(); err == nil {
					l.Printf("Stored: '%s' in %s", val, key)
				}
			}
		}
	})
}

func getAddr() string {
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}

	return ":6600"
}
