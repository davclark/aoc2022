package main

import (
  "os"
  "fmt"
  "bufio"
  "strconv"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func maxKeeper() func(int) int {
  max_total := 0
  return func(x int) int {
    fmt.Println("Total", x)
    if x > max_total {
      fmt.Println("New max!")
      max_total = x
    }

    return max_total
  }
}

func main() {
  f, err := os.Open("input1.txt")
  check(err)
  scanner := bufio.NewScanner(f)

  curr_total := 0
  maxTotal := maxKeeper()
  for scanner.Scan() {
    text := scanner.Text()
    num, err := strconv.Atoi(text)
    if err != nil {
      _ = maxTotal(curr_total)
      curr_total = 0
    } else {
      fmt.Println(text)
      curr_total += num
    }
  }

  fmt.Println("Largest total:", maxTotal(curr_total))
}

