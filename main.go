package main

/*
Commands:
	- SET key value
	- GET key
	- DELETE key
	- COUNT key
	- END
	- BEGIN
	- ROLLBACK
	- COMMIT

*/
func main() {
	b := NewBTree()
	c := NewBTree()
	db := NewDatabase(b, c)
	db.Run()

}
