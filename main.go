package main

import (
	"fmt"
	"net/http"
	"os"
    "os/signal"
    "syscall"
)
// import "rsc.io/quote"

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello %s welcome to the server", r.URL.Path[1:])
	// fmt.Fprintf(w, "hello %d, welcome to the Go server", r.URL.Path[1:])
	var log = r.URL.Path[1:]
	if log == "favicon.ico" {
		return;
	}
	fmt.Println(log)
	
}


func main() {
	
	http.HandleFunc("/", httpHandler)
	http.ListenAndServe(":9000", nil)

	stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

    fmt.Println("Server is running. Press CTRL+C to stop.")

    // Block until an interrupt signal is received
    <-stop

    fmt.Println("\nGracefully shutting down the server...")
    // Perform any cleanup tasks here

    fmt.Println("Server stopped.")

}






