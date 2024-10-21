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
	var ans = [5]int{5, 4, 5,6,7} //initializes array with elements
	fmt.Println(ans) 
	// var vec [5]int   --> declares an int array named vec
	fmt.Println("hello go")
	// for 2d array
	nums := [3][3]int {{1,2,3}, {4,5,6}}
	fmt.Println(nums)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	var call = fib(6)
	
	var name string = "hello"
	fmt.Println(name)

	varType := func(i interface{}) {
		switch t:=i.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("boolean")
		default:
			fmt.Println("others", t)
		}
	}

	varType("jsuhnjd")

	fmt.Print(call)
	// http.HandleFunc("/", httpHandler)
	// http.ListenAndServe(":9000", nil)

}


func fib(n int) int {
	if n==0 || n==1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}



