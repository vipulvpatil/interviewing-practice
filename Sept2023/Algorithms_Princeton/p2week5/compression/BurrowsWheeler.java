import edu.princeton.cs.algs4.BinaryStdIn;
import edu.princeton.cs.algs4.BinaryStdOut;
import edu.princeton.cs.algs4.Queue;

public class BurrowsWheeler {
  private static class IntegerQueue {
    private final Queue<Integer> queue;

    public IntegerQueue() {
      queue = new Queue<>();
    }

    public void enqueue(int n) {
      queue.enqueue(n);
    }

    public int dequeue() {
      return queue.dequeue();
    }

    public boolean isEmpty() {
      return queue.isEmpty();
    }
  }

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
    String s = BinaryStdIn.readString();
    char[] chars = s.toCharArray();
    IntegerQueue[] nextMap = new IntegerQueue[256];

    for (int i = 0; i < chars.length; i++) {
      char c = chars[i];
      if (nextMap[c] == null) {
        nextMap[c] = new IntegerQueue();
      }
      nextMap[c].enqueue(i);
    }
    int count = 0;
    char[] sortedT = new char[chars.length];
    int[] next = new int[sortedT.length];
    for (int i = 0; i < nextMap.length; i++) {
      char c = (char) i;
      if (nextMap[c] != null) {
        while (!nextMap[c].isEmpty()) {
          next[count] = nextMap[c].dequeue();
          sortedT[count++] = c;
        }
      }
    }

    int x = first;
    count = 0;
    chars = new char[sortedT.length];
    while (count < sortedT.length) {
      chars[count] = sortedT[x];
      x = next[x];
      count++;
    }
    BinaryStdOut.write(new String(chars));
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