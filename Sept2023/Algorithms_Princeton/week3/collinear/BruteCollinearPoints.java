import java.util.Arrays;

import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdDraw;
import edu.princeton.cs.algs4.StdOut;

public class BruteCollinearPoints {
  private LineSegment[] lineSegments;

  public BruteCollinearPoints(Point[] points) {
    if (points == null) {
      throw new IllegalArgumentException();
    }

    for (int i = 0; i < points.length; i++) {
      if (points[i] == null) {
        throw new IllegalArgumentException();
      }
    }

    this.lineSegments = new LineSegment[0];

    for (int i = 0; i < points.length; i++) {
      for (int j = i + 1; j < points.length; j++) {
        for (int k = j + 1; k < points.length; k++) {
          for (int l = k + 1; l < points.length; l++) {
            Point pI = points[i];
            Point pJ = points[j];
            Point pK = points[k];
            Point pL = points[l];
            if (pI.slopeTo(pJ) == pI.slopeTo(pK) &&
                pI.slopeTo(pJ) == pI.slopeTo(pL)) {
              int n = this.numberOfSegments();
              LineSegment[] newLineSegments = new LineSegment[n + 1];
              for (int x = 0; x < n; x++) {
                newLineSegments[x] = this.lineSegments[x];
              }

              Point[] collinearPoints = { pI, pJ, pK, pL };
              Arrays.sort(collinearPoints);
              newLineSegments[n] = new LineSegment(collinearPoints[0], collinearPoints[3]);
              this.lineSegments = newLineSegments;
            }
          }
        }
      }
    }
  }

  public int numberOfSegments() {
    return this.lineSegments.length;
  }

  public LineSegment[] segments() {
    return this.lineSegments;
  }

  public static void main(String[] args) {
    // read the n points from a file
    In in = new In(args[0]);
    int n = in.readInt();
    Point[] points = new Point[n];
    for (int i = 0; i < n; i++) {
      int x = in.readInt();
      int y = in.readInt();
      points[i] = new Point(x, y);
    }

    // draw the points
    StdDraw.enableDoubleBuffering();
    StdDraw.setXscale(0, 32768);
    StdDraw.setYscale(0, 32768);
    for (Point p : points) {
      p.draw();
    }
    StdDraw.show();

    // print and draw the line segments
    BruteCollinearPoints collinear = new BruteCollinearPoints(points);
    for (LineSegment segment : collinear.segments()) {
      segment.draw();
    }
    StdDraw.show();
  }
}
