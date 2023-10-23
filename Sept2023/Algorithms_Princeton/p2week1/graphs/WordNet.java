import java.util.ArrayList;

import edu.princeton.cs.algs4.Bag;
import edu.princeton.cs.algs4.BreadthFirstDirectedPaths;
import edu.princeton.cs.algs4.Digraph;
import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.ST;
import edu.princeton.cs.algs4.StdOut;

public class WordNet {

  private ST<String, Bag<Integer>> wordNetST;
  private ArrayList<String> words;
  private Digraph graph;

  // constructor takes the name of the two input files
  public WordNet(String synsets, String hypernyms) {
    if (synsets == null || hypernyms == null) {
      throw new IllegalArgumentException();
    }
    words = new ArrayList<>();
    wordNetST = processSynsets(new In(synsets), words);
    int verticesCount = wordNetST.size();
    graph = new Digraph(verticesCount);
    addHypernymEdgesToGraph(new In(hypernyms), graph);
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
    ShortestCommonAncestor s = new ShortestCommonAncestor(nounA, nounB);
    return s.shortestDistance();
  }

  // a synset (second field of synsets.txt) that is the common ancestor of nounA
  // and nounB
  // in a shortest ancestral path (defined below)
  public String sap(String nounA, String nounB) {
    ShortestCommonAncestor s = new ShortestCommonAncestor(nounA, nounB);
    return words.get(s.sca());
  }

  private class ShortestCommonAncestor {
    private int minDistance;
    private int shortestCommonAncestor;

    ShortestCommonAncestor(String nounA, String nounB) {
      if (!wordNetST.contains(nounA) || !wordNetST.contains(nounB)) {
        throw new IllegalArgumentException();
      }
      Bag<Integer> vertexA = wordNetST.get(nounA);
      Bag<Integer> vertexB = wordNetST.get(nounB);
      BreadthFirstDirectedPaths bfsA = new BreadthFirstDirectedPaths(graph, vertexA);
      BreadthFirstDirectedPaths bfsB = new BreadthFirstDirectedPaths(graph, vertexB);

      int minDist = (int) Double.POSITIVE_INFINITY;
      int selectedVertex = -1;
      for (int i = 0; i < graph.V(); i++) {
        if (bfsA.hasPathTo(i) && bfsB.hasPathTo(i)) {
          int distA = bfsA.distTo(i);
          int distB = bfsB.distTo(i);
          if (distA + distB < minDist) {
            selectedVertex = i;
            minDist = distA + distB;
          }
        }
      }
      if (selectedVertex == -1) {
        minDistance = -1;
      } else {
        minDistance = minDist;
      }
      shortestCommonAncestor = selectedVertex;
    }

    public int shortestDistance() {
      return minDistance;
    }

    public int sca() {
      return shortestCommonAncestor;
    }
  }

  // do unit testing of this class
  public static void main(String[] args) {
    WordNet wordNet = new WordNet(args[0], args[1]);
    StdOut.println(wordNet.distance("three", "eight2"));
    StdOut.println(wordNet.sap("three", "eight2"));
  }

  private ST<String, Bag<Integer>> processSynsets(In in, ArrayList<String> words) {
    ST<String, Bag<Integer>> st = new ST<String, Bag<Integer>>();
    while (!in.isEmpty()) {
      String line = in.readLine();
      String[] tokens = line.split(",");
      Integer val = Integer.parseInt(tokens[0]);
      words.add(tokens[1]);
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
      Integer v = Integer.parseInt(tokens[0]);
      for (int i = 1; i < tokens.length; i++) {
        Integer w = Integer.parseInt(tokens[i]);
        g.addEdge(v, w);
      }
    }
  }
}
