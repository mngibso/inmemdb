package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Transactioner interface {
	Clear()
	Begin()
	HasTransaction() bool
	GetTrxn() [][]string
	Rollback()
	Get(string) (string, bool)
	Set(string, string) (string, bool)
	Delete(string) (string, bool)
	Count(value string) (string, bool)
}

// A Trxn stores multiple transactions and implements the Datastorer interface. Transactions may be embedded
// within each other.
// Trxn must store the key/values as well as keep a running count of each value in the Datastorer.
type Trxn struct {
	datastore    Datastorer
	counter      Datastorer
	Transactions [][]string
	counts       []map[string]int
}

// New Trxn returns a pointer to a Trxn.
// `datastore` is a reference to the "live" storage, and `counter` is a reference to the "live" storage of
// counts of values.
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

// Begin starts a new transaction
func (t *Trxn) Begin() {
	tr := []string{}
	c := map[string]int{}
	t.Transactions = append(t.Transactions, tr)
	t.counts = append(t.counts, c)
}

// Clear removes all transactions
func (t *Trxn) Clear() {
	t.Transactions = [][]string{}
	t.counts = []map[string]int{}
}

// HasTransaction returns `true` if transactions exist
func (t Trxn) HasTransaction() bool {
	return len(t.Transactions) > 0
}

// GetTrxn() returns all active transaction arrays
func (t *Trxn) GetTrxn() [][]string {
	return t.Transactions
}

// Rollback removes the latest transaction
func (t *Trxn) Rollback() {
	if len(t.Transactions) == 0 {
		fmt.Println("TRANSACTION NOT FOUND")
		return
	}
	t.Transactions = t.Transactions[:len(t.Transactions)-1]
	t.counts = t.counts[:len(t.counts)-1]
}

// Delete removes `key` from the datastore.
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

// Set sets key to value in the datastore
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

// Get returns the value for `key` and `true`.  Returns `false` if the `key does not exist.
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

// Count returns the number of times `value` is stored in teh datastore
func (t *Trxn) Count(value string) (string, bool) {
	c := 0
	// get the count from the datastore
	if cntr, ok := t.counter.Get(value); ok {
		c, _ = strconv.Atoi(cntr)
	}

	// loop through the transaction count array and adjust the count value accordingly.
	for i := len(t.counts) - 1; i >= 0; i-- {
		count := t.counts[i]
		if v, ok := count[value]; ok {
			c += v
		}
	}
	return strconv.Itoa(c), false
}
