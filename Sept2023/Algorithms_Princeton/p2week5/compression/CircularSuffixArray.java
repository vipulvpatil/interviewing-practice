import java.util.Arrays;
import edu.princeton.cs.algs4.StdOut;

public class CircularSuffixArray {
  private final String original;
  private final Suffix[] circularSuffixArray;

  private class Suffix implements Comparable<Suffix> {
    private final String s;
    private final int i;

    public Suffix(String s, int i) {
      this.s = s;
      this.i = i;
    }

    public String string() {
      return s;
    }

    public int index() {
      return i;
    }

    public int compareTo(Suffix other) {
      return this.s.compareTo(other.s);
    }
  }

  // circular suffix array of s
  public CircularSuffixArray(String s) {
    if (s == null) {
      throw new IllegalArgumentException();
    }
    original = s;
    if (original.length() == 0) {
      circularSuffixArray = new Suffix[] {};
      return;
    }
    circularSuffixArray = new Suffix[original.length()];
    circularSuffixArray[0] = new Suffix(s, 0);
    for (int i = 1; i < original.length(); i++) {
      String prevSuffixString = circularSuffixArray[i - 1].string();
      String nextSuffixString = circularShiftLeft(prevSuffixString);
      circularSuffixArray[i] = new Suffix(nextSuffixString, i);
    }

    Arrays.sort(circularSuffixArray);
  }

  // length of s
  public int length() {
    return original.length();
  }

  // returns index of ith sorted suffix
  public int index(int i) {
    if (i >= length()) {
      throw new IllegalArgumentException();
    }
    return circularSuffixArray[i].index();
  }

  private String circularShiftLeft(String s) {
    char[] shiftedCharArray = new char[s.length()];
    char temp = s.charAt(0);
    for (int i = 0; i < s.length() - 1; i++) {
      shiftedCharArray[i] = s.charAt(i + 1);
    }
    shiftedCharArray[s.length() - 1] = temp;
    return new String(shiftedCharArray);
  }

  // unit testing (required)
  public static void main(String[] args) {
    CircularSuffixArray c = new CircularSuffixArray("ABRACADABRA!");
    if (c.length() != 12) {
      StdOut.printf("length of CircularSuffixArray should be 12 but is %d\n", c.length());
    }
    int[] expectedIndices = { 11, 10, 7, 0, 3, 5, 8, 1, 4, 6, 9, 2 };

    for (int i = 0; i < expectedIndices.length; i++) {
      if (c.index(i) != expectedIndices[i]) {
        StdOut.printf("index of %d should be %d but is %d\n", i, expectedIndices[i], c.index(i));
      }
    }
  }

}