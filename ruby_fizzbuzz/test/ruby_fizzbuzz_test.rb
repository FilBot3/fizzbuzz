# encoding: UTF-8
require "test_helper"

module RubyFizzbuzz
  class Test < Minitest::Test
    def test_fizzbuzz_returns_number
      assert_equal "1", RubyFizzbuzz.fizzbuzz(1)
    end
    
    def test_fizzbuzz_returns_fizz
      assert_equal "Fizz", RubyFizzbuzz.fizzbuzz(3)
    end
    
    def test_fizzbuzz_returns_buzz
      assert_equal "Buzz", RubyFizzbuzz.fizzbuzz(5)
    end
    
    def test_fizzbuzz_returns_fizzbuzz
      assert_equal "FizzBuzz", RubyFizzbuzz.fizzbuzz(15)
    end
  end
end