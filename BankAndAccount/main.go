package main

import (
	"fmt"

	"github.com/JHKim4532/BankAndAccount/accounts"
)

func main() {
	account := accounts.NewAccounts("jihong")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
