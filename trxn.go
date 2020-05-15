package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Trxn struct {
	datastore    Datastorer
	counter      Datastorer
	Transactions [][]string
	counts       []map[string]int
}

func NewTrxn(datastore, counter Datastorer) *Trxn {
	ar := [][]string{}
	c := []map[string]int{}
	trxn := Trxn{
		datastore,
		counter,
		ar,
		c,
	}
	return &trxn
}

func (t *Trxn) Begin() {
	tr := []string{}
	c := map[string]int{}
	t.Transactions = append(t.Transactions, tr)
	t.counts = append(t.counts, c)
}

func (t *Trxn) Clear() {
	t.Transactions = [][]string{}
	t.counts = []map[string]int{}
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
	t.counts = t.counts[:len(t.counts)-1]
}

func (t *Trxn) Delete(key string) (string, bool) {
	if t.HasTransaction() == false {
		return "", false
	}
	count := t.counts[len(t.counts)-1]
	v, ok := t.Get(key)
	if ok == true {
		if _, ok := count[v]; ok {
			count[v]--
		} else {
			count[v] = -1
		}
	}
	trxn := t.Transactions[len(t.Transactions)-1]
	trxn = append(trxn, key)
	t.Transactions[len(t.Transactions)-1] = trxn
	return "", false
}

func (t *Trxn) Set(key, value string) (string, bool) {
	if t.HasTransaction() == false {
		return "", false
	}
	count := t.counts[len(t.counts)-1]
	// decrement count of old value
	v, ok := t.Get(key)
	if ok {
		if _, ok := count[v]; ok {
			count[v]--
		} else {
			count[v] = -1
		}
	}
	// increment count of new value
	if _, ok := count[value]; ok {
		count[value]++
	} else {
		count[value] = 1
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

func (t *Trxn) Count(value string) (string, bool) {
	c := 0
	if cntr, ok := t.counter.Get(value); ok {
		c, _ = strconv.Atoi(cntr)
	}
	for i := len(t.counts) - 1; i >= 0; i-- {
		count := t.counts[i]
		if v, ok := count[value]; ok {
			// return strconv.Itoa(c + v), true
			c += v
		}
	}
	return strconv.Itoa(c), false
}
