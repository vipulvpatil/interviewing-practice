import java.util.ArrayList;

import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;
import edu.princeton.cs.algs4.StdRandom;

public class PercolationStats {

  private double[] results;
  private int trialCount;

  class Site {
    public int row;
    public int col;

    public Site(int row, int col) {
      this.row = row + 1;
      this.col = col + 1;
    }
  }

  // perform independent trials on an n-by-n grid
  public PercolationStats(int n, int trials) {
    if (n <= 0 || trials <= 0) {
      throw new IllegalArgumentException();
    }
    trialCount = trials;
    results = new double[trials];
    for (int t = 0; t < trials; t++) {
      Percolation perc = new Percolation(n);
      ArrayList<Site> fullsites = new ArrayList<Site>(n * n);
      for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
          fullsites.add(new Site(i, j));
        }
      }
      while (!perc.percolates()) {
        int randomlySelectedIndex = StdRandom.uniformInt(0, fullsites.size());
        Site randomlySelected = fullsites.get(randomlySelectedIndex);
        fullsites.remove(randomlySelectedIndex);
        perc.open(randomlySelected.row, randomlySelected.col);
      }
      results[t] = ((double) perc.numberOfOpenSites()) / (n * n);
    }
  }

  // sample mean of percolation threshold
  public double mean() {
    double sum = 0;
    for (int t = 0; t < trialCount; t++) {
      sum = sum + results[t];
    }
    return sum / trialCount;
  }

  // sample standard deviation of percolation threshold
  public double stddev() {
    double calculatedMean = this.mean();
    double sum = 0;
    for (int t = 0; t < trialCount; t++) {
      sum = sum + Math.pow((results[t] - calculatedMean), 2);
    }
    return Math.sqrt(sum / (trialCount - 1));
  }

  // low endpoint of 95% confidence interval
  public double confidenceLo() {
    return this.mean() - (1.96 * this.stddev() / Math.sqrt(trialCount));
  }

  // high endpoint of 95% confidence interval
  public double confidenceHi() {
    return this.mean() + (1.96 * this.stddev() / Math.sqrt(trialCount));
  }

  // test client (see below)
  public static void main(String[] args) {
    PercolationStats percStats = new PercolationStats(StdIn.readInt(), StdIn.readInt());
    StdOut.printf("mean                    = %f\n", percStats.mean());
    StdOut.printf("stddev                  = %f\n", percStats.stddev());
    StdOut.printf("95%% confidence interval = [%f, %f]\n", percStats.confidenceLo(), percStats.confidenceHi());
  }
}
