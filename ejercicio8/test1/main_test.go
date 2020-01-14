package main

import "testing"

func TestSum(t *testing.T) {
	// arrage
	var v int
	var expect int = 4
	// act
	v = Sum(2, 2)
	// assert
	if v != expect {
		t.Error("expect 4, got", v)
	}
}
