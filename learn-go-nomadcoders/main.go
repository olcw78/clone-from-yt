package main

import (
	"fmt"

	mydict "github.com/olcw78/yt/learn-go-nomadcoders/dict"
)

func main() {
	// account := accounts.NewAccount("yoon")
	// account.Deposit(1000)

	// if err := account.Withdraw(1001); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(account)

	dict := mydict.Dictionary{"first": "apple"}
	if v, err := dict.Search("first"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	dict.Add("second", "orange")
	if v, err := dict.Search("second"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	word := "hello"
	definition := "Greeting"

	if err := dict.Add(word, definition); err != nil {
		fmt.Println(err)
	}

	if hello, err := dict.Search(word); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(hello)
	}
}
