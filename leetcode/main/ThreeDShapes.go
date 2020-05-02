package main

import (
	"fmt"
)

func surfaceArea(grid [][]int) int {
	area := 0
	for i := range grid {
		for j := range grid[i] {
			high := grid[i][j]
			if high > 0 {
				area += 2
				if i-1 >= 0 {
					area += max(high-grid[i-1][j], 0)
				} else {
					area += high
				}
				if i+1 < len(grid) {
					area += max(high-grid[i+1][j], 0)
				} else {
					area += high
				}
				if j-1 >= 0 {
					area += max(high-grid[i][j-1], 0)
				} else {
					area += high
				}
				if j+1 < len(grid[i]) {
					area += max(high-grid[i][j+1], 0)
				} else {
					area += high
				}
			}
		}
	}
	return area
}
func main() {
	fmt.Println(surfaceArea([][]int{{1, 2}, {3, 4}}))
}
