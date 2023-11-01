import java.util.ArrayList;
import java.util.List;

import edu.princeton.cs.algs4.BinaryStdIn;
import edu.princeton.cs.algs4.BinaryStdOut;

public class MoveToFront {

  // apply move-to-front encoding, reading from standard input and writing to
  // standard output
  public static void encode() {
    char[] list = extendedAsciiList();
    List<Integer> ints = new ArrayList<>();
    while (!BinaryStdIn.isEmpty()) {
      char input = BinaryStdIn.readChar();
      int i = bringToFront(list, input);
      ints.add(i);
    }
    for (int i = 0; i < ints.size(); i++) {
      BinaryStdOut.write(ints.get(i), 8);
    }
    BinaryStdOut.close();
  }

  // apply move-to-front decoding, reading from standard input and writing to
  // standard output
  public static void decode() {
    char[] list = extendedAsciiList();
    List<String> chars = new ArrayList<>();
    while (!BinaryStdIn.isEmpty()) {
      int input = BinaryStdIn.readInt(8);
      char c = bringToFront(list, input);
      chars.add(String.valueOf(c));
    }
    for (int i = 0; i < chars.size(); i++) {
      char c = chars.get(i).charAt(0);
      BinaryStdOut.write(c);
    }
    BinaryStdOut.close();
  }

  private static char[] extendedAsciiList() {
    char[] list = new char[256];
    for (int i = 0; i < list.length; i++) {
      list[i] = (char) i;
    }
    return list;
  }

  private static int bringToFront(char[] chars, char x) {
    char prev;
    if (chars[0] == x) {
      return 0;
    }
    prev = chars[0];
    for (int i = 1; i < chars.length; i++) {
      char curr = chars[i];
      chars[i] = prev;
      if (curr == x) {
        chars[0] = curr;
        return i;
      }
      prev = curr;
    }
    return -1;
  }

  private static char bringToFront(char[] chars, int x) {
    char charX = chars[x];
    for (int i = x; i > 0; i--) {
      chars[i] = chars[i - 1];
    }
    chars[0] = charX;
    return charX;
  }

  // if args[0] is "-", apply move-to-front encoding
  // if args[0] is "+", apply move-to-front decoding
  public static void main(String[] args) {
    if (args == null || args.length == 0) {
      throw new IllegalArgumentException();
    }
    if ("-".equals(args[0])) {
      encode();
    } else if ("+".equals(args[0])) {
      decode();
    }
  }

}
