package main

import "testing"

func TestHelloFunc(t *testing.T) {
  t.Log("Checking that function prints what's expected...")
  result := GetHello()
  if result != "Hello World From Admin!" {
    t.Error("Expected 'Hello World From Admin!' but got", result)
  }
}
