package result

import "fmt"

// Result represents a union of a successful value and an error message.
type Result[T any] struct {
	value T
	err   error
}

// Ok creates a Result with a value.
func Ok[T any](value T) Result[T] {
	return Result[T]{value: value, err: nil}
}

// Err creates a Result with an error.
func Err[T any](err error) Result[T] {
	return Result[T]{value: *new(T), err: err} // Assumes T is a pointer type
}

// IsOk checks if the Result is successful.
func (r *Result[T]) IsOk() bool {
	return r.err == nil
}

// Value retrieves the value if it's successful, otherwise returns the zero value.
func (r *Result[T]) Value() T {
	return r.value
}

// Error retrieves the error message if there's an error.
func (r *Result[T]) Error() error {
	return r.err
}

// Returns
func (r *Result[T]) Unwrap() (T, error) {
	if r.IsOk() {
		return r.value, nil
	}
	return r.value, r.err
}

// String provides a string representation of the Result.
func (r *Result[T]) String() string {
	if r.IsOk() {
		return fmt.Sprintf("Ok(%v)", r.value)
	}
	return fmt.Sprintf("Err(%s)", r.err)
}
