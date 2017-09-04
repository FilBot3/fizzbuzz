/**
 *
 */
package fizzbuzz;

import fizzbuzz.FizzBuzz;
import org.junit.Assert;
import org.junit.Before;
import org.junit.Test;

public class FizzBuzzTest {
  private FizzBuzz fizz_buzz;
  
  @org.junit.Before
  public void beforeEach() {
    fizz_buzz = new FizzBuzz();
  }
  
  @org.junit.Test
  public void testFizzBuzz() {
    Assert.assertEquals("Fizz", fizz_buzz.fizzbuzz(3));
    Assert.assertEquals("Buzz", fizz_buzz.fizzbuzz(5));
    Assert.assertEquals("FizzBuzz", fizz_buzz.fizzbuzz(15));
    Assert.assertEquals("1", fizz_buzz.fizzbuzz(1));
  }
}