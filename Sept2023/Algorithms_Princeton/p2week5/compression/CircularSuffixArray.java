import java.util.HashMap;

import edu.princeton.cs.algs4.MSD;
import edu.princeton.cs.algs4.StdOut;

public class CircularSuffixArray {
  String original;
  HashMap<String, Integer> indices;
  String[] circularSuffixArray;

  // circular suffix array of s
  public CircularSuffixArray(String s) {
    original = s;
    circularSuffixArray = new String[original.length()];
    circularSuffixArray[0] = s;
    indices = new HashMap<>();
    indices.put(circularSuffixArray[0], 0);
    for (int i = 1; i < original.length(); i++) {
      circularSuffixArray[i] = circularShiftLeft(circularSuffixArray[i - 1]);
      indices.put(circularSuffixArray[i], i);
    }

    MSD.sort(circularSuffixArray);
  }

  // length of s
  public int length() {
    return original.length();
  }

  // returns index of ith sorted suffix
  public int index(int i) {
    return indices.get(circularSuffixArray[i]);
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
    int[] expectedIndices = new int[] { 11, 10, 7, 0, 3, 5, 8, 1, 4, 6, 9, 2 };

    for (int i = 0; i < expectedIndices.length; i++) {
      if (c.index(i) != expectedIndices[i]) {
        StdOut.printf("index of %d should be %d but is %d\n", i, expectedIndices[i], c.index(i));
      }
    }
  }

}