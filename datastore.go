package main

import (
	bt "github.com/google/btree"
	"hash/fnv"
	"log"
)

const DEGREE = 5

// A Datastorer stores, retrieves and deletes data.
type Datastorer interface {
	Get(string) (string, bool)
	Set(string, string) (string, bool)
	Delete(string) (string, bool)
}

// A BTree implements a Datastorer using a BTree.
type BTree struct {
	btree *bt.BTree
}

// NewBTree returns a pointer to an initialized BTree
func NewBTree() *BTree {
	bt := bt.New(DEGREE)
	var d = BTree{bt}
	return &d
}

// hashKey hashes a string into a number. Collisions are rare but possible.
func hashKey(s string) uint32 {
	h := fnv.New32()
	h.Write([]byte(s))
	return h.Sum32()
}

// Get returns the value of `key` in the BTree, or `false` if the key isn't present in the BTree.
func (b BTree) Get(key string) (string, bool) {
	node := NodeImpl{
		Key: hashKey(key),
	}
	n := b.btree.Get(node)

	// not found
	if n == nil {
		return "NULL", false
	}
	out, ok := n.(Node)
	if ok == false {
		log.Fatal("tree item is not a Node")
	}
	return out.GetValue(), true
}

// Set sets `key` to `value` in the BTree.  It returns the value replaced if `key` already existed, or `false`
// if the key did not exist before the Set call.
func (b BTree) Set(key, value string) (string, bool) {
	node := NodeImpl{
		Key:   hashKey(key),
		Value: value,
	}

	// return the value replaced, if item exists
	i := b.btree.ReplaceOrInsert(node)
	if i == nil {
		return "", false
	}
	out, ok := i.(Node)
	if ok == false {
		log.Fatal("tree item is not a Node")
	}
	return out.GetValue(), true
}

// Delete removes `key` from the BTree. Returns the value deleted and `true`, or `false` if `key` doesn't exist.
func (b BTree) Delete(key string) (string, bool) {
	node := NodeImpl{
		Key: hashKey(key),
	}

	// return the value deleted, if item exists
	i := b.btree.Delete(node)
	if i == nil {
		return "", false
	}
	out, ok := i.(Node)
	if ok == false {
		log.Fatal("tree item is not a Node")
	}
	return out.GetValue(), true
}
