package main

import (
	"fmt"

	"github.com/JHKim4532/nomadGo/BankAndAccount/accounts"
)

func main() {
	account := accounts.NewAccounts("jihong")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account)
}
