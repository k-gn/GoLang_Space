package accounts

import (
	"errors"
	"fmt"
)

// 주석작성 시 항상 제목을 앞에 적어줘야한다.

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// NewAccount create Account
func NewAccount(owner string) *Account { // 객체 주소를 반환
	account := Account{owner: owner, balance: 0}
	return &account
}

// *Account == 포인터를 사용하여 복사본을 만들지 않고 실제 객체를 사용해라! == Pointer receiver
// * 를 사용하지 않으면 복사본을 만든다.

// Deposit x amount on your account (method)
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {

	/*
		- exception handling
		Go 에는 exception, try~catch 이런 것들이 없다.
		직접 error 를 Return 하고 체크해야한다.
	*/

	if a.balance < amount {
		// return error.Error()
		return errNoMoney
	}

	a.balance -= amount
	return nil // == null or none
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of your account
func (a Account) Owner() string {
	return a.owner
}

// String for print
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}
