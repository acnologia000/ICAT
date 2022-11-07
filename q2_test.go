package main

import "testing"

func TestMoveYoshita(t *testing.T) {
	out := moveYoshita("LLRDDR", []int{0, 0})
	if out[0] != 0 || out[1] != -2 {
		t.Error("something went wrong")
	}
}
