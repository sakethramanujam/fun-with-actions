package main

import "fmt"

func fib(n int32) int {
	if n <= 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)

}

func main() {
	result := fib(5)
	fmt.Print(result)
}
