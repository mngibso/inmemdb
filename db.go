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
	trxn := NewTrxn()
	return &Database{
		ds,
		cntr,
		trxn,
	}
}

func (d Database) Run() {
	fmt.Print("SET key value\n" +
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

// processCommand parses line into a command and parameters and
// executes the command
// Possible commands:
// 	Set key value
// 	Get key
// 	Delete key
func (d Database) decrementCount(value string) {
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
	fields := strings.Fields(line)
	command := fields[0]
	switch command {
	case "GET":
		if len(fields) != 2 {
			fmt.Println("Invalid argument count for 'GET'")
			return
		}
		v, _ := d.Datastore.Get(fields[1])
		fmt.Println(v)

	case "SET":
		if len(fields) != 3 {
			fmt.Println("Invalid argument count for 'SET'")
			return
		}
		v, ok := d.Datastore.Set(fields[1], fields[2])

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
		v, ok := d.Datastore.Delete(fields[1])
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
	case "ROLLBACK":
	case "COMMIT":
		fmt.Println("Not implemented")
	default:
		fmt.Println("Invalid Command")
	}
}
