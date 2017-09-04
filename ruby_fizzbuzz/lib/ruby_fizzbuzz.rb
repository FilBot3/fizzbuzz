# encoding: UTF-8
require 'ruby_fizzbuzz/version'

# Used to display the fizzbuzz challenge.
module RubyFizzbuzz
  # Displays Fizz, Buzz, and FizzBuzz for numbers divisible by 3, 5, and 3 and 5
  # @param [Integer] num
  # @return [String] output
  def self.fizzbuzz(num = 100)
    output = if (num % 3).zero? && (num % 5).zero? && num != 0
               'FizzBuzz'
             elsif (num % 3).zero? && num != 0
               'Fizz'
             elsif (num % 5).zero? && num != 0
               'Buzz'
             else
               num.to_s
             end
    output
  end
end
