import java.util.ArrayList;

import edu.princeton.cs.algs4.Bag;
import edu.princeton.cs.algs4.Digraph;
import edu.princeton.cs.algs4.DirectedCycle;
import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.ST;
import edu.princeton.cs.algs4.StdOut;

public class WordNet {

  private final ST<String, Bag<Integer>> wordNetST;
  private final ArrayList<String> words;
  private final SAP mySap;

  // constructor takes the name of the two input files
  public WordNet(String synsets, String hypernyms) {
    if (synsets == null || hypernyms == null) {
      throw new IllegalArgumentException();
    }
    words = new ArrayList<>();
    wordNetST = processSynsets(new In(synsets), words);
    int verticesCount = wordNetST.size();
    Digraph graph = new Digraph(verticesCount);
    addHypernymEdgesToGraph(new In(hypernyms), graph);

    boolean rootFound = false;
    for (int v = 0; v < graph.V(); v++) {
      if (graph.outdegree(v) == 0 && graph.indegree(v) > 0) {
        if (rootFound) {
          throw new IllegalArgumentException("too many roots");
        } else {
          rootFound = true;
        }
      }
    }

    if (!rootFound) {
      throw new IllegalArgumentException("no root found");
    }

    DirectedCycle dc = new DirectedCycle(graph);
    if (dc.hasCycle()) {
      throw new IllegalArgumentException();
    }

    mySap = new SAP(graph);
  }

  // returns all WordNet nouns
  public Iterable<String> nouns() {
    return wordNetST.keys();
  }

  // is the word a WordNet noun?
  public boolean isNoun(String word) {
    return wordNetST.contains(word);
  }

  // distance between nounA and nounB (defined below)
  public int distance(String nounA, String nounB) {
    if (!wordNetST.contains(nounA) || !wordNetST.contains(nounB)) {
      throw new IllegalArgumentException();
    }
    Bag<Integer> verticesA = wordNetST.get(nounA);
    Bag<Integer> verticesB = wordNetST.get(nounB);
    return mySap.length(verticesA, verticesB);
  }

  // a synset (second field of synsets.txt) that is the common ancestor of nounA
  // and nounB
  // in a shortest ancestral path (defined below)
  public String sap(String nounA, String nounB) {
    if (!wordNetST.contains(nounA) || !wordNetST.contains(nounB)) {
      throw new IllegalArgumentException();
    }
    Bag<Integer> verticesA = wordNetST.get(nounA);
    Bag<Integer> verticesB = wordNetST.get(nounB);
    int index = mySap.ancestor(verticesA, verticesB);
    return words.get(index);
  }

  // do unit testing of this class
  public static void main(String[] args) {
    WordNet wordNet = new WordNet(args[0], args[1]);
    String word1 = "taxonomy"; // "three"
    String word2 = "villainy"; // "three"
    StdOut.println(wordNet.distance(word1, word2));
    StdOut.println(wordNet.sap(word1, word2));
  }

  private ST<String, Bag<Integer>> processSynsets(In in, ArrayList<String> w) {
    ST<String, Bag<Integer>> st = new ST<String, Bag<Integer>>();
    while (!in.isEmpty()) {
      String line = in.readLine();
      String[] tokens = line.split(",");
      int val = Integer.parseInt(tokens[0]);
      w.add(tokens[1]);
      String[] keys = tokens[1].split(" ");
      for (String key : keys) {
        if (st.get(key) == null) {
          st.put(key, new Bag<>());
        }
        Bag<Integer> b = st.get(key);
        b.add(val);
        st.put(key, b);
      }
    }
    return st;
  }

  private void addHypernymEdgesToGraph(In in, Digraph g) {
    while (!in.isEmpty()) {
      String line = in.readLine();
      String[] tokens = line.split(",");
      int v = Integer.parseInt(tokens[0]);
      for (int i = 1; i < tokens.length; i++) {
        int w = Integer.parseInt(tokens[i]);
        g.addEdge(v, w);
      }
    }
  }
}
