package main

import (
  // #include "math.h"
  "C"
  "fmt"
)

func main() {
  fmt.Println("Welcome from Go")
  fmt.Println("Let's call some functions from our C library")
  fmt.Printf("C: math_add(2, 3)=%d\n", C.math_add(2, 3))
  fmt.Printf("C: math_sub(5, 2)=%d\n", C.math_sub(5, 2))
}
