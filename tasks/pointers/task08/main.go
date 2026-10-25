package main

import "fmt"

func main() {
	var account Account = Account{Owner: "Иван", Balance: 100}
	fmt.Println(account.deposit(-100))
	fmt.Println(account.Balance)
	fmt.Println(account.withdraw(101))
	fmt.Println(account.Balance)
}

type Account struct {
	Owner   string
	Balance int
}

func (a *Account) deposit(amount int) bool {
	if amount > 0 {
		(*a).Balance += amount
		return true
	}
	return false
}

func (a *Account) withdraw(amount int) bool {
	if amount > 0 && (*a).Balance >= amount {
		(*a).Balance -= amount
		return true
	}
	return false
}
