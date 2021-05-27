package banking

import (
	"errors"
)

type Account struct {
	owner      string
	balance    int
	isCustomer bool
}

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	account.isCustomer = true
	return &account
}

func SetAccountType(owner *Account, isCustomer bool) {
	owner.isCustomer = isCustomer
}

func GetAccountType(owner *Account) bool {
	return owner.isCustomer
}

func GetBalance(owner *Account) int {
	return owner.balance
}

///////////////////////////////////////////////////////
// Method

// 값을 변경하고 싶을 때
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if a.balance <= 0 {
		return errors.New("Balance is 0...")
	} else if a.balance < amount {
		return errors.New("Over...")
	}

	a.balance -= amount
	return nil
}

// 값을 변경하지 않을 때
func (a Account) Balance() int {
	return a.balance
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

// func (a Account) String() string {
// 	return
// }
