package main

import (
  "fmt"
)

func GetHello() string {
  return "Hello World From Admin!"
}

func main() {
  output := GetHello()
  fmt.Println(output)
}
