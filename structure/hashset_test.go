package structure

import (
	"testing"

	"iain.fyi/aoc2024/utils"
)

func TestNewHashSet(t *testing.T) {
	t.Run("NewHashSet()", func(t *testing.T) {
		hashset := NewHashSet[int]()
		got := hashset.entries != nil
		utils.CheckEqual(got, true, t)
	})
}

func TestHashset(t *testing.T) {
	t.Run("Add()", func(t *testing.T) {
		hashset := NewHashSet[int]()
		hashset.Add(1)
		want := true
		got := hashset.entries[1]

		utils.CheckEqual(got, want, t)
	})

	t.Run("AddFromSlice()", func(t *testing.T) {
		hashset := NewHashSet[int]()
		hashset.AddFromSlice([]int{1, 2, 3})

		utils.CheckEqual(hashset.entries[1], true, t)
		utils.CheckEqual(hashset.entries[2], true, t)
		utils.CheckEqual(hashset.entries[3], true, t)
	})

	t.Run("AddAll()", func(t *testing.T) {
		hashset := NewHashSet[int]()
		hashset.AddAll(1, 2, 3)

		utils.CheckEqual(hashset.entries[1], true, t)
		utils.CheckEqual(hashset.entries[2], true, t)
		utils.CheckEqual(hashset.entries[3], true, t)
	})

	t.Run("Size()", func(t *testing.T) {
		hashset := NewHashSet[int]()
		hashset.entries = map[int]bool{
			1: true,
			2: true,
			3: true,
		}
		want := 3
		got := hashset.Size()

		utils.CheckEqual(got, want, t)
	})

	t.Run("Remove()", func(t *testing.T) {
		hashset := NewHashSet[int]()
		hashset.entries = map[int]bool{
			1: true,
			2: true,
			3: true,
		}

		hashset.Remove(1)

		utils.CheckEqual(len(hashset.entries), 2, t)
		utils.CheckEqual(hashset.entries[1], false, t)
		utils.CheckEqual(hashset.entries[2], true, t)
		utils.CheckEqual(hashset.entries[3], true, t)
	})

	t.Run("Contains()", func(t *testing.T) {
		hashset := NewHashSet[int]()
		hashset.entries = map[int]bool{
			1: true,
			2: true,
		}

		utils.CheckEqual(hashset.Contains(1), true, t)
		utils.CheckEqual(hashset.Contains(2), true, t)
		utils.CheckEqual(hashset.Contains(3), false, t)
	})

	t.Run("Clone()", func(t *testing.T) {
		hashset := NewHashSet[int]()
		hashset.entries = map[int]bool{
			1: true,
			2: true,
			3: true,
		}

		clone := hashset.Clone()

		utils.CheckEqual(len(clone.entries), 3, t)
		utils.CheckEqual(clone.entries[1], true, t)
		utils.CheckEqual(clone.entries[2], true, t)
		utils.CheckEqual(clone.entries[3], true, t)
	})
}
