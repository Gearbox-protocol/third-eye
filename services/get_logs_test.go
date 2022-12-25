package services

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInsertInSlice(t *testing.T) {
	a := make([]int, 0, 5)
	a = append(a, []int{2, 3, 6}...)

	a = insertInSlice(a, 0)
	if len(a) != 4 || cap(a) != 5 || !cmp.Equal(a, []int{0, 2, 3, 6}) {
		t.Fatal(a, len(a), cap(a))
	}
	a = insertInSlice(a, 7)
	if len(a) != 5 || cap(a) != 5 || !cmp.Equal(a, []int{0, 2, 3, 6, 7}) {
		t.Fatal(a, len(a), cap(a))
	}

	a = insertInSlice(a, 5)
	if len(a) != 6 || !cmp.Equal(a, []int{0, 2, 3, 5, 6, 7}) {
		t.Fatal(a, len(a), cap(a))
	}
	a = insertInSlice(a, 8)
	if len(a) != 7 || !cmp.Equal(a, []int{0, 2, 3, 5, 6, 7, 8}) {
		t.Fatal(a, len(a), cap(a))
	}
}

func TestDeleteInSlice(t *testing.T) {
	a := make([]int, 0, 5)
	a = append(a, []int{0, 2, 3, 6, 7}...)

	a = deleteInSlice(a, 5, map[int]map[string][]Log{})

	if len(a) != 2 || cap(a) != 5 || !cmp.Equal(a, []int{6, 7}) {
		t.Fatal(a, len(a), cap(a))
	}
}
