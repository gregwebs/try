package assert

import (
	"fmt"
)

var (
	// P is a production Asserter that sets panic objects to errors which
	// allows err2 handlers to catch them.
	P = AsserterToError

	// D is a development Asserter that sets panic objects to strings that
	// doesn't by caught by err2 handlers.
	D Asserter = 0

	// DefaultAsserter is a default asserter used for package-level functions
	// like assert.That(). It is the same as the production asserter P, which
	// treats assert failures as Go errors, but in addition to that, it formats
	// the assertion message properly. Naturally, only if err2 handlers are
	// found in the call stack, these errors are caught.
	//
	// You are free to set it according to your current preferences. For
	// example, it might be better to panic about every assertion fault during
	// the tests. When in other cases, throw an error.
	DefaultAsserter = AsserterToError | AsserterFormattedCallerInfo
)

// That asserts that the term is true. If not it panics with the given
// formatting string. Note! That is the most performant of all the assertion
// functions.
func That(term bool, a ...any) {
	if !term {
		DefaultAsserter.reportAssertionFault("", a...)
	}
}

// NotNil asserts that the value is not nil. If it is it panics/errors (default
// Asserter) with the given message.
func NotNil[T any](p *T, a ...any) {
	if p == nil {
		defMsg := "pointer is nil"
		DefaultAsserter.reportAssertionFault(defMsg, a...)
	}
}

// SNotNil asserts that the slice is not nil. If it is it panics/errors (default
// Asserter) with the given message.
func SNotNil[T any](s []T, a ...any) {
	if s == nil {
		defMsg := "slice is nil"
		DefaultAsserter.reportAssertionFault(defMsg, a...)
	}
}

// CNotNil asserts that the channel is not nil. If it is it panics/errors
// (default Asserter) with the given message.
func CNotNil[T any](c chan T, a ...any) {
	if c == nil {
		defMsg := "channel is nil"
		DefaultAsserter.reportAssertionFault(defMsg, a...)
	}
}

// MNotNil asserts that the map is not nil. If it is it panics/errors (default
// Asserter) with the given message.
func MNotNil[T comparable, U any](m map[T]U, a ...any) {
	if m == nil {
		defMsg := "map is nil"
		DefaultAsserter.reportAssertionFault(defMsg, a...)
	}
}

// NotEqual asserts that the values are equal. If not it panics/errors (current
// Asserter) with the given message.
func NotEqual[T comparable](val, want T, a ...any) {
	if want == val {
		defMsg := fmt.Sprintf("got %v, want %v", val, want)
		DefaultAsserter.reportAssertionFault(defMsg, a...)
	}
}

// Equal asserts that the values are equal. If not it panics/errors (current
// Asserter) with the given message.
func Equal[T comparable](val, want T, a ...any) {
	if want != val {
		defMsg := fmt.Sprintf("got %v, want %v", val, want)
		DefaultAsserter.reportAssertionFault(defMsg, a...)
	}
}

// SLen asserts that the length of the slice is equal to given. If not it
// panics/errors (current Asserter) with the given message. Note! This is
// reasonable fast but not as fast as 'That' because of lacking inlining for the
// current implementation of Go's type parametric functions.
func SLen[T any](obj []T, length int, a ...any) {
	l := len(obj)

	if l != length {
		defMsg := fmt.Sprintf("got %d, want %d", l, length)
		DefaultAsserter.reportAssertionFault(defMsg, a...)
	}
}

// MLen asserts that the length of the map is equal to given. If not it
// panics/errors (current Asserter) with the given message. Note! This is
// reasonable fast but not as fast as 'That' because of lacking inlining for the
// current implementation of Go's type parametric functions.
func MLen[T comparable, U any](obj map[T]U, length int, a ...any) {
	l := len(obj)

	if l != length {
		defMsg := fmt.Sprintf("got %d, want %d", l, length)
		DefaultAsserter.reportAssertionFault(defMsg, a...)
	}
}

// NotEmpty asserts that the string is NOT empty. If not it panics/errors
// (current Asserter) with the given message. Note! This is reasonable fast but
// not as fast as 'That' because of lacking inlining for the current
// implementation of Go's type parametric functions.
func NotEmpty(obj string, a ...any) {
	if obj == "" {
		defMsg := "string shouldn't be empty"
		DefaultAsserter.reportAssertionFault(defMsg, a...)
	}
}

// SNotEmpty asserts that the slice is NOT empty. If not it panics/errors
// (current Asserter) with the given message. Note! This is reasonable fast but
// not as fast as 'That' because of lacking inlining for the current
// implementation of Go's type parametric functions.
func SNotEmpty[T any](obj []T, a ...any) {
	l := len(obj)

	if l == 0 {
		defMsg := "slice shouldn't be empty"
		DefaultAsserter.reportAssertionFault(defMsg, a...)
	}
}

// MNotEmpty asserts that the map is NOT empty. If not it panics/errors
// (current Asserter) with the given message. Note! This is reasonable fast but
// not as fast as 'That' because of lacking inlining for the current
// implementation of Go's type parametric functions.
func MNotEmpty[T comparable, U any](obj map[T]U, length int, a ...any) {
	l := len(obj)

	if l == 0 {
		defMsg := "map shouldn't be empty"
		DefaultAsserter.reportAssertionFault(defMsg, a...)
	}
}

func combineArgs(format string, a []any) []any {
	args := make([]any, 1, len(a)+1)
	args[0] = format
	args = append(args, a...)
	return args
}
