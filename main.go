// package main

// import (
// 	"fmt"
// 	"net/http"
// )
// // import "rsc.io/quote"

// func httpHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "hello %s welcome to the server", r.URL.Path[1:])
// 	// fmt.Fprintf(w, "hello %d, welcome to the Go server", r.URL.Path[1:])
// 	var log = r.URL.Path[1:]
// 	if log == "favicon.ico" {
// 		return;
// 	}
// 	fmt.Println(log)

// }

// func main() {
// 	var nums = []int{1,2,3,4}
// 	fmt.Println(nums)

// 	var arr = make([]int, 2, 7)
// 	arr = append(arr, 1)
// 	arr = append(arr, 5)
// 	fmt.Println(arr)
// 	http.HandleFunc("/", httpHandler)
// 	http.ListenAndServe(":9000", nil)

// }

package main

import "fmt"

type shape interface {
	area()
	perimeter()
}

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}
type rect struct {
	width int
	height int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" || mToSend.sender.number == 0{
		return false
	}
	return true
}

func area(r rect) int {
	return r.width * r.height
}


func getUser(resolve,  err string) string {
	return "some data"
}
func main() {
	res := canSendMessage(messageToSend{message: "hello", sender: user{name: "", number: 738839}, recipient: user{name: "gungun", number: 5555}})
	fmt.Println(res)

	rectArea:= area(rect{width: 7, height: 8})
	fmt.Println((rectArea))

	// error handling in go

resolve, err := getUser()

if err != nil {
	fmt.Println(err)
	return
} else {
	fmt.Println(resolve)
}
}




