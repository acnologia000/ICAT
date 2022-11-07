package main

func findWXYZ(input int) (int, int, int, int) {

	return 0, 0, 0, 0
}

func allFactorCombinations([]int) [][]int {

	return nil
}

func factorize(input int) []int {
	out := make([]int, 0)

	for i := 1; i < input-1; i++ {
		if input%i == 0 {
			out = append(out, i)
		}
	}
	return out
}
