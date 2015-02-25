Learning how to test in Go.

Create a file XXX_test.go in the same directory and package than SUT.
This file must import "testing" package.
Every test function must have the name TestXXXX and receive a pointer to *testing.T
To run
	go test -c
	go test

Reference: http://golang.org/pkg/testing/

