# inmemdb

## Description

In memory key/value database.  This project was created as part of challenge 
to provide an in memory key/value store with O(log n) read and write performance. 
It supports embedded transactions. Interface is through the console.

## Setup and run
### Prerequisites
- [Go](https://golang.org/dl/)

Install and set up your go environment.

### Get inmemdb

`$ go get github.com/mngibso/inmemdb`

### Start inmemdb

```
$ cd $GOPATH/src/github.com/mngibson/inmemdb
$ go build
$ ./inmemdb
```

## Commands

`key` and `value` must contain no whitespace

- `SET key value`
- `GET key`
- `DELETE key`
- `COUNT value` : Provides a count of the number of keys that are set to `value` in the database.
- `BEGIN` : Start a transaction.
- `ROLLBACK` : Rollback the most recent transaction.
- `COMMIT` : Commit *all* transactions.
- `END` : Quit the application. 

## Example:

```
$ ./inmemdb
COMMANDS:

SET key value
GET key
DELETE key
COUNT value
END
BEGIN
ROLLBACK
COMMIT
---------------------------

SET a foo
SET b bar
SET c foo
COUNT foo
2
BEGIN
DELETE a
COUNT foo
1
GET a 
NULL
BEGIN
DELETE c
COUNT foo
0
ROLLBACK
COUNT foo
1
COMMIT
GET a 
NULL
GET b
bar
GET c
foo
END
```
