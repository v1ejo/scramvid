package assert

import (
	"reflect"
	"testing"
)

func Equal[T comparable](t *testing.T, got, expected T) {
	t.Helper()

	if got != expected {
		t.Errorf("got: %v;expected: %v", got, expected)
	}
}

func NotEqual[T comparable](t *testing.T, got, expected T) {
	t.Helper()

	if got == expected {
		t.Errorf("got: %v;expected: %v", got, expected)
	}
}

func EqualSlice[T comparable](t *testing.T, got, expected []T) {
	t.Helper()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got: %v;expected: %v", got, expected)
	}
}
