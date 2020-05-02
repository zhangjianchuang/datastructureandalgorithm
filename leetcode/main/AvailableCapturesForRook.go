package main

func main() {
	// 65 - 90 是大写 A - Z
	// 97 - 122 是小写 a - z
	println(numRookCaptures([][]byte{{'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', 'R', '.', '.', '.', 'p'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', 'p', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.'}}))
}

func numRookCaptures(board [][]byte) int {
	count := 0
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				for k := 0; k < 4; k++ {
					x := i
					y := j
					for true {
						x += dx[k]
						y += dy[k]
						if x < 0 || x >= 8 || y < 0 || y >= 8 || board[x][y] == 'B' {
							break
						}
						if board[x][y] == 'p' {
							count++
							break
						}
					}
				}
				return count
			}
		}
	}
	return count
}
