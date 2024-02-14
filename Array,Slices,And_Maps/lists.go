package lists

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99, 12.99, 29.99, 100.10)
	prices = prices[1:]
	fmt.Println(prices)

	// appending the slice or array to the existing price
	discountedPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountedPrices...)
	fmt.Println("The updated prices are : ", prices)
}

// // func main() {

// 	prices := []float64{10.99, 5.64}

// 	//fmt.Println("The array is : ", prices)

// 	//We can also dynamically add the elements into the array by using the append function
// 	//updatedPrice := append(prices, 8.99)
// 	// the above line of code creates an new array with the already existing value along with the new appended value

// 	//fmt.Println("The new updated array: ", updatedPrice, prices)

// 	// So to get the changes made in the array we should assign the prices only again while appending the value , So that go will ditch the original created array and use the newly created array
// 	prices = append(prices, 13.99)
// 	fmt.Println("The updated array is : ", prices)

// }

// // func main() {

// // 	var productNames [5]string = [5]string{"Mango", "Apple"}
// // 	//in the below line of code we are alterring the specific element
// // 	productNames[1] = "grapes"
// // 	productNames[2] = "dragon"
// // 	fmt.Println(productNames)
// // 	prices := [4]float64{10.98, 9.74, 7.65, 6.89}

// // 	fmt.Println("The List of the prices: ", prices)
// // 	// Reading the specific value in an array
// // 	fmt.Println(prices[3]) // here we are accessing the 3rd name

// // 	fmt.Println(productNames)

// // 	//---------creating the slices-------//
// // 	featuredPrices := prices[1:4]
// // 	//here in above slices the first element is included and the second element is exculded

// // 	//-----Vomiting the first element and printing the slices------//
// // 	featuredPrices1 := prices[1:]

// // 	// We can also operate the silced slices
// // 	silcedslice := featuredPrices1[:1]

// // 	// here using the slice to alter the original array
// // 	silcedslice[0] = 199.99

// // 	//Printing the slices of the array
// // 	fmt.Println("The length and capacity of the silced array is :  ", len(silcedslice), cap(silcedslice))
// // 	fmt.Println("The original array after altering is : ", featuredPrices)
// // 	fmt.Println("The sliced slice : ", silcedslice)
// // 	fmt.Println("The featuredPrices1 are : ", featuredPrices1)
// // 	fmt.Println("The featuredPrices are : ", featuredPrices)

// // 	///We will check the capacity after altering the silcedslice

// // 	silcedslice = silcedslice[:3]
// // 	fmt.Println(silcedslice)
// // 	fmt.Println(len(silcedslice), cap(silcedslice))

// // So as u can see the o/p now gives u 3, 3 because now we have taken the base slice as silcedslice and not the feautured slice
// //As u can observe that we have rezied the sliceslice and have given more than the original slicedslice
// //but go internally memoixes te sizemain()

// // }

// //Note: len()-- it gives number of values present in an slice  i.e 1
// //cap()--- it gives how many elements the slice can hold without the reallocation of the memory
// //i.e : now slicedslice is based of featuredslice and featuredslice is based on the base silce "prices" which has 3 more places theoritically hence it gives the o/p as 3
// //---- it cannot give 4 as the o/p because we cannot acces the left of the slice once it is decalred we acn access to the end but not teh starting

// //Note:Slicing
// //-> slices are like pointers for array using slices we can alter the original array
// // -> slicies are like the windows of the array
// //-> highest array an be accessed as mentioned when declaring , if access more than the mentioned array it throws an arrayindexoutofBound error
// //-> Also negative signing doesn't work in slicing of golang
