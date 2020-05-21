package error

import (
	"fmt"

	"github.com/pkg/errors"
)

// 1. String-based errors

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

// 2. Custom errors with data

type SyntaxError struct {
	Line int
	Col  int
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("%d:%d: syntax error", e.Line, e.Col)
}

type InternalError struct {
	Path string
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("parse %v: internal error", e.Path)
}

func Foo() error {
	return &InternalError{Path: ""}
}

func ExampleCustomError() {
	if err := Foo(); err != nil {
		switch e := err.(type) {
		case *SyntaxError:
			fmt.Println("Syntax")
			// Do something interesting with e.Line and e.Col.
		case *InternalError:
			fmt.Println("Internal")
			// Abort and file an issue.
		default:
			fmt.Println(e)
		}
	}

	// Output:
}

// 3.

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
	if e, ok := err.(errDuplicateKey); ok {
		// ok
		return e.DuplicateKey()
	}
	// way 2:
	err = errors.Cause(err)
	if e, ok := err.(errDuplicateKey); ok {
		// ok
		return e.DuplicateKey()
	}

	return false
}

func ExampleCustomError2() {
	err := &errDuplicateKey{error: errors.Errorf("document %s is exist", "123")}
	if IsErrDuplicateKey(err) {
		fmt.Println("Dup")
	}
	// Output:
	// Dup
}
