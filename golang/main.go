package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func render(grid *[][]int) {
	for i := 0; i < 9; i++ {
		var row string
		for _, v := range (*grid)[i] {
			row += strconv.Itoa(v) + ", "
		}
		fmt.Printf("[ %s ] \n", row[0:25])
	}
}

func load(numStr string, size int) [][]int {
	field := [][]int{}
	for i := 0; i < size; i++ {
		start := i * size
		end := start + size
		array := []int{}
		row := numStr[start:end]
		for _, v := range strings.Split(row, "") {
			value, _ := strconv.Atoi(v)
			array = append(array, value)
		}
		field = append(field, array)
	}

	return field
}

func validate(grid *[][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			n := (*grid)[i][j]
			if n == 0 || !possible(grid, i, j, n) {
				return false
			}
		}
	}

	return true
}

func possible(grid *[][]int, y int, x int, n int) bool {
	for i := 0; i < 9; i++ {
		if (i != x && (*grid)[y][i] == n) || (i != y && (*grid)[i][x] == n) {
			return false
		}
	}

	x0 := int(math.Floor(float64(x)/3) * 3)
	y0 := int(math.Floor(float64(y)/3) * 3)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if x != x0+j && y != y0+i && (*grid)[y0+i][x0+j] == n {
				return false
			}
		}
	}

	return true
}

func solve(grid *[][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if (*grid)[i][j] == 0 {
				for k := 0; k < 9; k++ {
					if possible(grid, i, j, k+1) {
						(*grid)[i][j] = k + 1
						if solve(grid) {
							return true
						} else {
							(*grid)[i][j] = 0
						}
					}
				}
				return false
			}
		}
	}

	return true
}

func main() {
	const size = 9
	flag.Parse()
	numStr := flag.Arg(0)
	field := load(numStr, size)

	fmt.Println("[START]")
	render(&field)
	fmt.Println("")

	solve(&field)
	if validate(&field) {
		fmt.Println("[END]")
	} else {
		fmt.Println("[END] Invalid!!!!!!!!!!")
	}

	render(&field)
}
