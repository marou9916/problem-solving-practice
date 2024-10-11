package training

import "fmt"

type BankAccount struct {
	Owner string
	Sold  float64
}

func (account *BankAccount) Deposit(money float64) {
	account.Sold += money
}

func (account *BankAccount) Withdraw(money float64) {
	if money > account.Sold {
		fmt.Println("Not enough money buddy, i'm sorry")
		return
	}
	account.Sold -= float64(money)
}

func (account BankAccount) DisplaySold() {
	fmt.Printf("Sold : %f", account.Sold)
}
