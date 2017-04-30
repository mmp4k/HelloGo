package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	if "Hello World" != HelloWorld() {
		t.Fail()
	}
}
