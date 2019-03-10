package main

import (
  "fmt"
  "flag"
  "strings"
  "strconv"
)

func main() {
  const size = 9
  flag.Parse()
  numStr := flag.Arg(0)
  field := load(numStr, size)

  printResult(field, size)
}

func load(numStr string, size int) [][]int {
  field := [][]int{}
  for i := 0; i < size ; i++ {
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
