package error

import (
	"fmt"

	"github.com/pkg/errors"
)

func ExampleStringBased() {
	// simple string-based error
	err1 := errors.New("math: square root of negative number")
	// with formatting
	err2 := fmt.Errorf("math: square root of negative number %g", -1.1)

	fmt.Println(err1)
	fmt.Println(err2)

	// Output:
	// math: square root of negative number
	// math: square root of negative number -1.1
}

type errDuplicateKey struct {
	error
}

// DuplicateKey return true
func (e *errDuplicateKey) DuplicateKey() bool {
	return true
}

// IsErrDuplicateKey return true
func IsErrDuplicateKey(err error) bool {
	type errDuplicateKey interface {
		DuplicateKey() bool
	}
	// way 1:
	if _, ok := err.(errDuplicateKey); ok {
		// ok
	}
	// way 2:
	err = errors.Cause(err)
	if e, ok := err.(errDuplicateKey); ok {
		// ok
		return e.DuplicateKey()
	}

	return false
}

func ExampleCustomError() {
	err := &errDuplicateKey{error: errors.Errorf("document %s is exist", "123")}
	if IsErrDuplicateKey(err) {
		fmt.Println("Dup")
	}
	// Output:
	// Dup
}
