package main

import "fmt"

func main() {

	/// This is an example of integers
	x, y := Swap(5, 9)
	fmt.Println("The numbers After Swapping: ", x, y)

	// This is an example of Strings
	x1, y1 := Swap("Meghana.S", "Sahana.S")
	fmt.Println("The two Strings after Swapping: ", x1, y1)

}

func Swap[T any](a, b T) (T, T) {
	return b, a
}
