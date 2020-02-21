package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Fatal("Add(1, 2) must be 3 but was not")
	}
}
