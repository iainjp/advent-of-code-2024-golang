package structure

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestNewList(t *testing.T) {
	want := List[int]{
		data: []int{},
	}

	got := NewList[int]()

	utils.CheckEqual(got, want, t)
}

func TestList(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {
		got := NewList[int]()
		got.Add(1)

		want := List[int]{
			data: []int{1},
		}

		utils.CheckEqual(got, want, t)
	})

	t.Run("AsSlice()", func(t *testing.T) {
		list := NewList[int]()
		list.data = []int{1, 2, 3}
		got := list.AsSlice()

		want := []int{1, 2, 3}

		utils.CheckEqual(got, want, t)
	})

}
