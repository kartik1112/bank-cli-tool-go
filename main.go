package main

import (
	"fmt"

	bank "com.example/kartik/Bank"
)


func main() {
	AccountBalance := 0.0
	fmt.Println("***************************************************************************")
	fmt.Println("                              Bank Management")
	fmt.Println("***************************************************************************")
	exit := false
	for exit != true {
		var choice int
		fmt.Println("1. Add Money")
		fmt.Println("2. Withdraw Money")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice :")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			bank.AddMoney(&AccountBalance)
			fmt.Printf(`Balance Remaining : %v `,AccountBalance)
			fmt.Println("")
		case 2:
			bank.WithdrawMoney(&AccountBalance)
			fmt.Printf(`Balance Remaining : %v`,AccountBalance)
			fmt.Println("")
		case 3:
			bank.ShowBalance(&AccountBalance)
		case 4:
			exit = true
		default:
			fmt.Println("Enter Valid Input !!!")
		}

	}
}
