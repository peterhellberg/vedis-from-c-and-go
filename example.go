package main

import (
	"fmt"
	"os"

	"github.com/icholy/vedis"
)

func main() {
	store, err := vedis.Open(getDB())
	check(err, "Error opening Vedis database.")
	defer store.Close()

	// First GET foo set from C
	result, err := store.ExecResult("GET foo")
	check(err, "Could not execute result command:")

	fmt.Println(" foo:", result.String())

	// Then SET bar to 123
	err = store.Exec("SET bar 123")
	check(err, "Could not execute command:")

	result, err = store.ExecResult("GET bar")
	check(err, "Could not execute result command:")

	// Check if the result is numeric
	if result.IsNumeric() {
		fmt.Println(" bar:", result.Int())
	}
}

func getDB() string {
	if len(os.Args) == 2 {
		return os.Args[1]
	}

	// Fall back to using an in memory database
	return ":mem:"
}

func check(err error, message interface{}) {
	if err != nil {
		fmt.Println(message, err)
		os.Exit(1)
	}
}
