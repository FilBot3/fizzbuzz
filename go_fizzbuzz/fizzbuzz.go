package fizzbuzz

import (
  "fmt"
)

// Calculates multiples of 3, and 5, then prints Fizz on multiples 3 and prints
// Buzz on multiples of 5. Then on multiples of 3 and 5, print FizzBuzz.
func FizzBuzz(fizzBuzz int) (output string) {
  if (fizzBuzz % 3 == 0 && fizzBuzz % 5 == 0 && fizzBuzz != 0) {
    output = "FizzBuzz"
  } else if (fizzBuzz % 3 == 0 && fizzBuzz != 0) {
    output = "Fizz"
  } else if (fizzBuzz % 5 == 0 && fizzBuzz != 0) {
    output = "Buzz"
  } else {
    output = fmt.Sprintf("%v", fizzBuzz)
  }
  return output
}

func UseFizzBuzz() {
  fizzBuzz := 15
  for num := 0; num <= fizzBuzz; num++ {
    fmt.Println(FizzBuzz(num))
  }
}