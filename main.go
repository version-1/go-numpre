package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const size = 9
	flag.Parse()
	numStr := flag.Arg(0)
	x := 7
	y := 2
	field := load(numStr, size)
	cursorValue := field[y][x]
	candidate := []int{}
	for i := 0; i < size; i++ {
		candidate = append(candidate, i+1)
	}

	fmt.Println(cursorValue)
	printResult(field, size)
	// for i := 0; i < size; i++ {
	// 	for j := 0; j < size; j++ {
	// 		check(field, x, j, size, candidate)
	// 	}
	// }
	value, ok, filteredCandidate := filter(field, x, y, size, candidate)
	if ok {
		field[y][x] = value
	}

	fmt.Println(value, ok, filteredCandidate)
	printResult(field, size)
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

func printResult(field [][]int, rowLength int) {
	for i := 0; i < rowLength; i++ {
		fmt.Println(field[i])
	}
}

func filter(field [][]int, x int, y int, size int, candidate []int) (int, bool, []int) {
	_candidate := make([]int, len(candidate))
	copy(_candidate, candidate)
	_candidate = filterVertical(field, x, y, size, _candidate)
	if len(_candidate) == 1 {
		return _candidate[0], true, _candidate
	}
	_candidate = filterHorizontal(field, x, y, size, _candidate)
	fmt.Println(_candidate)
	if len(_candidate) == 1 {
		return _candidate[0], true, _candidate
	}
	_candidate = filterBox(field, x, y, size, _candidate)
	if len(_candidate) == 1 {
		return _candidate[0], true, _candidate
	}

	return 0, false, _candidate
}

func filterVertical(field [][]int, x int, y int, size int, candidate []int) []int {
	_candidate := make([]int, len(candidate))
	copy(_candidate, candidate)
	fmt.Println(_candidate)
	for i := 0; i < size; i++ {
		value := field[i][x]
		if y != i && value != 0 {
			fmt.Println(value)
			_candidate = remove(_candidate, value)
		}
	}
	fmt.Println(_candidate)
	return _candidate
}

func filterHorizontal(field [][]int, x int, y int, size int, candidate []int) []int {
	_candidate := make([]int, len(candidate))
	copy(_candidate, candidate)
	for i := 0; i < size; i++ {
		if x != i && field[y][i] != 0 {
			_candidate = remove(_candidate, field[y][i])
		}
	}
	return _candidate
}

func filterBox(field [][]int, x int, y int, size int, candidate []int) []int {
	const sectionSize = 3
	_candidate := make([]int, len(candidate))
	copy(_candidate, candidate)
	row := y / sectionSize
	col := x / sectionSize
	indexY := row * sectionSize
	indexX := col * sectionSize
	boxWidth := size / sectionSize

	for i := 0; i < boxWidth; i++ {
		for j := 0; j < boxWidth; j++ {
			_x := indexX + i
			_y := indexY + j
			value := field[_y][_x]
			if value != 0 || (y != _y || x != _x) {
				_candidate = remove(_candidate, value)
			}
		}
	}
	return _candidate
}

func remove(numbers []int, search int) []int {
	result := []int{}
	for _, num := range numbers {
		if num != search {
			result = append(result, num)
		}
	}
	return result
}
