/**
 * Caclulates and displays FizzBuzz Challenge.
 */
package fizzbuzz;

/**
 * Importing the version of the package. its not really needed, so I commeneted
 * it out.
 */
//import sample.version.Version;

/**
 * The Public Class FizzBuzz. Accessable by other packages using this package.
 */
public class FizzBuzz {
  /**
   * Performs the FizzBuzz Challenge.
   *
   * @param [int] num
   * @return [String] output
   */
  public static String fizzbuzz(int num) {
    String output;
    if( num % 3 == 0 && num % 5 ==0 && num != 0 ) {
      output = "FizzBuzz";
    } else if( num % 3 == 0 && num != 0) {
      output = "Fizz";
    } else if( num % 5 == 0 && num != 0) {
      output = "Buzz";
    } else {
      output = String.format("%s", num);
    }
    
    return output;
  }
  
  /**
   * The main entry point for the FizzBuzz application.
   *
   * @param [String] args
   * @return [Int] 0 
   */
  public static void main(String[] args) {
    for(int num = 0; num <= 100; num = num + 1) {
      System.out.println(fizzbuzz(num));
    }
  }
}