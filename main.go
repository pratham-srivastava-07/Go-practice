package main

import "fmt"
// import "rsc.io/quote"

func main() {
	// fmt.Println(quote.Go())
	var x = 60
	var arr = []int{40, 50, x, 90}

	fmt.Println("hello go")

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	var call = fib(6)

	fmt.Print(call)
}


func fib(n int) int {
	if n==0 || n==1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}