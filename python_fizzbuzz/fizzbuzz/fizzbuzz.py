#!/usr/bin/env
#
#
#

def fizzbuzz( num = 100 ):
    '''
    Calculate the modulus of 3's and 5's printing Fizz on 3's and Buzz on 5's
    then printing FizzBuzz on modulus of 3 and 5.
    Returns the output of either Fizz, Buzz, or FizzBuzz. 
    '''
    if num % 3 == 0 and num != 0 :
        output = "Fizz"
    elif num % 5 == 0 and num != 0 :
        output = "Buzz"
    elif num % 3 == 0 and num % 5 == 0 and num != 0 :
        output = "FizzBuzz"
    else:
        output = "%s", (num)
    # end of if...elif...end
    return output
# end of fizzbuzz