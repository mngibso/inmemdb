package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

func TestTrxn(t *testing.T) {
	Convey("Given a Transaction object", t, func() {
		key := "foo"
		value := "bar"
		d := MockDatastorer{mock.Mock{}}
		c := MockDatastorer{mock.Mock{}}
		Convey("With no transactions", func() {
			trxn := NewTrxn(&d, &c)
			Convey("Get should be called on the datastore object", func() {
				d.On("Get", key).Once().Return("", false)
				trxn.Get(key)
				d.AssertExpectations(t)
			})
			Convey("Set should not cause an error", func() {
				v, ok := trxn.Set(key, value)
				So(ok, ShouldEqual, false)
				So(v, ShouldEqual, "")
			})
			Convey("Delete should not cause an error", func() {
				v, ok := trxn.Delete(key)
				So(ok, ShouldEqual, false)
				So(v, ShouldEqual, "")
			})
			Convey("Clear should not cause an error", func() {
				trxn.Clear()
			})
			Convey("Rollback should not cause an error", func() {
				trxn.Rollback()
			})
			Convey("HasTransaction should return false", func() {
				b := trxn.HasTransaction()
				So(b, ShouldEqual, false)
			})
			Convey("GetTransaction should return no transactions", func() {
				b := trxn.GetTrxn()
				So(len(b), ShouldEqual, 0)
			})
			Convey("Count should return call count.Get", func() {
				c.On("Get", value).Once().Return("0", false)
				v, ok := trxn.Count(value)
				So(v, ShouldEqual, "0")
				So(ok, ShouldEqual, false)
			})
			Convey("Begin should add a transaction", func() {
				trxn.Begin()
				b := trxn.GetTrxn()
				So(len(b), ShouldEqual, 1)
			})
		})
		Convey("With transactions", func() {
			trxn := NewTrxn(&d, &c)
			key := "foo"
			value := "bar"
			trxn.Begin()
			// _,_ = trxn.Set("key","bar")
			Convey("Get'ing an item that doesn't exist should call datastore.Get", func() {
				d.On("Get", key).Once().Return("", false)
				_, _ = trxn.Get(key)
				d.AssertExpectations(t)
			})
			Convey("Set'ing an item that exist in the transaction adds it to the transaction", func() {
				d.On("Get", key).Once().Return("", false)
				_, _ = trxn.Set(key, value)
				b := trxn.GetTrxn()
				So(len(b), ShouldEqual, 1)
				So(len(b[0]), ShouldEqual, 1)
				v := []string{key + " " + value}
				So(b[0], ShouldResemble, v)
			})
			/*
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

			*/
		})
	})
}
