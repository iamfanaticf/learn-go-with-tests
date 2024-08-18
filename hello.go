package main

import "fmt"

var englishHelloPrefix = "Hello, "

func Hello(name string) string {
  if name == "" {
    name = "world"
  }
  return englishHelloPrefix + name
}

func main() {
  fmt.Println(Hello("world"))
}
