package main

import "fmt"

// Sequence can only go upto 93 due to overflow

func fib_recursive(prev1 uint64, prev2 uint64, counter int, n int) uint64 {
	if n == 0 {return uint64(0)}
	if n == 1 {return uint64(1)}
	if n+1 == counter {return prev2}

	return fib_recursive(prev2, prev1+prev2, counter+1, n)
}

func fib_iterative(n int) uint64 {
	if n == 0 {return uint64(0)}
	if n == 1 {return uint64(1)}

	temp := uint64(0)
	prev1 := uint64(0)
	prev2 := uint64(1)

	for i := 2; i<n+1; i++ {
		temp = prev1 + prev2
		prev1 = prev2
		prev2 = temp
	}
	return prev2
}

func getInput() int {
	var input int

	for {
		fmt.Println("Enter a number between 0 and 93, anything larger will cause an error due to overflow")

		_, err := fmt.Scanln(&input)
		if input < 0 || input > 93 || err != nil {
			fmt.Println("Invalid input")
			continue
		}
		return input
	}
}

func main() {
	var n int

	n = getInput()

	fmt.Println("\nRecursive solution")
	fmt.Println(fib_recursive(0,1,2,n))

	fmt.Println("\nIterative solution")
	fmt.Println(fib_iterative(n))
}
