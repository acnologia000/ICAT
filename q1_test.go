package main

import "testing"

func TestFactorize(t *testing.T) {
	ans := []int{1, 2, 4}
	out := factorize(8)

	for i := 0; i < len(out); i++ {
		if out[i] != ans[i] {
			t.Fatal("failed", out, ans)
		}
	}
}

func TestAllCombination(t *testing.T) {
	allFactorCombinations(factorize(8))
}
