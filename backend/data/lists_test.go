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

/*
func TestCompleteList(t *testing.T) {
	//case true
	lt := NewListTask()
	lt.AddList(testlist[0])
	lt.AddList(testlist[1])
	err := lt.CompleteList(0)
	if lt.GetList()[1].ID != testlist[0].ID {
		t.Error("CompleteList is not working case true")
		t.Log(lt.GetList()[0].Success)
		t.Log(lt.GetList()[1].Success)
	}

	//case false
	ltf := NewListTask()
	ltf.AddList(testlist[0])
	ltf.AddList(testlist[1])
	err = lt.CompleteList(1)
	if ltf.GetList()[0].ID != testlist[1].ID {
		t.Error("CompleteList is not working case false")
		t.Log(lt.GetList()[0].Success)
		t.Log(lt.GetList()[1].Success)
	}
	if err != nil {
		t.Error("Unexpected CompleteList error")
	}

	//Incorrect ID
	err = lt.CompleteList(-1)
	if err == nil {
		t.Error("Not handled out of range error")
	}
}
*/
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
