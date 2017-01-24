package main

import (
  "fmt"
  "os"
)

func main() {
  args := os.Args

  if len(args) > 1 {
    for index, arg := range(args) {
      if index > 0 {
        fmt.Println(arg)
      }
    }
  } else {
    fmt.Println("No args were passed!")
  }
}
