import edu.princeton.cs.algs4.WeightedQuickUnionUF;

public class Percolation {

  private boolean[][] grid;
  private int openSiteCount;

  private WeightedQuickUnionUF set;

  private int virtualTopIndex;
  private int virtualBottomIndex;
  private int totalCols;
  private int totalRows;

  // creates n-by-n grid, with all sites initially blocked
  public Percolation(int n) {
    if (n <= 0) {
      throw new IllegalArgumentException();
    }
    this.grid = new boolean[n][n];
    for (int i = 0; i < n; i++) {
      for (int j = 0; j < n; j++) {
        this.grid[i][j] = false;
      }
    }
    this.openSiteCount = 0;
    this.set = new WeightedQuickUnionUF(n * n + 2);
    this.virtualTopIndex = n * n;
    this.virtualBottomIndex = n * n + 1;
    this.totalCols = n;
    this.totalRows = n;
  }

  // opens the site (row, col) if it is not open already
  public void open(int row, int col) {
    if (row < 1 || col < 1 || row > totalRows || col > totalCols) {
      throw new IllegalArgumentException();
    }
    if (this.isOpen(row, col)) {
      return;
    }
    int zeroIndexedRow = row - 1;
    int zeroIndexedCol = col - 1;
    this.grid[zeroIndexedRow][zeroIndexedCol] = true;
    this.openSiteCount = this.openSiteCount + 1;
    int currentIndexInUnion = indexInUnion(zeroIndexedRow, zeroIndexedCol);
    if (zeroIndexedRow == 0) {
      this.set.union(currentIndexInUnion, virtualTopIndex);
    } else if (this.isOpen(row - 1, col)) {
      int top = indexInUnion(zeroIndexedRow - 1, zeroIndexedCol);
      this.set.union(currentIndexInUnion, top);
    }
    if (zeroIndexedRow == totalRows - 1) {
      this.set.union(currentIndexInUnion, virtualBottomIndex);
    } else if (this.isOpen(row + 1, col)) {
      int bottom = indexInUnion(zeroIndexedRow + 1, zeroIndexedCol);
      this.set.union(currentIndexInUnion, bottom);
    }
    if (zeroIndexedCol > 0 && this.isOpen(row, col - 1)) {
      int left = indexInUnion(zeroIndexedRow, zeroIndexedCol - 1);
      this.set.union(currentIndexInUnion, left);
    }
    if (zeroIndexedCol < this.totalCols - 1 && this.isOpen(row, col + 1)) {
      int right = indexInUnion(zeroIndexedRow, zeroIndexedCol + 1);
      this.set.union(currentIndexInUnion, right);
    }
  }

  // is the site (row, col) open?
  public boolean isOpen(int row, int col) {
    if (row < 1 || col < 1 || row > totalRows || col > totalCols) {
      throw new IllegalArgumentException();
    }
    int zeroIndexedRow = row - 1;
    int zeroIndexedCol = col - 1;
    return this.grid[zeroIndexedRow][zeroIndexedCol];
  }

  // is the site (row, col) full?
  public boolean isFull(int row, int col) {
    if (row < 1 || col < 1 || row > totalRows || col > totalCols) {
      throw new IllegalArgumentException();
    }
    int zeroIndexedRow = row - 1;
    int zeroIndexedCol = col - 1;
    int unionIndex = indexInUnion(zeroIndexedRow, zeroIndexedCol);
    return this.set.find(unionIndex) == this.set.find(virtualTopIndex);
  }

  // returns the number of open sites
  public int numberOfOpenSites() {
    return this.openSiteCount;
  }

  // does the system percolate?
  public boolean percolates() {
    return this.set.find(virtualTopIndex) == this.set.find(virtualBottomIndex);
  }

  // test client (optional)
  public static void main(String[] args) {

  }

  private int indexInUnion(int zeroIndexedRow, int zeroIndexedCol) {
    return zeroIndexedRow * this.totalCols + zeroIndexedCol;
  }
}
