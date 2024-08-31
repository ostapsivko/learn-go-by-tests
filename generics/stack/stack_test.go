package stack_test

import (
	"testing"
)

type StackOfInts struct {
	values []int
}

func (s *StackOfInts) Push(value int) {
	s.values = append(s.values, value)
}

func (s *StackOfInts) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackOfInts) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

type StackOfStrings struct {
	values []string
}

func (s *StackOfStrings) Push(value string) {
	s.values = append(s.values, value)
}

func (s *StackOfStrings) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackOfStrings) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}

func TestStack(t *testing.T) {
	t.Run("stack of ints", func(t *testing.T) {
		myStackOfInts := StackOfInts{}

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
	})

	t.Run("stack of strings", func(t *testing.T) {
		myStackOfStrings := StackOfStrings{}

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
