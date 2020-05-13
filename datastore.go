package main

import (
	bt "github.com/google/btree"
	"hash/fnv"
	"log"
)

const DEGREE = 5

type Datastorer interface {
	Get(string) string
	Set(string, string)
	Delete(string)
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

func (b BTree) Get(key string) string {
	node := NodeImpl{
		Key: hashKey(key),
	}
	n := b.btree.Get(node)

	// not found
	if n == nil {
		return "NULL"
	}
	out, ok := n.(Node)
	if ok == false {
		log.Fatal("tree item is not a Node")
	}
	return out.GetValue()
}

func (b BTree) Set(key, value string) {
	node := NodeImpl{
		Key:   hashKey(key),
		Value: value,
	}
	b.btree.ReplaceOrInsert(node)
}

func (b BTree) Delete(key string) {
	node := NodeImpl{
		Key: hashKey(key),
	}
	b.btree.Delete(node)
}
