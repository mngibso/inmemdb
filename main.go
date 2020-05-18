package main

// set up the in memory database and run
func main() {
	b := NewBTree()
	c := NewBTree()
	db := NewDatabase(b, c)
	db.Run()
}
