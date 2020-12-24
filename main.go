package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func spiralize(mx [][]int) []int { // check nxn
	spiralLen := len(mx) * len(mx)
	res := make([]int, 0, spiralLen)

	if len(mx) != len(mx[0]) {
		fmt.Println("Only square matricies are supported")
		return res
	}

	var i, j, minn, ii, jj, nn int
	jj = 1
	nn = len(mx) - 1
	for k := 0; k < spiralLen; k++ {
		res = append(res, mx[i][j])
		i += ii
		j += jj
		if k == 0 {
			continue
		}
		if i == j && i < nn { // top-left corner
			nn--
			minn++
			j = minn
			i = j
			ii, jj = jj, ii
			ii *= -1
			jj *= -1
		} else if max(i, j) == nn && (min(i, j) == minn || min(i, j) == nn) { // all other corners
			ii, jj = jj, ii
			if i == j && j == nn {
				ii *= -1
				jj *= -1
			}
		}
	}
	return res
}

func main() {
	input1 := [][]int{{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}

	input2 := [][]int{{1, 2, 3, 1},
		{4, 5, 6, 4},
		{7, 8, 9, 7},
		{7, 8, 9, 7}}
	fmt.Println(spiralize(input1))
	fmt.Println(spiralize(input2))
}
