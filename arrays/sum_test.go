package arrays

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 items", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("git %d want %d", got, want)
		}
	})
}
