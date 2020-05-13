package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Database struct {
	Datastore   Datastorer
	Counter     Datastorer
	Transaction *Trxn
}

func NewDatabase(ds, cntr Datastorer) *Database {
	trxn := NewTrxn(cntr, ds)
	return &Database{
		ds,
		cntr,
		trxn,
	}
}

func (d Database) Run() {
	fmt.Print("COMMANDS:\n\n" +
		"SET key value\n" +
		"GET key\n" +
		"DELETE key\n" +
		"COUNT value\n" +
		"END\n" +
		"BEGIN\n" +
		"ROLLBACK\n" +
		"COMMIT\n")
	fmt.Println("---------------------------\n")
	scanner := bufio.NewScanner(os.Stdin)
	var t string
	for {
		scanner.Scan()
		t = scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from input: ", err)
		}

		if t == "END" {
			break
		}
		if t == "" {
			continue
		}
		d.processCommand(t)
	}
}

/*
func (d Database) getCounter() Datastorer {
	if(d.Transaction.HasTransaction()) {
		return d.Counter
		//return d.Transaction
	}
	return d.Counter
}

*/

func (d Database) getStorage() Datastorer {
	if d.Transaction.HasTransaction() {
		return d.Transaction
	}
	return d.Datastore
}

// processCommand parses line into a command and parameters and
// executes the command
func (d Database) decrementCount(value string) {
	if d.Transaction.HasTransaction() {
		// Ignore count within transaction for now
		return
	}
	v, ok := d.Counter.Delete(value)
	if ok == true {
		count, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("Invalid non-integer value in counter")
		}
		if count == 1 {
			return
		}
		d.Counter.Set(value, strconv.Itoa(count-1))
	}
}
func (d Database) incrementCount(value string) {
	if d.Transaction.HasTransaction() {
		// Ignore count within transaction for now
		return
	}
	v, ok := d.Counter.Set(value, "1")
	if ok == false {
		return
	}
	count, err := strconv.Atoi(v)
	if err != nil {
		log.Fatal("Invalid non-integer value in counter")
	}
	d.Counter.Set(value, strconv.Itoa(count+1))
}
func (d Database) processCommand(line string) {
	datastore := d.getStorage()
	fields := strings.Fields(line)
	command := fields[0]
	switch command {
	case "GET":
		if len(fields) != 2 {
			fmt.Println("Invalid argument count for 'GET'")
			return
		}
		v, _ := datastore.Get(fields[1])
		fmt.Println(v)

	case "SET":
		if len(fields) != 3 {
			fmt.Println("Invalid argument count for 'SET'")
			return
		}
		v, ok := datastore.Set(fields[1], fields[2])

		// setting to the same value, counter doesn't change
		if ok == true && v == fields[2] {
			return
		}

		if ok == true {
			// value is being replaced
			// decrement counter for old value
			d.decrementCount(v)
		}
		// increment counter for inserted value
		d.incrementCount(fields[2])

	case "DELETE":
		if len(fields) != 2 {
			fmt.Println("Invalid argument count for 'DELETE'")
		}
		v, ok := datastore.Delete(fields[1])
		// nothing was removed
		if ok == false {
			return
		}
		// decrement counter for old value
		d.decrementCount(v)

	case "COUNT":
		if len(fields) != 2 {
			fmt.Println("Invalid argument count for 'COUNT'")
			return
		}
		v, ok := d.Counter.Get(fields[1])

		if ok == false {
			v = "0"
		}
		fmt.Println(v)
	case "BEGIN":
		d.Transaction.Begin()
	case "ROLLBACK":
		d.Transaction.Rollback()
	case "COMMIT":
		transactions := d.Transaction.GetTrxn()
		d.Transaction.Commit()
		for i := 0; i < len(transactions); i++ {
			trxn := transactions[i]
			for y := 0; y < len(trxn); y++ {
				values := strings.Split(trxn[y], " ")
				if len(values) == 2 {
					d.processCommand("SET " + " " + values[0] + " " + values[1])
				} else {
					d.processCommand("DELETE " + " " + values[0])
				}
			}
		}
	default:
		fmt.Println("Invalid Command")
	}
}
