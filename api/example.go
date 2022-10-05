package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func init() {
	fmt.Println("System init start!")
	// initialize project  e.g.
	// 1.Check and get environment variable
	// 2.Initialize database connection
	// ...
}

func Run() {
	// Hello world, the web server
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		_, err := io.WriteString(w, "Hello, world!\n")
		if err != nil {
			return
		}
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
