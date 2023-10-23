import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdOut;

public class Outcast {
  private final WordNet wn;

  // constructor takes a WordNet object
  public Outcast(WordNet wordnet) {
    wn = wordnet;
  }

  // given an array of WordNet nouns, return an outcast
  public String outcast(String[] nouns) {
    int maxDist = 0;
    String theOutcast = "";
    for (String n : nouns) {
      int d = 0;
      for (String m : nouns) {
        d += wn.distance(n, m);
      }
      if (d > maxDist) {
        maxDist = d;
        theOutcast = n;
      }
    }
    return theOutcast;
  }

  // see test client below
  public static void main(String[] args) {
    WordNet wordnet = new WordNet(args[0], args[1]);
    Outcast outcast = new Outcast(wordnet);
    for (int t = 2; t < args.length; t++) {
      In in = new In(args[t]);
      String[] nouns = in.readAllStrings();
      StdOut.println(args[t] + ": " + outcast.outcast(nouns));
    }
  }
}
