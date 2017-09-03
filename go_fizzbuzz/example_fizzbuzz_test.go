package fizzbuzz_test

import (
  "fmt"
  fb "github.com/predatorian3/fizzbuzz/go_fizzbuzz"
)

func ExampleFizzBuzz() {
  for fizzBuzz := 0; fizzBuzz <= 15; fizzBuzz++ {
    fmt.Println(fb.FizzBuzz(fizzBuzz))
  }
  // Output:
  // 0
  // 1
  // 2
  // Fizz
  // 4
  // Buzz
  // Fizz
  // 7
  // 8
  // Fizz
  // Buzz
  // 11
  // Fizz
  // 13
  // 14
  // FizzBuzz
}