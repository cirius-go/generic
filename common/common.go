package common

// Zero returns a zero value of type T.
func Zero[T any]() T {
	var (
		t T
	)

	return t
}

// IsZero returns true if the value is zero.
func IsZero[T comparable](v T) bool {
	return v == Zero[T]()
}

// Something represents something that not important.
type Something struct{}

// SomethingIntf represents something that used to implement purpose only.
type SomethingIntf interface{}

// Pointer returns pointer of value T.
func Pointer[T any](value T) *T {
	return &value
}
