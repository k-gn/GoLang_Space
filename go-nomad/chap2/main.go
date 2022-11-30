package main

import (
	"fmt"

	"github.com/k-gn/learngo/go-nomad/chap2/accounts"
	"github.com/k-gn/learngo/go-nomad/chap2/mydict"
)

func main() {
	account := accounts.NewAccount("gyunam")
	fmt.Println(account)

	account.Deposit(10)
	fmt.Println(account.Balance())

	err := account.Withdraw(20)
	// error check
	if err != nil {
		fmt.Println(err)
		// log.Fatalln(err) // err log + exit
	}

	account.ChangeOwner("dongyu")
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)

	// ------------------------------------------------------

	dictionary := mydict.Dictionary{"first": "first word"}
	value, err := dictionary.Search("first")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	word := "hello"
	definition := "Greeting"

	addErr := dictionary.Add(word, definition)
	if addErr != nil {
		fmt.Println(addErr)
	}

	addValue, _ := dictionary.Search(word)
	fmt.Println(addValue)

	addErr2 := dictionary.Add(word, definition)
	if addErr2 != nil {
		fmt.Println(addErr2)
	}

	upErr := dictionary.Update(word, "update Greeting")
	if upErr != nil {
		fmt.Println(addErr2)
	}

	upValue, _ := dictionary.Search(word)
	fmt.Println(upValue)

	dictionary.Delete(word)
	delValue, delErr := dictionary.Search(word)
	fmt.Println(delValue, delErr)
}
