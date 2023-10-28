package utils

import "encoding/json"

// Result is a generic interface that can either hold a value or an error.
// It is used to return values from functions that can either return a value or an error.
// It has two types that implement it: ok[T] and err[T].
type Result[T any] interface {
	IsOK() bool
	IsErr() bool
	Unwrap() T
	UnwrapErr() error
	MarshalJSON() ([]byte, error)
}

// ok is a type that implements Result[T] interface. It is used to represent a value.
type ok[T any] struct {
	value T
}

// Ok returns a Result[T] that holds the given value.
func Ok[T any](value T) Result[T] {
	return ok[T]{value: value}
}

// IsOK returns true, since ok[T] always holds a value.
func (o ok[T]) IsOK() bool {
	return true
}

// IsErr returns false, since ok[T] never holds an error.
func (o ok[T]) IsErr() bool {
	return false
}

// Unwrap returns the value that ok[T] holds.
func (o ok[T]) Unwrap() T {
	return o.value
}

// UnwrapErr panics, since ok[T] never holds an error that could be unwrapped.
func (o ok[T]) UnwrapErr() error {
	panic("called UnwrapErr on ok[T]")
}

// MarshalJSON marshals the value that ok[T] holds into a JSON string.
func (o ok[T]) MarshalJSON() ([]byte, error) {
	value, err := json.Marshal(o.value)
	if err != nil {
		return nil, err
	}
	return []byte(`{"Ok": true, "Value": ` + string(value) + `}`), nil
}

// err is a type that implements Result[T] interface. It is used to represent an error.
type err[T any] struct {
	value error
}

// Err returns a Result[T] that holds the given error.
func Err[T any](value error) Result[T] {
	return err[T]{value: value}
}

// IsOK returns false, since err[T] never holds a value.
func (e err[T]) IsOK() bool {
	return false
}

// IsErr returns true, since err[T] always holds an error.
func (e err[T]) IsErr() bool {
	return true
}

// Unwrap panics, since err[T] never holds a value that could be unwrapped.
func (e err[T]) Unwrap() T {
	panic("called Unwrap on err[T]")
}

// UnwrapErr returns the error that err[T] holds.
func (e err[T]) UnwrapErr() error {
	return e.value
}

// MarshalJSON marshals the error that err[T] holds into a JSON string.
func (e err[T]) MarshalJSON() ([]byte, error) {
	value, err := json.Marshal(e.value.Error())
	if err != nil {
		return nil, err
	}
	return []byte(`{"Ok": false, "Error": ` + string(value) + `}`), nil
}
