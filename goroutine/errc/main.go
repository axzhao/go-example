package main

import (
	"fmt"
	"net/http"
	"os"
)

// lsof -i :5555

func serve() error {
	errc := make(chan error, 3)
	go func() {
		err := http.ListenAndServe("localhost:5555", nil)
		errc <- fmt.Errorf("listening on failed: %v", err)
	}()

	return <-errc
}

func main() {
	if err := serve(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
