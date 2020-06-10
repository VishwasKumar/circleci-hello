package main

import (
  "fmt"
)

func GetHello() string {
  return "Hello World From App!!!"
}

func main() {
  output := GetHello()
  fmt.Println(output)
}
