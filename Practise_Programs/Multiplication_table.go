package main

import "fmt"

func main() {

	var number int

	fmt.Println("Enter the number : ")
	fmt.Scan(&number)

	for i := 1; i < 11; i++ {
		result := number * i
		fmt.Printf("%d x %d = %d\n", number, i, result)
	}

}
