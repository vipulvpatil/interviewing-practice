import edu.princeton.cs.algs4.WeightedQuickUnionUF;

public class Percolation {

  private static final byte CLOSED = 1;
  private static final byte OPEN = 2;
  private static final byte FULL = 4;
  private static final byte DRAINED = 8;

  private boolean[] grid;
  private byte[] status;
  private int openSiteCount;
  private boolean percolatedStatus;

  private WeightedQuickUnionUF set;

  private int totalCols;
  private int totalRows;

  // creates n-by-n grid, with all sites initially blocked
  public Percolation(int n) {
    if (n <= 0) {
      throw new IllegalArgumentException();
    }
    this.percolatedStatus = false;
    this.grid = new boolean[n * n];
    for (int i = 0; i < n * n; i++) {
      this.grid[i] = false;
    }
    this.status = new byte[n * n];
    for (int i = 0; i < n * n; i++) {
      this.status[i] = CLOSED;
    }
    this.openSiteCount = 0;
    this.set = new WeightedQuickUnionUF(n * n);
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
    int currentIndexInUnion = indexInUnion(zeroIndexedRow, zeroIndexedCol);
    byte newStatus = this.status[this.set.find(currentIndexInUnion)];
    newStatus |= OPEN;
    this.openSiteCount = this.openSiteCount + 1;
    this.grid[currentIndexInUnion] = false;
    if (zeroIndexedRow == 0) {
      newStatus |= FULL;
    } else if (this.isOpen(row - 1, col)) {
      int top = indexInUnion(zeroIndexedRow - 1, zeroIndexedCol);
      newStatus |= this.status[this.set.find(top)];
      this.set.union(currentIndexInUnion, top);
    }
    if (zeroIndexedRow == totalRows - 1) {
      newStatus |= DRAINED;
    } else if (this.isOpen(row + 1, col)) {
      int bottom = indexInUnion(zeroIndexedRow + 1, zeroIndexedCol);
      newStatus |= this.status[this.set.find(bottom)];
      this.set.union(currentIndexInUnion, bottom);
    }
    if (zeroIndexedCol > 0 && this.isOpen(row, col - 1)) {
      int left = indexInUnion(zeroIndexedRow, zeroIndexedCol - 1);
      newStatus |= this.status[this.set.find(left)];
      this.set.union(currentIndexInUnion, left);
    }
    if (zeroIndexedCol < this.totalCols - 1 && this.isOpen(row, col + 1)) {
      int right = indexInUnion(zeroIndexedRow, zeroIndexedCol + 1);
      newStatus |= this.status[this.set.find(right)];
      this.set.union(currentIndexInUnion, right);
    }
    if ((newStatus & FULL) != 0 && (newStatus & DRAINED) != 0) {
      this.percolatedStatus = true;
    }
    this.status[this.set.find(currentIndexInUnion)] = newStatus;
  }

  // is the site (row, col) open?
  public boolean isOpen(int row, int col) {
    if (row < 1 || col < 1 || row > totalRows || col > totalCols) {
      throw new IllegalArgumentException();
    }
    int zeroIndexedRow = row - 1;
    int zeroIndexedCol = col - 1;
    int unionIndex = indexInUnion(zeroIndexedRow, zeroIndexedCol);
    return (this.status[this.set.find(unionIndex)] & OPEN) != 0;
  }

  // is the site (row, col) full?
  public boolean isFull(int row, int col) {
    if (row < 1 || col < 1 || row > totalRows || col > totalCols) {
      throw new IllegalArgumentException();
    }
    int zeroIndexedRow = row - 1;
    int zeroIndexedCol = col - 1;
    int unionIndex = indexInUnion(zeroIndexedRow, zeroIndexedCol);
    return (this.status[this.set.find(unionIndex)] & FULL) != 0;
  }

  // returns the number of open sites
  public int numberOfOpenSites() {
    return this.openSiteCount;
  }

  // does the system percolate?
  public boolean percolates() {
    return this.percolatedStatus;
  }

  // test client (optional)
  public static void main(String[] args) {

  }

  private int indexInUnion(int zeroIndexedRow, int zeroIndexedCol) {
    return zeroIndexedRow * this.totalCols + zeroIndexedCol;
  }
}
