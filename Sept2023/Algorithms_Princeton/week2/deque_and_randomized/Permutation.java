import edu.princeton.cs.algs4.StdIn;

public class Permutation {
  public static void main(String[] args) {
    int k = Integer.parseInt(args[0]);
    RandomizedQueue<String> q = new RandomizedQueue<String>();
    while (!StdIn.isEmpty()) {
      String str = StdIn.readString();
      q.enqueue(str);
    }
    while (k > 0) {
      System.out.println(q.dequeue());
      k--;
    }
  }
}
