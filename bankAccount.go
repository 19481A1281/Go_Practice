package main

import "fmt"

type BankAccount struct {
	AccountNumber string
	Balance       float64
}

func (account *BankAccount) Deposit(amount float64) {
	account.Balance += amount
}

func (account *BankAccount) Withdraw(amount float64) error {
	if amount > account.Balance {
		return fmt.Errorf("Insufficient balance")
	}

	account.Balance -= amount
	return nil
}

func (account BankAccount) GetBalance() float64 {
	return account.Balance
}

func main() {
	account := BankAccount{AccountNumber: "1234", Balance: 2000.67}
	fmt.Println("Initial Balance:", account.GetBalance())

	account.Deposit(400)
	fmt.Println("Balance after deposit:", account.GetBalance())

	var amount float64
	fmt.Print("Enter amount to withdraw:")
	fmt.Scan(&amount)

	err := account.Withdraw(amount)
	if err != nil {
		fmt.Println("Error while widhdrawing:", err)
	} else {
		fmt.Println("Balnce after withdrawl:", account.GetBalance())
	}

}
