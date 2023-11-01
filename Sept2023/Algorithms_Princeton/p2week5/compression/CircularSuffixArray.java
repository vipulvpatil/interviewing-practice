import java.util.Arrays;
import edu.princeton.cs.algs4.StdOut;

public class CircularSuffixArray {
  private final String original;
  private final Suffix[] circularSuffixArray;

  private class Suffix implements Comparable<Suffix> {
    private final String s;
    private final int offset;

    public Suffix(String s, int offset) {
      this.s = s;
      this.offset = offset;
    }

    private char charAt(int i) {
      return s.charAt((offset + i) % s.length());
    }

    public int index() {
      return offset;
    }

    public int compareTo(Suffix other) {
      for (int i = 0; i < s.length(); i++) {
        int comp = Character.compare(this.charAt(i), other.charAt(i));
        if (comp != 0) {
          return comp;
        }
      }
      return 0;
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
    for (int i = 0; i < original.length(); i++) {
      circularSuffixArray[i] = new Suffix(s, i);
    }

    Arrays.sort(circularSuffixArray);
  }

  // length of s
  public int length() {
    return original.length();
  }

  // returns index of ith sorted suffix
  public int index(int i) {
    if (i >= length() || i < 0) {
      throw new IllegalArgumentException();
    }
    return circularSuffixArray[i].index();
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