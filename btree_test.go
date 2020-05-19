package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBTree(t *testing.T) {
	Convey("Given a BTree", t, func() {
		bTree := NewBTree()
		key := "foo"
		keyDoesntExist := "foobar"
		value := "bar"
		value2 := "bar2"
		Convey("Get'ing an item that doesn't exist should return ('Null', false)", func() {
			key := "foo"
			v, ok := bTree.Get(key)
			So(ok, ShouldEqual, false)
			So(v, ShouldEqual, "NULL")
		})
		Convey("Set'ing an item that does not already exists should return ('', false)", func() {
			v, ok := bTree.Set(key, value)
			So(ok, ShouldEqual, false)
			So(v, ShouldEqual, "")
		})
		Convey("Set'ing an item that does already exists should return ('<prev_value>', true)", func() {
			_, _ = bTree.Set(key, value)
			v, ok := bTree.Set(key, value2)
			So(ok, ShouldEqual, true)
			So(v, ShouldEqual, value)
		})
		Convey("Get'ing an item that does exist should return (<value>, true)", func() {
			_, _ = bTree.Set(key, value)
			v, ok := bTree.Get(key)
			So(v, ShouldEqual, value)
			So(ok, ShouldEqual, true)
		})
		Convey("Delete'ing an item that does'nt exist should return ('', false)", func() {
			_, _ = bTree.Set(key, value)
			v, ok := bTree.Delete(keyDoesntExist)
			So(v, ShouldEqual, "")
			So(ok, ShouldEqual, false)
		})
		Convey("Delete'ing an item that does exist should return (<dltd_value>, true)", func() {
			_, _ = bTree.Set(key, value)
			v, ok := bTree.Delete(key)
			So(v, ShouldEqual, value)
			So(ok, ShouldEqual, true)
		})
	})
	Convey("Given the hash function", t, func() {
		key := "foo"
		dif := "bar"
		Convey("Hashing the same value should return the same hash", func() {
			h1 := hashKey(key)
			h2 := hashKey(key)
			So(h1, ShouldEqual, h2)
		})
		Convey("Hashing the different values should return different hash values", func() {
			h1 := hashKey(key)
			h2 := hashKey(dif)
			So(h1, ShouldNotEqual, h2)
		})
	})
}
