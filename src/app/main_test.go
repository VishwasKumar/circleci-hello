package main

import "testing"

func TestHelloFunc(t *testing.T) {
  t.Log("Checking that function prints what's expected...")
  result := GetHello()
  if result != "Hello World From App!!" {
    t.Error("Expected 'Hello World From App!!' but got", result)
  }
}
