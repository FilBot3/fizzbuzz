package fizzbuzz

import (
  "testing"
)

func TestFizzBuzz(t *testing.T) {
  var tests = []struct{
    in int
    expected string
  }{
    {1, "1"},
    {2, "2"},
    {3, "Fizz"},
    {4, "4"},
    {5, "Buzz"},
    {6, "Fizz"},
    {7, "7"},
    {8, "8"},
    {9, "Fizz"},
    {10, "Buzz"},
    {11, "11"},
    {12, "Fizz"},
    {13, "13"},
    {14, "14"},
    {15, "FizzBuzz"},
  }
  
  for _, tt := range tests {
    output := FizzBuzz(tt.in)
    if output != tt.expected {
      t.Errorf("FizzBuzz gave %q, wanted %q", output, tt.expected)
    }
  }
}