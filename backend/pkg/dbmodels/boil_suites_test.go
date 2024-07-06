// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Messages", testMessages)
}

func TestSoftDelete(t *testing.T) {}

func TestQuerySoftDeleteAll(t *testing.T) {}

func TestSliceSoftDeleteAll(t *testing.T) {}

func TestDelete(t *testing.T) {
	t.Run("Messages", testMessagesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Messages", testMessagesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Messages", testMessagesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Messages", testMessagesExists)
}

func TestFind(t *testing.T) {
	t.Run("Messages", testMessagesFind)
}

func TestBind(t *testing.T) {
	t.Run("Messages", testMessagesBind)
}

func TestOne(t *testing.T) {
	t.Run("Messages", testMessagesOne)
}

func TestAll(t *testing.T) {
	t.Run("Messages", testMessagesAll)
}

func TestCount(t *testing.T) {
	t.Run("Messages", testMessagesCount)
}

func TestInsert(t *testing.T) {
	t.Run("Messages", testMessagesInsert)
	t.Run("Messages", testMessagesInsertWhitelist)
}

func TestReload(t *testing.T) {
	t.Run("Messages", testMessagesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Messages", testMessagesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Messages", testMessagesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Messages", testMessagesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Messages", testMessagesSliceUpdateAll)
}