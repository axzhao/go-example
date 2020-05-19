package error

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func foo1() error {
	return sql.ErrNoRows
}

func bar1() error {
	return foo1()
}

func ExampleError1() {
	err := bar1()
	if err != nil {
		fmt.Printf("got err, %+v\n", err)
	}

	// Output:
	// got err, sql: no rows in result set
}

func ExampleError2() {
	err := bar1()
	// if foo func return fmt.Errorf("foo err, %v", sql.ErrNoRows)
	if err == sql.ErrNoRows {
		fmt.Printf("data not found, %+v\n", err)
		return
	}
	if err != nil {
		// Unknown error
	}

	// Output:
	// data not found, sql: no rows in result set
}

func foo() error {
	return errors.Wrap(sql.ErrNoRows, "foo failed")
	// return errors.WithStack(sql.ErrNoRows)
}

func bar() error {
	return errors.WithMessage(foo(), "bar failed")
}

func ExampleError3() {
	err := bar()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found, %v\n", err)
		// fmt.Printf("%+v\n", err)
		return
	}
	if err != nil {
		// unknown error
	}

	// Output:
}
