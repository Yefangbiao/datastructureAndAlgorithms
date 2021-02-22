package main

import "fmt"

type wallet struct {
	balance int
}

// 存钱
func (w *wallet) addBalance(amount int) error {
	if amount < 0 {
		return fmt.Errorf("error: addBalance")
	}
	w.balance += amount
	return nil
}

// 取钱
func (w *wallet) reduceBalance(amount int) error {
	if amount < 0 || amount > w.balance {
		return fmt.Errorf("error: reduceBalance")
	}
	w.balance -= amount
	return nil
}
