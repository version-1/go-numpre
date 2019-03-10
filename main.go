package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Result struct {
	N         int
	X         int
	Y         int
	Ok        bool
	Value     int
	Candidate []int
	Field     [][]int
}

func main() {
	const size = 9
	flag.Parse()
	numStr := flag.Arg(0)
	field := load(numStr, size)

	fmt.Println("")
	fmt.Println("[START]")
	result := Result{Field: field}
	printResult(result, size, false)
	n := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			n++
			cursorValue := field[i][j]
			if cursorValue == 0 {
				candidate := buildCandidate(size)
				value, ok, filteredCandidate := filter(field, j, i, size, candidate)
				if ok {
					field[i][j] = value
				}
				result := Result{N: n, X: j, Y: i, Ok: ok, Value: value, Candidate: filteredCandidate, Field: field}
				printResult(result, size, true)
			}
		}
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

func printResult(r Result, size int, header bool) {
	if header {
		fmt.Println("")
		fmt.Println("n :", r.N)
		fmt.Println("(x, y) :", r.X+1, r.Y+1)
		fmt.Println("ok :", r.Ok)
		fmt.Println("value :", r.Value)
		fmt.Println("candidate :", r.Candidate)
	}
	for i := 0; i < size; i++ {
		fmt.Println(r.Field[i])
	}
}

func buildCandidate(size int) []int {
	candidate := []int{}
	for i := 0; i < size; i++ {
		candidate = append(candidate, i+1)
	}
	return candidate
}

func filter(field [][]int, x int, y int, size int, candidate []int) (int, bool, []int) {
	_candidate := make([]int, len(candidate))
	copy(_candidate, candidate)
	_candidate = filterVertical(field, x, y, size, _candidate)
	if len(_candidate) == 1 {
		return _candidate[0], true, _candidate
	}
	_candidate = filterHorizontal(field, x, y, size, _candidate)
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
	for i := 0; i < size; i++ {
		value := field[i][x]
		if y != i && value != 0 {
			_candidate = remove(_candidate, value)
		}
	}
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
