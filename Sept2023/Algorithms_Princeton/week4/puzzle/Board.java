import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdOut;

public class Board {

  private Iterable<Board> b;
  private int[][] currentTiles;

  // create a board from an n-by-n array of tiles,
  // where tiles[row][col] = tile at (row, col)
  public Board(int[][] tiles) {
    this.currentTiles = tiles;
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
    int dim = this.dimension();
    int outOfPlaceCount = 0;
    for (int row = 0; row < this.currentTiles.length; row++) {
      for (int col = 0; col < this.currentTiles[row].length; col++) {
        if (this.currentTiles[row][col] != row * dim + col) {
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
        if (this.currentTiles[row][col] != row * dim + col) {
          int targetRow = this.currentTiles[row][col] / dim;
          int targetCol = this.currentTiles[row][col] % dim;
          noOfMoves = noOfMoves + Math.abs(targetRow - row) + Math.abs(targetCol - col);
        }
      }
    }

    return noOfMoves;
  }

  // is this board the goal board?
  public boolean isGoal() {
    return false;
  }

  // does this board equal y?
  public boolean equals(Object y) {
    return false;
  }

  // all neighboring boards
  public Iterable<Board> neighbors() {
    return this.b;
  }

  // a board that is obtained by exchanging any pair of tiles
  public Board twin() {
    return this;
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
    StdOut.println(initial.manhattan());
  }

}
