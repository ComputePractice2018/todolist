package data

import (
	"testing"
)

var testlist = []List{
	{
		ID:      1,
		Name:    "abc",
		Success: true,
	},
	{
		ID:      2,
		Name:    "abc",
		Success: false,
	},
	{
		ID:      3,
		Name:    "ghi",
		Success: true,
	},
}

func TestAddList(t *testing.T) {
	lt := NewListTask()
	lt.AddList(testlist[0])
	if lt.GetList()[0] != testlist[0] {
		t.Error("AddList is not working")
	}
}

func TestEditList(t *testing.T) {
	lt := NewListTask()
	lt.AddList(testlist[1])
	err := lt.EditList(testlist[1], 1)
	if lt.GetList()[0] != testlist[1] {
		t.Error("EditList is not working")
	}
	if err != nil {
		t.Error("Unexpected EditList error")
	}

	err = lt.EditList(testlist[2], -1)
	if err == nil {
		t.Error("Not handled out of range error")
	}
}

func TestDeleteList(t *testing.T) {
	lt := NewListTask()
	lt.AddList(testlist[0])
	lt.AddList(testlist[1])

	err := lt.DeleteList(1)
	if err != nil {
		t.Error("Unexpected EditList error")
	}
	if lt.GetList()[0] != testlist[1] {
		t.Error("DeleteList is not working")
	}
	err = lt.DeleteList(-1)
	if err == nil {
		t.Error("Not handled out of range error")
	}
}
