import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdRandom;

public class Permutation {
  public static void main(String[] args) {
    int k = Integer.parseInt(args[0]);
    int included = 0;
    RandomizedQueue<String> q = new RandomizedQueue<String>();
    while (!StdIn.isEmpty()) {
      String str = StdIn.readString();
      if (k > 0) {
        if (included < k) {
          q.enqueue(str);
          included++;
        } else if (StdRandom.bernoulli(0.5)) {
          q.dequeue();
          q.enqueue(str);
        }
      }
    }
    while (k > 0) {
      System.out.println(q.dequeue());
      k--;
    }
  }
}
