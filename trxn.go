package main

import (
	"fmt"
	"strings"
)

type Trxn struct {
	counter      Datastorer
	datastore    Datastorer
	Transactions [][]string
	commiting    bool
}

func NewTrxn(counter, datastore Datastorer) *Trxn {
	ar := [][]string{}
	trxn := Trxn{
		counter,
		datastore,
		ar,
		false,
	}
	return &trxn
}

func (t *Trxn) Begin() {
	tr := []string{}
	t.Transactions = append(t.Transactions, tr)
}

func (t *Trxn) Commit() {
	t.Transactions = [][]string{}
}

func (t Trxn) HasTransaction() bool {
	return len(t.Transactions) > 0
}

func (t *Trxn) GetTrxn() [][]string {
	return t.Transactions
}

func (t *Trxn) Rollback() {
	// removes the last transacton
	if len(t.Transactions) == 0 {
		fmt.Println("TRANSACTION NOT FOUND")
		return
	}
	t.Transactions = t.Transactions[:len(t.Transactions)-1]
}

func (t *Trxn) Delete(key string) (string, bool) {
	if t.HasTransaction() == false {
		return "", false
	}
	trxn := t.Transactions[len(t.Transactions)-1]
	trxn = append(trxn, key)
	return "", false
}

func (t *Trxn) Set(key, value string) (string, bool) {
	if t.HasTransaction() == false {
		return "", false
	}
	trxn := t.Transactions[len(t.Transactions)-1]
	trxn = append(trxn, key+" "+value)
	t.Transactions[len(t.Transactions)-1] = trxn
	return "", false
}

func (t *Trxn) Get(key string) (string, bool) {
	// check the transactions for this key
	for i := len(t.Transactions) - 1; i >= 0; i-- {
		trxn := t.Transactions[i]
		for y := len(trxn) - 1; y >= 0; y-- {
			values := strings.Split(trxn[y], " ")
			if values[0] == key {
				if len(values) == 2 {
					// SET
					return values[1], true
				}
				// DELETE
				return "NULL", false
			}
		}
	}
	return t.datastore.Get(key)
}

