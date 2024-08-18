package sum

import (
  "testing"
  "slices"
)

func TestSum(t *testing.T) {
  t.Run("collection of any size", func(t *testing.T) {
    numbers := []int {1, 2, 3}

    got := Sum(numbers)
    want := 6

    if got != want {
    t.Errorf("got %d want %d", got, want)
     }
  })
}

func TestSumAll(t *testing.T) {
  t.Run("sum of elements of slices", func(t *testing.T) {
    slice1 := []int {1, 2, 3}
    slice2 := []int {4, 5}
    got := SumAll(slice1, slice2)
    want := 15

    if got != want {
    t.Errorf("got %d want %d", got, want)
     }
  })
}

func TestFlatten(t *testing.T) {
  t.Run("flatten the slices", func(t *testing.T) {
    slice1 := []int {1, 2, 3}
    slice2 := []int {4, 5}
    got := Flatten(slice1, slice2)
    want := []int {1, 2, 3, 4, 5}

    if !slices.Equal(got, want) {
    t.Errorf("got %d want %d", got, want)
     }
  })
}

func TestSumAllTails(t *testing.T) {
  t.Run("sum of all tails element", func(t *testing.T) {
    slice1 := []int {1, 2, 3}
    slice2 := []int {4, 5}

    got := SumAllTails(slice1, slice2)
    want := []int {5, 5}

    if !slices.Equal(got, want) {
    t.Errorf("got %d want %d", got, want)
     }
  })
  t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
    if !slices.Equal(got, want) {
    t.Errorf("got %d want %d", got, want)
    }
	})
}


