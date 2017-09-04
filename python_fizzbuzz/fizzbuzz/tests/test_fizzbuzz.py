#!/usr/bin/env python
#
#
#

from unittest import TestCase

import fizzbuzz as fb

class TestFizzBuzz(TestCase):
    def test_fizzbuzz(self):
        output = fb.fizzbuzz(1)
        self.assertTrue(output, "1")
        
        output = fb.fizzbuzz(3)
        self.assertTrue(output, "Fizz")
        
        output = fb.fizzbuzz(5)
        self.assertTrue(output, "Buzz")
        
        output = fb.fizzbuzz(15)
        self.assertTrue(output, "FizzBuzz")
    # end test_fizzbuzz
# end TestFizzBuzz