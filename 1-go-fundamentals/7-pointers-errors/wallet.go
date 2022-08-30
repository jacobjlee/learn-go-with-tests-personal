package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

/*
- & 연산자: 변수의 메모리 주소를 가져옴
- * 포인터: wallet의 포인터를 취해 원본 값을 변경 가능
- 포인터 구조체는 선언하면 자동으로 역참조 됨 (return 부분에 (*w).balance와 같은 방식으로 사용하지 않아도 됨
- nil eqauls to null from other programing languages
*/
