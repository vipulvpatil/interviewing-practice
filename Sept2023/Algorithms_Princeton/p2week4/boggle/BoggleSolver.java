import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdOut;

public class BoggleSolver {
  private final Trie trie;

  // Initializes the data structure using the given array of strings as the
  // dictionary.
  // (You can assume each word in the dictionary contains only the uppercase
  // letters A through Z.)
  public BoggleSolver(String[] dictionary) {
    this.trie = new Trie(dictionary);
  }

  // Returns the set of all valid words in the given Boggle board, as an Iterable.
  public Iterable<String> getAllValidWords(BoggleBoard board) {
    BoggleBoardTrie bt = new BoggleBoardTrie(board);
    HashMap<String, Integer> collectedWords = new HashMap<>();
    for (int i = 0; i < board.rows(); i++) {
      for (int j = 0; j < board.cols(); j++) {
        BGTNode bNode = bt.nodes[i][j];
        bNode.visit();
        HashMap<String, Integer> newWords = this.trie.collectWords(this.trie.root, bt, bNode);
        collectedWords.putAll(newWords);
        bNode.unvisit();
      }
    }
    return collectedWords.keySet();
  }

  // Returns the score of the given word if it is in the dictionary, zero
  // otherwise.
  // (You can assume the word contains only the uppercase letters A through Z.)
  public int scoreOf(String word) {
    return this.trie.findWordScore(word);
  }

  public static void main(String[] args) {
    In in = new In(args[0]);
    String[] dictionary = in.readAllStrings();
    BoggleSolver solver = new BoggleSolver(dictionary);
    BoggleBoard board = new BoggleBoard(args[1]);
    int score = 0;
    for (String word : solver.getAllValidWords(board)) {
      StdOut.println(word);
      score += solver.scoreOf(word);
    }
    StdOut.println("Score = " + score);
  }

  private static int intForChar(char c) {
    return c - 65;
  }

  private static int scoreForCount(int count) {
    int score = 0;
    if (count >= 8) {
      score = 11;
    } else if (count >= 7) {
      score = 5;
    } else if (count >= 6) {
      score = 3;
    } else if (count >= 5) {
      score = 2;
    } else if (count >= 3) {
      score = 1;
    }
    return score;
  }

  private class Node {
    private Node[] children;
    private String word;
    private int score;

    public Node() {
      children = new Node[26];
    }
  }

  private class Trie {
    private Node root;

    public Trie(String[] dictionary) {
      for (String s : dictionary) {
        root = add(root, s.toUpperCase(), "", 0);
      }
    }

    public int findWordScore(String s) {
      Node word = findWord(root, s.toUpperCase());
      if (word != null) {
        return word.score;
      }
      return 0;
    }

    private Node add(Node n, String s, String path, int count) {
      if (n == null) {
        n = new Node();
      }
      if (s.length() == 0) {
        n.word = path;
        n.score = scoreForCount(count);
        return n;
      }
      char c = s.charAt(0);
      int i = intForChar(c);
      n.children[i] = add(n.children[i], s.substring(1), path + c, count + 1);
      return n;
    }

    private Node findWord(Node n, String s) {
      if (s.length() == 0) {
        return n;
      }
      char c = s.charAt(0);
      int i = intForChar(c);
      Node nextNode = n.children[i];
      if (nextNode == null) {
        return null;
      }
      return findWord(nextNode, s.substring(1));
    }

    private HashMap<String, Integer> collectWords(Node n, BoggleBoardTrie t, BGTNode b) {
      int i = intForChar(b.character);
      Node nextNode = n.children[i];
      if (b.character == "Q".charAt(0)) {
        if (nextNode != null) {
          nextNode = nextNode.children[intForChar("U".charAt(0))];
        }
      }
      HashMap<String, Integer> furtherWords = new HashMap<String, Integer>();
      if (nextNode != null) {
        BGTNode[] nextBGTNodes = t.unvisitedNeighbors(b);
        for (int x = 0; x < nextBGTNodes.length; x++) {
          BGTNode nextBGTNode = nextBGTNodes[x];
          nextBGTNode.visit();
          furtherWords.putAll(collectWords(nextNode, t, nextBGTNode));
          nextBGTNode.unvisit();
        }
        if (nextNode.score > 0) {
          furtherWords.put(nextNode.word, nextNode.score);
        }
      }
      return furtherWords;
    }
  }

  private class BGTNode {
    private final char character;
    private boolean visited;
    private final int i;
    private final int j;

    public BGTNode(int i, int j, char character) {
      this.i = i;
      this.j = j;
      this.visited = false;
      this.character = character;
    }

    public void visit() {
      visited = true;
    }

    public void unvisit() {
      visited = false;
    }
  }

  private class BoggleBoardTrie {
    public BGTNode[][] nodes;
    private final int rows;
    private final int cols;

    public BoggleBoardTrie(BoggleBoard board) {
      rows = board.rows();
      cols = board.cols();
      nodes = new BGTNode[rows][cols];
      for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
          nodes[i][j] = new BGTNode(i, j, board.getLetter(i, j));
        }
      }
    }

    public BGTNode[] unvisitedNeighbors(BGTNode b) {
      List<BGTNode> list = new ArrayList<>();
      for (int i = b.i - 1; i <= b.i + 1; i++) {
        for (int j = b.j - 1; j <= b.j + 1; j++) {
          if (i >= 0 && j >= 0 && i < rows && j < cols) {
            BGTNode neighbor = nodes[i][j];
            if (!neighbor.visited) {
              list.add(neighbor);
            }
          }
        }
      }
      return list.toArray(new BGTNode[0]);
    }
  }

}
