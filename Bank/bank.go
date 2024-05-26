package bank

import "fmt"

// import "fmt"

func AddMoney(AccountBalance *float64) {
	var deposit float64
	fmt.Print("Enter deposit amount : ")
	fmt.Scan(&deposit)
	*AccountBalance = *AccountBalance + deposit
}

func WithdrawMoney(AccountBalance *float64) {
	var withdraw float64
	fmt.Print("Enter withdraw amount : ")
	fmt.Scan(&withdraw)
	if *AccountBalance-withdraw <= 0 {
		fmt.Println("Insufficient Funds")
	} else {
		*AccountBalance = *AccountBalance - withdraw
	}

}

func ShowBalance(AccountBalance *float64) {
	fmt.Printf(`Balance Remaining : %v`,*AccountBalance)
}
