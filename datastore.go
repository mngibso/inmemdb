package main

import (
	bt "github.com/google/btree"
	"hash/fnv"
	"log"
)

const DEGREE = 5

type Datastorer interface {
	Get(string) (string, bool)
	Set(string, string) (string, bool)
	Delete(string) (string, bool)
}

type BTree struct {
	btree *bt.BTree
}

func NewBTree() *BTree {
	bt := bt.New(DEGREE)
	var d = BTree{bt}
	return &d
}

func hashKey(s string) uint32 {
	h := fnv.New32()
	h.Write([]byte(s))
	return h.Sum32()
}

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
