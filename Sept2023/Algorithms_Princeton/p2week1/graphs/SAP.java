import edu.princeton.cs.algs4.Bag;
import edu.princeton.cs.algs4.BreadthFirstDirectedPaths;
import edu.princeton.cs.algs4.Digraph;
import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;

public class SAP {
  private final Digraph graph;

  // constructor takes a digraph (not necessarily a DAG)
  public SAP(Digraph G) {
    graph = new Digraph(G);
  }

  // length of shortest ancestral path between v and w; -1 if no such path
  public int length(int v, int w) {
    Bag<Integer> vs = new Bag<Integer>();
    vs.add(v);
    Bag<Integer> ws = new Bag<Integer>();
    ws.add(w);
    return length(vs, ws);
  }

  // a common ancestor of v and w that participates in a shortest ancestral path;
  // -1 if no such path
  public int ancestor(int v, int w) {
    Bag<Integer> vs = new Bag<Integer>();
    vs.add(v);
    Bag<Integer> ws = new Bag<Integer>();
    ws.add(w);
    return ancestor(vs, ws);
  }

  // length of shortest ancestral path between any vertex in v and any vertex in
  // w; -1 if no such path
  public int length(Iterable<Integer> v, Iterable<Integer> w) {
    if (v == null || w == null) {
      throw new IllegalArgumentException();
    }
    if (!v.iterator().hasNext() || !w.iterator().hasNext()) {
      return -1;
    }
    SAPCalculator sapc = new SAPCalculator(graph, v, w);
    return sapc.shortestDistance();
  }

  // a common ancestor that participates in shortest ancestral path; -1 if no such
  // path
  public int ancestor(Iterable<Integer> v, Iterable<Integer> w) {
    if (v == null || w == null) {
      throw new IllegalArgumentException();
    }
    if (!v.iterator().hasNext() || !w.iterator().hasNext()) {
      return -1;
    }
    SAPCalculator sapc = new SAPCalculator(graph, v, w);
    return sapc.sca();
  }

  // do unit testing of this class
  public static void main(String[] args) {
    In in = new In(args[0]);
    Digraph G = new Digraph(in);
    SAP sap = new SAP(G);
    while (!StdIn.isEmpty()) {
      int v = StdIn.readInt();
      int w = StdIn.readInt();
      int length = sap.length(v, w);
      int ancestor = sap.ancestor(v, w);
      StdOut.printf("length = %d, ancestor = %d\n", length, ancestor);
    }
  }

  private class SAPCalculator {
    private int minDistance;
    private final int shortestCommonAncestor;

    public SAPCalculator(Digraph graph, Iterable<Integer> v, Iterable<Integer> w) {
      BreadthFirstDirectedPaths bfsPathsV = new BreadthFirstDirectedPaths(graph, v);
      BreadthFirstDirectedPaths bfsPathsW = new BreadthFirstDirectedPaths(graph, w);

      int minDist = (int) Double.POSITIVE_INFINITY;
      int selectedVertex = -1;
      for (int i = 0; i < graph.V(); i++) {
        if (bfsPathsV.hasPathTo(i) && bfsPathsW.hasPathTo(i)) {
          int distA = bfsPathsV.distTo(i);
          int distB = bfsPathsW.distTo(i);
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
}
