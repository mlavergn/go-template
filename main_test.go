package main

import (
	"testing"
)

func TestFoo(t *testing.T) {
	var ok = true
	if !ok {
		t.Fatalf("Foo test failed")
	}
}

func TestMain(t *testing.T) {
	var ok = true
	if !ok {
		t.Fatalf("Setup test failed")
	}
}
