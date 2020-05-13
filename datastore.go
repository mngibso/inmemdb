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

func hashKey(s string) uint64 {
	h := fnv.New64()
	h.Write([]byte(s))
	return h.Sum64()
}

func (b BTree) Get(key string) string {
	return ""
}

func (b BTree) Set(key, value string) {
}

func (b BTree) Delete(key string) {
}
