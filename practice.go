package main
import "fmt"
func practice() {
	// fmt.Println(quote.Go())
	var x = 60
	var arr = []int{40, 50, x, 90}
	// var ans = [5]int{5, 4, 5,6,7} //initializes array with elements
	var ans = make([]int, 5)
	fmt.Println(len(ans))
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
	// var arr1 = make([]int, 2, 5)
	// arr = append(arr, 1)
	// arr = append(arr, 1)
	// arr = append(arr, 1) // the difference between len() and cap() is len defines the no. of elements in the slice whereas 
	// arr = append(arr, 1)  // cap() defines the size, underlying array can hold upto 
	// arr = append(arr, 1)
	// arr = append(arr, 1)
	// arr = append(arr, 1)
	// arr = append(arr, 1)
	// arr = append(arr, 1)
	// arr = append(arr, 1)
	// arr = append(arr, 1)
	// arr = append(arr, 1)
	// fmt.Println(arr)
	// fmt.Println(len(arr))
	// fmt.Println(cap(arr))

	/* copy slices in go */
	
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// fmt.Println(nums)
	// var nums1 = make([]int, len(nums))
	// copy(nums1, nums)
	// fmt.Println(nums1)
	fmt.Print(call)
}

func fib(n int) int {
	if n==0 || n==1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}