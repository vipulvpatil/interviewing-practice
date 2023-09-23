import edu.princeton.cs.algs4.StdIn;
import edu.princeton.cs.algs4.StdOut;
import edu.princeton.cs.algs4.StdRandom;
import edu.princeton.cs.algs4.StdStats;

public class PercolationStats {

  private static final double CONFIDENCE_95 = 1.96;
  private double[] results;
  private int trialCount;

  // perform independent trials on an n-by-n grid
  public PercolationStats(int n, int trials) {
    if (n <= 0 || trials <= 0) {
      throw new IllegalArgumentException();
    }
    trialCount = trials;
    results = new double[trials];
    for (int t = 0; t < trials; t++) {
      Percolation perc = new Percolation(n);
      int[] randomSiteList = StdRandom.permutation(n * n);
      int index = 0;
      while (!perc.percolates()) {
        int siteIndex = randomSiteList[index];
        int siteRow = siteIndex / n;
        int siteCol = siteIndex % n;
        perc.open(siteRow + 1, siteCol + 1);
        index++;
      }
      results[t] = ((double) perc.numberOfOpenSites()) / (n * n);
    }
  }

  // sample mean of percolation threshold
  public double mean() {
    return StdStats.mean(results);
  }

  // sample standard deviation of percolation threshold
  public double stddev() {
    return StdStats.stddev(results);
  }

  // low endpoint of 95% confidence interval
  public double confidenceLo() {
    return this.mean() - (CONFIDENCE_95 * this.stddev() / Math.sqrt(trialCount));
  }

  // high endpoint of 95% confidence interval
  public double confidenceHi() {
    return this.mean() + (CONFIDENCE_95 * this.stddev() / Math.sqrt(trialCount));
  }

  // test client (see below)
  public static void main(String[] args) {
    int n = Integer.parseInt(StdIn.readString());
    int trials = Integer.parseInt(StdIn.readString());
    PercolationStats percStats = new PercolationStats(n, trials);
    StdOut.printf("mean                    = %f\n", percStats.mean());
    StdOut.printf("stddev                  = %f\n", percStats.stddev());
    StdOut.printf("95%% confidence interval = [%f, %f]\n", percStats.confidenceLo(), percStats.confidenceHi());
  }
}
