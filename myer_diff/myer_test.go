package main

import (
	"strings"
)

func ExampleDiff() {
	a := strings.Split(`ABCABBAAABC`, ``)
	b := strings.Split(`CBABACAABBD`, ``)
	generateDiff(a, b)
	// Output:
	// -A
	// -B
	//  C
	// +B
	//  A
	//  B
	// -B
	//  A
	// +C
	//  A
	//  A
	//  B
	// -C
	// +B
	// +D
}
