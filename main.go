package main

import (
	"fmt"
	"net/http"
)
// import "rsc.io/quote"

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello %s welcome to the server", r.URL.Path[1:])
}


func main() {
	// fmt.Println(quote.Go())
	var x = 60
	var arr = []int{40, 50, x, 90}

	fmt.Println("hello go")

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	var call = fib(6)
	
	var name string = "hello"
	fmt.Println(name)

	fmt.Print(call)
	http.HandleFunc("/", httpHandler)
	http.ListenAndServe(":9000", nil)

}


func fib(n int) int {
	if n==0 || n==1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}