package main

import (
	"fmt"
)

const accountBalanceFile = "balance.txt"

func main() {
	var AccountBalance, err = GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		panic("Can't continue Sorry...!!!")
	}

	fmt.Println("Welcome to Go Bank!")
	for {
		PresenterOptions()

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is ", AccountBalance)
		case 2:
			fmt.Println("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid Input! must be greater than 0 ...")
				continue // Here the return keyword is added because after the display of this statement the further code should not be executed
			}

			AccountBalance += depositAmount // AccountBalance
			fmt.Println("Balance update! New amount: ", AccountBalance)
			WriteValueToFile(AccountBalance, accountBalanceFile)
		case 3:
			fmt.Println("Withdrawal Amount: ")
			var Withdrawal float64
			fmt.Scan(&Withdrawal)
			if Withdrawal <= 0 {
				fmt.Println("Invalid Input! must be greater than 0 ...")
				continue // this continue keyword will not continue to excute code the remaining code but it will restart the code again to give
				//nother chance and it will display the options again]
			}
			if Withdrawal > AccountBalance {
				fmt.Println("Invalid amount . You can't withdraw more than you have")
				continue
			}
			AccountBalance -= Withdrawal
			fmt.Println("Balance update! After Withdrawal: ", AccountBalance)
			WriteValueToFile(AccountBalance, accountBalanceFile)

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for Choosing our Bank.... have a nice day..")
			return

		}
	}
}

//When you want to break out of the loop the switch case statement willl be a wrong choice we can use the ifelse statement instead
