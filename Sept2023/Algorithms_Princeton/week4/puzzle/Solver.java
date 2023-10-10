import java.util.Comparator;
import java.util.Iterator;
import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.MinPQ;
import edu.princeton.cs.algs4.StdOut;

public class Solver {
  private class GameTreeNode {
    private Board board;
    private GameTreeNode previous;

    public GameTreeNode(Board b) {
      board = b;
    }
  }

  private GameTreeNode winningGameTreeNode;

  private class HammingComparator implements Comparator<GameTreeNode> {
    public int compare(GameTreeNode a, GameTreeNode b) {
      if (a.board.hamming() < b.board.hamming()) {
        return -1;
      }
      if (a.board.hamming() > b.board.hamming()) {
        return 1;
      }
      return 0;
    }
  }

  private class ManhattanComparator implements Comparator<GameTreeNode> {
    public int compare(GameTreeNode a, GameTreeNode b) {
      if (a.board.manhattan() < b.board.manhattan()) {
        return -1;
      }
      if (a.board.manhattan() > b.board.manhattan()) {
        return 1;
      }
      return 0;
    }
  }

  // find a solution to the initial board (using the A* algorithm)
  public Solver(Board initial) {
    if (initial == null) {
      throw new IllegalArgumentException();
    }
    OneStepSolver mainSolution = new OneStepSolver(initial);
    Board alt = initial.twin();
    OneStepSolver altSolution = new OneStepSolver(alt);
    while (true) {
      winningGameTreeNode = mainSolution.NextStep();
      GameTreeNode altGameTreeNode = altSolution.NextStep();
      if (isSolvable()) {
        break;
      }
      if (altGameTreeNode != null) {
        break;
      }
    }
  }

  // is the initial board solvable? (see below)
  public boolean isSolvable() {
    return winningGameTreeNode != null;
  }

  // min number of moves to solve initial board; -1 if unsolvable
  public int moves() {
    if (!isSolvable()) {
      return -1;
    }
    int movesCount = 0;
    GameTreeNode gameTreeNode = winningGameTreeNode;
    while (gameTreeNode != null) {
      movesCount++;
      gameTreeNode = gameTreeNode.previous;
    }
    return movesCount - 1;
  }

  // sequence of boards in a shortest solution; null if unsolvable
  public Iterable<Board> solution() {
    if (!isSolvable()) {
      return null;
    }
    GameTreeNode gameTreeNode = winningGameTreeNode;
    BoardIterable reverseIterator = new BoardIterable();
    while (gameTreeNode != null) {
      reverseIterator.add(gameTreeNode.board);
      gameTreeNode = gameTreeNode.previous;
    }
    return reverseIterator;
  }

  private class OneStepSolver {
    private MinPQ<GameTreeNode> queue;
    private GameTreeNode winningGameTreeNode;

    public OneStepSolver(Board initial) {
      GameTreeNode gameTreeRoot = new GameTreeNode(initial);
      queue = new MinPQ<GameTreeNode>(0, new ManhattanComparator());
      queue.insert(gameTreeRoot);
    }

    public GameTreeNode NextStep() {
      if (!queue.isEmpty()) {
        GameTreeNode gameTreeNode = queue.delMin();
        // solutionForNode(gameTreeNode);
        if (gameTreeNode.board.isGoal()) {
          winningGameTreeNode = gameTreeNode;
        }
        for (Board board : gameTreeNode.board.neighbors()) {
          if (gameTreeNode.previous == null || !board.equals(gameTreeNode.previous.board)) {
            GameTreeNode newGameTreeNode = new GameTreeNode(board);
            newGameTreeNode.previous = gameTreeNode;
            queue.insert(newGameTreeNode);
          }
        }
      }

      return winningGameTreeNode;
    }

    // public void solutionForNode(GameTreeNode gameTreeNode) {
    // StdOut.println("solutionForNode");
    // // while (gameTreeNode != null) {
    // StdOut.println(gameTreeNode.board);
    // // gameTreeNode = gameTreeNode.previous;
    // // }
    // }
  }

  private class BoardIterable implements Iterable<Board> {
    BoardStack iter;

    public BoardIterable() {
      iter = new BoardStack();
    }

    public void add(Board b) {
      iter.add(b);
    }

    public Iterator<Board> iterator() {
      return iter;
    }

    private class BoardStack implements Iterator<Board> {
      private class BoardNode {
        public Board board;
        public BoardNode previous;

        public BoardNode(Board b) {
          this.board = b;
        }
      }

      BoardNode current;

      public void add(Board b) {
        BoardNode node = new BoardNode(b);
        if (current == null) {
          current = node;
        } else {
          node.previous = current;
          current = node;
        }
      }

      public boolean hasNext() {
        return current != null;
      }

      public Board next() {
        if (!hasNext()) {
          throw new java.util.NoSuchElementException();
        }
        Board board = current.board;
        current = current.previous;
        return board;
      }
    }
  }

  // test client (see below)
  public static void main(String[] args) {

    // create initial board from file
    In in = new In(args[0]);
    int n = in.readInt();
    int[][] tiles = new int[n][n];
    for (int i = 0; i < n; i++)
      for (int j = 0; j < n; j++)
        tiles[i][j] = in.readInt();
    Board initial = new Board(tiles);

    // solve the puzzle
    Solver solver = new Solver(initial);

    // print solution to standard output
    if (!solver.isSolvable())
      StdOut.println("No solution possible");
    else {
      StdOut.println("Minimum number of moves = " + solver.moves());
      for (Board board : solver.solution())
        StdOut.println(board);
    }
  }

}
