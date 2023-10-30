import java.util.HashMap;

import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdOut;

public class BoggleSolver {
  Trie trie;

  // Initializes the data structure using the given array of strings as the
  // dictionary.
  // (You can assume each word in the dictionary contains only the uppercase
  // letters A through Z.)
  public BoggleSolver(String[] dictionary) {
    this.trie = new Trie(dictionary);
    StdOut.println(this.trie.findWords("friend"));
    StdOut.println(this.trie.findWords("friendly"));
    StdOut.println(this.trie.findWords("fry"));
    StdOut.println(this.trie.findWords("fraud"));
    StdOut.println(this.trie.findWords("defraud"));
  }

  // Returns the set of all valid words in the given Boggle board, as an Iterable.
  public Iterable<String> getAllValidWords(BoggleBoard board) {

    return null;
  }

  // Returns the score of the given word if it is in the dictionary, zero
  // otherwise.
  // (You can assume the word contains only the uppercase letters A through Z.)
  public int scoreOf(String word) {
    return 0;
  }

  public static void main(String[] args) {
    // In in = new In(args[0]);
    // String[] dictionary = in.readAllStrings();
    String[] dictionary = new String[] {
        "friend", "friendly", "fry", "deep", "friendship",
    };
    BoggleSolver solver = new BoggleSolver(dictionary);
    // BoggleBoard board = new BoggleBoard(args[1]);
    int score = 0;
    // for (String word : solver.getAllValidWords(board)) {
    // StdOut.println(word);
    // score += solver.scoreOf(word);
    // }
    StdOut.println("Score = " + score);
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

    private HashMap<String, Integer> findWords(String s) {
      return findWords(root, s.toUpperCase());
    }

    private Node add(Node n, String s, String path, int count) {
      if (n == null) {
        n = new Node();
      }
      if (s.length() == 0) {
        n.word = path;
        n.score = scoreForCount(count + 1);
        return n;
      }
      char c = s.charAt(0);
      int i = intForChar(c);
      n.children[i] = add(n.children[i], s.substring(1), path + c, count + 1);
      return n;
    }

    private HashMap<String, Integer> findWords(Node n, String s) {
      HashMap<String, Integer> newWordFound = new HashMap<>();
      if (n.score > 0) {
        newWordFound.put(n.word, n.score);
      }
      if (s.length() == 0) {
        return newWordFound;
      }
      char c = s.charAt(0);
      int i = intForChar(c);
      Node nextNode = n.children[i];
      if (nextNode == null) {
        return newWordFound;
      }
      HashMap<String, Integer> furtherWords = findWords(nextNode, s.substring(1));
      if (n.score > 0) {
        furtherWords.put(n.word, n.score);
      }
      return furtherWords;
    }

    private final int intForChar(char c) {
      return c - 65;
    }

    private final int scoreForCount(int count) {
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
  }

}
