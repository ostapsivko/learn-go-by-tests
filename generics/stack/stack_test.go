package stack_test

import (
	"testing"
)

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

func TestStack(t *testing.T) {
	t.Run("stack of ints", func(t *testing.T) {
		myStackOfInts := new(Stack[int])

		//check stack is empty
		AssertTrue(t, myStackOfInts.IsEmpty())

		//add a thing, then check it's not empty
		myStackOfInts.Push(1)
		AssertFalse(t, myStackOfInts.IsEmpty())

		//add another thing and pop it back again
		myStackOfInts.Push(2)
		val, ok := myStackOfInts.Pop()
		AssertFalse(t, myStackOfInts.IsEmpty())
		AssertTrue(t, ok)
		AssertEqual(t, val, 2)
		val, ok = myStackOfInts.Pop()
		AssertTrue(t, ok)
		AssertEqual(t, val, 1)
		AssertTrue(t, myStackOfInts.IsEmpty())

		//getting typed numbers back
		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		first, _ := myStackOfInts.Pop()
		second, _ := myStackOfInts.Pop()
		AssertEqual(t, first+second, 3)
	})

	t.Run("stack of strings", func(t *testing.T) {
		myStackOfStrings := new(Stack[string])

		//check stack is empty
		AssertTrue(t, myStackOfStrings.IsEmpty())

		//add a thing, then check it's not empty
		myStackOfStrings.Push("1")
		AssertFalse(t, myStackOfStrings.IsEmpty())

		//add another thing and pop it back again
		myStackOfStrings.Push("2")
		val, ok := myStackOfStrings.Pop()
		AssertFalse(t, myStackOfStrings.IsEmpty())
		AssertTrue(t, ok)
		AssertEqual(t, val, "2")
		val, ok = myStackOfStrings.Pop()
		AssertTrue(t, ok)
		AssertEqual(t, val, "1")
		AssertTrue(t, myStackOfStrings.IsEmpty())
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("did not want %+v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
