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


