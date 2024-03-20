package main

import (
	"fmt"
	"time"
)

func main() {
  a := make(chan int)
  b := make(chan int)

  go func() {
    defer close(a)
    for i := 0; i < 10; i++ {
      time.Sleep(2 * time.Second)
      fmt.Println(i)
    }
  }()

  go func() {
    defer close(b)
    for i := 0; i < 5; i++ {
      time.Sleep(1 * time.Second)
      fmt.Println(i)
    }
  }()

  select {
  case <- a:
    fmt.Println("Hello world!")
  case <- b:
    fmt.Println("Hello world!")
  }

  fmt.Println("Hello me!")
}
