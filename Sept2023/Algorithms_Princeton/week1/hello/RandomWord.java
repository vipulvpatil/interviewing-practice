// StdIn.readString(): reads and returns the next string from standard input.
// StdIn.isEmpty(): returns true if there are no more strings available on standard input, and false otherwise.
// StdOut.println(): prints a string and terminating newline to standard output. It’s also fine to use System.out.println() instead.
// StdRandom.bernoulli(p)

import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;
import edu.princeton.cs.algs4.StdRandom;

public class RandomWord {
  public static void main(String[] args) {
    String champion = "";
    double index = 0;
    while (!StdIn.isEmpty()) {
      String nextWord = StdIn.readString();
      if (StdRandom.bernoulli(1 / (index + 1))) {
        champion = nextWord;
      }
      index++;
    }
    StdOut.println(champion);
  }
}
