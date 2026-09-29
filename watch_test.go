package main

import "testing"

func TestDelta(t *testing.T) {
	if _, ok := Delta(0, 5, false); ok {
		t.Fatal("first reading is not a change")
	}
	if d, ok := Delta(5, 7, true); !ok || d != 2 {
		t.Fatalf("got %d %v", d, ok)
	}
	if _, ok := Delta(7, 7, true); ok {
		t.Fatal("no change")
	}
}
