package main

import "fmt"

func main() {
	var account1 Account = Account{Owner: "Иван", Balance: 100}
	var account2 Account = Account{Owner: "Пётр", Balance: 20}
	fmt.Println(transfer(&account1, &account2, 100))
	fmt.Println(account1.Balance)
	fmt.Println(account2.Balance)
}

type Account struct {
	Owner   string
	Balance int
}

func transfer(from *Account, to *Account, amount int) bool {
	if from != nil && to != nil && amount > 0 && (*from).Balance >= amount {
		(*from).Balance -= amount
		(*to).Balance += amount
		return true
	}
	return false
}
