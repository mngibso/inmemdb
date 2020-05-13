package main

import "fmt"

type Transaction interface {
	Begin()
	Rollback()
	Commit()
	HasTransaction() bool
}

type Trxn struct {
	Transactions []map[string]string
}

func NewTrxn() *Trxn {
	ar := make([]map[string]string, 2)
	trxn := Trxn{
		Transactions: ar,
	}
	return &trxn
}

func (t Trxn) Begin() {
	tr := make(map[string]string)
	t.Transactions = append(t.Transactions, tr)
}

func (t Trxn) Commit() {
}

func (t Trxn) HasTransaction() bool {
	return len(t.Transactions) > 0
}

func (t Trxn) Rollback() {
	// removes the last transacton
	if len(t.Transactions) == 0 {
		fmt.Println("TRANSACTION NOT FOUND")
	}
	t.Transactions = t.Transactions[:len(t.Transactions)-1]
}
