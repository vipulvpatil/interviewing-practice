import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

import edu.princeton.cs.algs4.BinaryStdIn;
import edu.princeton.cs.algs4.BinaryStdOut;
import edu.princeton.cs.algs4.MSD;
import edu.princeton.cs.algs4.Queue;
import edu.princeton.cs.algs4.StdOut;

public class BurrowsWheeler {

  // apply Burrows-Wheeler transform,
  // reading from standard input and writing to standard output
  public static void transform() {
    String original = BinaryStdIn.readString();
    CircularSuffixArray circularSuffixArray = new CircularSuffixArray(original);
    int first = 0;
    char[] t = new char[original.length()];
    for (int i = 0; i < circularSuffixArray.length(); i++) {
      int j = circularSuffixArray.index(i);
      if (j == 0) {
        first = i;
        t[i] = original.charAt(circularSuffixArray.length() - 1);
      } else {
        t[i] = original.charAt(j - 1);
      }
    }
    BinaryStdOut.write(first);
    BinaryStdOut.write(new String(t), 8);
    BinaryStdOut.close();
  }

  // apply Burrows-Wheeler inverse transform,
  // reading from standard input and writing to standard output
  public static void inverseTransform() {
    int first = BinaryStdIn.readInt();
    List<String> t = new ArrayList<>();
    HashMap<String, Queue<Integer>> nextMap = new HashMap<>();
    int length = 0;
    while (!BinaryStdIn.isEmpty()) {
      char c = BinaryStdIn.readChar(8);
      String s = String.valueOf(c);
      t.add(s);
      if (!nextMap.containsKey(s)) {
        nextMap.put(s, new Queue<>());
      }
      nextMap.get(s).enqueue(length);
      length++;
    }
    String[] sortedT = t.toArray(new String[length]);
    MSD.sort(sortedT);

    int[] next = new int[length];
    for (int i = 0; i < length; i++) {
      int index = nextMap.get(sortedT[i]).dequeue();
      next[i] = index;
    }

    int x = first;
    int count = 0;
    StringBuilder sb = new StringBuilder();
    while (count < length) {
      sb.append(sortedT[x]);
      x = next[x];
      count++;
    }
    BinaryStdOut.write(sb.toString());
    BinaryStdOut.close();
  }

  // if args[0] is "-", apply Burrows-Wheeler transform
  // if args[0] is "+", apply Burrows-Wheeler inverse transform
  public static void main(String[] args) {
    if (args == null || args.length == 0) {
      throw new IllegalArgumentException();
    }
    if ("-".equals(args[0])) {
      transform();
    } else if ("+".equals(args[0])) {
      inverseTransform();
    }
  }
}