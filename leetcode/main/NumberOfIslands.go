package main

func main() {
	// [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]
	println(numIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))
	println(numIslands([][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}))
	println(numIslands([][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}}))
}

func numIslands(grid [][]byte) int {
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				flag := false
				left := j == 0 || grid[i][j-1] == '0'
				top := i == 0 || grid[i-1][j] == '0'
				if left && top {
					flag = true
				}

				if flag {
					count++
				}
			}
		}
	}
	return count
}
