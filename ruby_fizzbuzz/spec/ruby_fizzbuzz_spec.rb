# encoding: UTF-8
require 'spec_helper'

describe RubyFizzbuzz do
  it 'must return a string if no matches' do
    RubyFizzbuzz.fizzbuzz(1).must_equal '1'
  end
  
  it 'must return Fizz when modulus 3' do
    RubyFizzbuzz.fizzbuzz(3).must_equal 'Fizz'
  end
  
  it 'must return Buzz when modulus 5' do
    RubyFizzbuzz.fizzbuzz(5).must_equal 'Buzz'
  end
  
  it 'must return FizzBuzz when modulus 3 and 5' do
    RubyFizzbuzz.fizzbuzz(15).must_equal 'FizzBuzz'
  end
end