package wordwrap_test

import (
	"fmt"
	"github.com/andreyvit/wordwrap"
)

const (
	lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`
)

func ExampleWrapSlice() {
	for _, line := range wordwrap.WrapSlice(lorem, 40, wordwrap.Options{}) {
		fmt.Println(line)
	}
	// Output:
	// Lorem ipsum dolor sit amet, consectetur
	// adipiscing elit, sed do eiusmod tempor
	// incididunt ut labore et dolore magna
	// aliqua. Ut enim ad minim veniam, quis
	// nostrud exercitation ullamco laboris
	// nisi ut aliquip ex ea commodo consequat.
	// Duis aute irure dolor in reprehenderit
	// in voluptate velit esse cillum dolore eu
	// fugiat nulla pariatur. Excepteur sint
	// occaecat cupidatat non proident, sunt in
	// culpa qui officia deserunt mollit anim
	// id est laborum.
}

func ExampleWrap() {
	wordwrap.Wrap(lorem, 40, wordwrap.Options{}, func(line string) {
		fmt.Println(line)
	})
	// Output:
	// Lorem ipsum dolor sit amet, consectetur
	// adipiscing elit, sed do eiusmod tempor
	// incididunt ut labore et dolore magna
	// aliqua. Ut enim ad minim veniam, quis
	// nostrud exercitation ullamco laboris
	// nisi ut aliquip ex ea commodo consequat.
	// Duis aute irure dolor in reprehenderit
	// in voluptate velit esse cillum dolore eu
	// fugiat nulla pariatur. Excepteur sint
	// occaecat cupidatat non proident, sunt in
	// culpa qui officia deserunt mollit anim
	// id est laborum.
}

func ExampleWrapString() {
	fmt.Print(wordwrap.WrapString(lorem, 40, wordwrap.Options{}))
	// Output:
	// Lorem ipsum dolor sit amet, consectetur
	// adipiscing elit, sed do eiusmod tempor
	// incididunt ut labore et dolore magna
	// aliqua. Ut enim ad minim veniam, quis
	// nostrud exercitation ullamco laboris
	// nisi ut aliquip ex ea commodo consequat.
	// Duis aute irure dolor in reprehenderit
	// in voluptate velit esse cillum dolore eu
	// fugiat nulla pariatur. Excepteur sint
	// occaecat cupidatat non proident, sunt in
	// culpa qui officia deserunt mollit anim
	// id est laborum.
}
