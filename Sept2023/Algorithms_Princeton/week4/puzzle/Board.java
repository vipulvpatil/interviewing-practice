import java.util.Arrays;
import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.Queue;
import edu.princeton.cs.algs4.StdOut;

public class Board {
  private int[][] currentTiles;

  // create a board from an n-by-n array of tiles,
  // where tiles[row][col] = tile at (row, col)
  public Board(int[][] tiles) {
    this.currentTiles = new int[tiles.length][tiles.length];
    for (int row = 0; row < this.currentTiles.length; row++) {
      for (int col = 0; col < this.currentTiles[row].length; col++) {
        this.currentTiles[row][col] = tiles[row][col];
      }
    }
  }

  // string representation of this board
  public String toString() {
    String outputString = String.format("%d\n", this.dimension());
    for (int row = 0; row < this.currentTiles.length; row++) {
      for (int col = 0; col < this.currentTiles[row].length; col++) {
        outputString = outputString.concat(String.format(" %d", this.currentTiles[row][col]));
      }
      outputString = outputString.concat("\n");
    }

    return outputString;
  }

  // board dimension n
  public int dimension() {
    return this.currentTiles.length;
  }

  // number of tiles out of place
  public int hamming() {
    int outOfPlaceCount = 0;
    for (int row = 0; row < this.currentTiles.length; row++) {
      for (int col = 0; col < this.currentTiles[row].length; col++) {
        if (this.currentTiles[row][col] != 0 && this.currentTiles[row][col] != goal(row, col)) {
          outOfPlaceCount++;
        }
      }
    }

    return outOfPlaceCount;
  }

  // sum of Manhattan distances between tiles and goal
  public int manhattan() {
    int dim = this.dimension();
    int noOfMoves = 0;
    for (int row = 0; row < this.currentTiles.length; row++) {
      for (int col = 0; col < this.currentTiles[row].length; col++) {
        int g = goal(row, col);
        if (this.currentTiles[row][col] != g) {
          if (this.currentTiles[row][col] != 0) {
            int targetRow = (this.currentTiles[row][col] - 1) / dim;
            int targetCol = (this.currentTiles[row][col] - 1) % dim;
            noOfMoves = noOfMoves + Math.abs(targetRow - row) + Math.abs(targetCol - col);
          }
        }
      }
    }

    return noOfMoves;
  }

  // is this board the goal board?
  public boolean isGoal() {
    for (int row = 0; row < this.currentTiles.length; row++) {
      for (int col = 0; col < this.currentTiles[row].length; col++) {
        if (this.currentTiles[row][col] != goal(row, col)) {
          return false;
        }
      }
    }

    return true;
  }

  // does this board equal y?
  public boolean equals(Object y) {
    if (y == null) {
      return false;
    }
    if (y.getClass() != this.getClass()) {
      return false;
    }

    Board board = (Board) y;
    if (this.dimension() != board.dimension()) {
      return false;
    }
    return Arrays.deepEquals(this.currentTiles, board.currentTiles);
  }

  // all neighboring boards
  public Iterable<Board> neighbors() {
    Queue<Board> queue = new Queue<Board>();

    // BoardIterable iter = new BoardIterable();
    int dim = this.dimension();
    int row0 = 0;
    int col0 = 0;
    for (int row = 0; row < this.currentTiles.length; row++) {
      for (int col = 0; col < this.currentTiles[row].length; col++) {
        if (this.currentTiles[row][col] == 0) {
          row0 = row;
          col0 = col;
        }
      }
    }
    if (col0 > 0) {
      int[][] moveTiles = new int[dim][dim];
      int movingCell = this.currentTiles[row0][col0 - 1];
      for (int row = 0; row < this.currentTiles.length; row++) {
        for (int col = 0; col < this.currentTiles[row].length; col++) {
          moveTiles[row][col] = this.currentTiles[row][col];
        }
      }
      moveTiles[row0][col0] = movingCell;
      moveTiles[row0][col0 - 1] = 0;

      queue.enqueue(new Board(moveTiles));
    }
    if (col0 < dim - 1) {
      int[][] moveTiles = new int[dim][dim];
      int movingCell = this.currentTiles[row0][col0 + 1];
      for (int row = 0; row < this.currentTiles.length; row++) {
        for (int col = 0; col < this.currentTiles[row].length; col++) {
          moveTiles[row][col] = this.currentTiles[row][col];
        }
      }
      moveTiles[row0][col0] = movingCell;
      moveTiles[row0][col0 + 1] = 0;

      queue.enqueue(new Board(moveTiles));
    }
    if (row0 > 0) {
      int[][] moveTiles = new int[dim][dim];
      int movingCell = this.currentTiles[row0 - 1][col0];
      for (int row = 0; row < this.currentTiles.length; row++) {
        for (int col = 0; col < this.currentTiles[row].length; col++) {
          moveTiles[row][col] = this.currentTiles[row][col];
        }
      }
      moveTiles[row0][col0] = movingCell;
      moveTiles[row0 - 1][col0] = 0;

      queue.enqueue(new Board(moveTiles));
    }
    if (row0 < dim - 1) {
      int[][] moveTiles = new int[dim][dim];
      int movingCell = this.currentTiles[row0 + 1][col0];
      for (int row = 0; row < this.currentTiles.length; row++) {
        for (int col = 0; col < this.currentTiles[row].length; col++) {
          moveTiles[row][col] = this.currentTiles[row][col];
        }
      }
      moveTiles[row0][col0] = movingCell;
      moveTiles[row0 + 1][col0] = 0;

      queue.enqueue(new Board(moveTiles));
    }
    return queue;
  }

  // a board that is obtained by exchanging any pair of tiles
  public Board twin() {
    int dim = this.dimension();

    int[][] newTiles = new int[dim][dim];
    for (int row = 0; row < this.currentTiles.length; row++) {
      for (int col = 0; col < this.currentTiles[row].length; col++) {
        newTiles[row][col] = this.currentTiles[row][col];
      }
    }

    int targetRow1 = 0 / dim;
    int targetCol1 = 0 % dim;
    int targetRow2 = (dim * dim - 1) / dim;
    int targetCol2 = (dim * dim - 1) % dim;

    if (newTiles[targetRow1][targetCol1] == 0) {
      targetRow1++;
    }
    if (newTiles[targetRow2][targetCol2] == 0) {
      targetRow2--;
    }

    int temp = newTiles[targetRow1][targetCol1];
    newTiles[targetRow1][targetCol1] = newTiles[targetRow2][targetCol2];
    newTiles[targetRow2][targetCol2] = temp;

    return new Board(newTiles);
  }

  private int goal(int row, int col) {
    int dim = dimension();
    return (row * dim + col + 1) % (dim * dim);
  }

  // unit testing (not graded)
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
    // Solver solver = new Solver(initial);

    // print solution to standard output
    // if (!solver.isSolvable())
    // StdOut.println("No solution possible");
    // else {
    // StdOut.println("Minimum number of moves = " + solver.moves());
    // for (Board board : solver.solution())
    // StdOut.println(board);
    // }

    StdOut.println(initial);

    StdOut.println(initial.hamming());
    StdOut.println("manhattan");
    StdOut.println(initial.manhattan());
    tiles[1][1] = 5;
    tiles[2][1] = 2;
    StdOut.println("manhattan");
    StdOut.println(initial.manhattan());

    StdOut.println("neighbors");
    for (Board board : initial.neighbors())
      StdOut.println(board);

    StdOut.println("twins");
    StdOut.println(initial.twin());

    StdOut.println("other");
  }
}
