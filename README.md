Learning how to code in Go
========


Testing
--------
Create a file XXX_test.go in the same directory and package than SUT.

This file must import "testing" package.

Every test function must have the name TestXXXX and receive a pointer to *testing.T

To run

    go test [package]

Package is optional if you are in that package    

Reference: [http://golang.org/pkg/testing/](http://golang.org/pkg/testing/)

