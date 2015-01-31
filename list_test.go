package main

import (
	"testing"

	"github.com/PreetamJinka/listgen/intlist"
	"github.com/PreetamJinka/listgen/stringlist"
)

func TestIntList(t *testing.T) {
	l := intlist.NewList()
	l.Insert(4)
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(-1)
	l.Insert(-2)
	l.Insert(9)
	l.Insert(5)

	err := l.Insert(4)
	if err == nil {
		t.Fatal("expected error:", intlist.ErrIntListValueExists)
	}

	err = l.Remove(4)
	if err != nil {
		t.Fatal(err)
	}

	err = l.Insert(4)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(l)

	expectedOrder := []int{
		-2, -1, 1, 2, 3, 4, 5, 9,
	}

	i := l.NewIterator()

	for i.Next() {
		v, err := i.Value()
		if len(expectedOrder) == 0 {
			if err == nil {
				t.Errorf("extra value: %v", v)
			}

			t.Error(err)
			continue
		}

		if err != nil {
			t.Error(err)
			continue
		}

		if v != expectedOrder[0] {
			t.Errorf("expected value %v, got %v", expectedOrder[0], v)
			continue
		}

		expectedOrder = expectedOrder[1:]
	}

	if i.Next() {
		t.Error("expected Next to return false")
	}

	if v, err := i.Value(); err == nil {
		t.Errorf("expected err to be non-nil. Got value %v", v)
	}

	if size := l.Size(); size != 8 {
		t.Errorf("expected size to be %d, got %d", 8, size)
	}
}

func TestStringList(t *testing.T) {
	l := stringlist.NewList()
	l.Insert("f")
	l.Insert("o")
	l.Insert("b")
	l.Insert("a")
	l.Insert("r")
	l.Insert("z")

	err := l.Insert("o")
	if err == nil {
		t.Fatal("expected error:", intlist.ErrIntListValueExists)
	}

	err = l.Remove("o")
	if err != nil {
		t.Fatal(err)
	}

	err = l.Insert("o")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(l)

	expectedOrder := []string{
		"a", "b", "f", "o", "r", "z",
	}

	i := l.NewIterator()

	for i.Next() {
		v, err := i.Value()
		if len(expectedOrder) == 0 {
			if err == nil {
				t.Errorf("extra value: %v", v)
			}

			t.Error(err)
			continue
		}

		if err != nil {
			t.Error(err)
			continue
		}

		if v != expectedOrder[0] {
			t.Errorf("expected value %v, got %v", expectedOrder[0], v)
			continue
		}

		expectedOrder = expectedOrder[1:]
	}

	if i.Next() {
		t.Error("expected Next to return false")
	}

	if v, err := i.Value(); err == nil {
		t.Errorf("expected err to be non-nil. Got value %v", v)

	}

	if size := l.Size(); size != 6 {
		t.Errorf("expected size to be %d, got %d", 6, size)
	}
}
