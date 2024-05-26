package main

import (
	"fmt"

	accounts "github.com/olcw78/yt/learn-go-nomadcoders/accounts"
)

func main() {
	account := accounts.NewAccount("yoon")
	account.Deposit(1000)

	if err := account.Withdraw(1001); err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}
