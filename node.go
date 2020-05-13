package main

import (
	"log"

	bt "github.com/google/btree"
)

type Node interface {
	Less(than bt.Item) bool
	GetKey() uint32
}

type NodeImpl struct {
	Key   uint32
	Value string
}

func (n NodeImpl) GetKey() uint32 {
	return n.Key
}

func (n NodeImpl) Less(than bt.Item) bool {
	b, ok := than.(Node)
	if ok == false {
		log.Fatal("tree item is not a Node")
	}
	return n.Key < b.GetKey()
}
