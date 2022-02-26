// This is simple example of http server with Go language.
// GET / - returns "Hello, World!"
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", HelloWorldHandler)
	APP_PORT := os.Getenv("APP_PORT")
	if APP_PORT == "" {
		APP_PORT = "9090"
	}
	http.ListenAndServe(":"+APP_PORT, nil)
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// log stdout access_log as json "{"method":"GET","path":"/","status":200,"size":0}"
	log.Printf("{\"method\":\"%s\",\"path\":\"%s\",\"status\":%d,\"size\":%d}\n",
		r.Method, r.URL.Path, 200, 0)
	fmt.Fprintf(w, "Hello, World!")
}
