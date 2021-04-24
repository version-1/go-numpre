package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Render(grid *[][]int) {
	for i := 0; i < 9; i++ {
		var row string
		for _, v := range (*grid)[i] {
			row += strconv.Itoa(v) + ", "
		}
		fmt.Printf("[ %s ] \n", row[0:25])
	}
}

func Load(numStr string, size int) [][]int {
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

func Possible(grid *[][]int, y int, x int, n int) bool {
	for i := 0; i < 9; i++ {
		if i != x && (*grid)[y][i] == n {
			return false
		}
	}

	for j := 0; j < 9; j++ {
		if j != x && (*grid)[j][x] == n {
			return false
		}
	}

	x0 := int(math.Floor(float64(x)/3) * 3)
	y0 := int(math.Floor(float64(y)/3) * 3)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (*grid)[y0+i][x0+j] == n {
				return false
			}
		}
	}

	return true
}

func Solve(grid *[][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if (*grid)[i][j] == 0 {
				for k := 0; k < 9; k++ {
					if Possible(grid, i, j, k+1) {
						(*grid)[i][j] = k + 1
						if Solve(grid) {
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
	field := Load(numStr, size)

	fmt.Println("[START]")
	Render(&field)

	Solve(&field)

	fmt.Println("[END]")
	Render(&field)
}
