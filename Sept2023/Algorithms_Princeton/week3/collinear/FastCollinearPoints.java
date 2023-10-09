import java.util.Arrays;

import edu.princeton.cs.algs4.In;
import edu.princeton.cs.algs4.StdDraw;

public class FastCollinearPoints {
  private LineSegment[] lineSegments;

  public FastCollinearPoints(Point[] points) {
    if (points == null) {
      throw new IllegalArgumentException();
    }

    for (int i = 0; i < points.length; i++) {
      if (points[i] == null) {
        throw new IllegalArgumentException();
      }
    }

    this.lineSegments = new LineSegment[0];

    Arrays.sort(points);

    for (int i = 0; i < points.length; i++) {
      Point p0 = points[i];
      Point[] otherPoints = new Point[points.length - i - 1];
      for (int j = 0; j < points.length - i - 1; j++) {
        otherPoints[j] = points[j + i + 1];
      }

      Arrays.sort(otherPoints, p0.slopeOrder());

      Point[] collinearPoints = new Point[0];
      for (int j = 0; j < otherPoints.length - 1; j++) {
        if (p0.slopeTo(otherPoints[j]) == p0.slopeTo(otherPoints[j + 1])) {
          if (collinearPoints.length == 0) {
            collinearPoints = append(collinearPoints, otherPoints[j]);
          }
          collinearPoints = append(collinearPoints, otherPoints[j + 1]);
        } else {
          if (collinearPoints.length >= 3) {
            collinearPoints = append(collinearPoints, p0);
            Arrays.sort(collinearPoints);
            LineSegment lineSegment = new LineSegment(collinearPoints[0], collinearPoints[collinearPoints.length - 1]);
            this.lineSegments = append(this.lineSegments, lineSegment);
          }
          collinearPoints = new Point[0];
        }
      }
      if (collinearPoints.length >= 3) {
        collinearPoints = append(collinearPoints, p0);
        Arrays.sort(collinearPoints);
        LineSegment lineSegment = new LineSegment(collinearPoints[0], collinearPoints[collinearPoints.length - 1]);
        this.lineSegments = append(this.lineSegments, lineSegment);
      }
    }
  }

  private LineSegment[] append(LineSegment[] lineSegments, LineSegment lineSegment) {
    int n = lineSegments.length;
    LineSegment[] newLineSegments = new LineSegment[n + 1];
    for (int x = 0; x < n; x++) {
      newLineSegments[x] = lineSegments[x];
    }
    newLineSegments[n] = lineSegment;

    return newLineSegments;
  }

  private Point[] append(Point[] points, Point point) {
    int n = points.length;
    Point[] newPoints = new Point[n + 1];
    for (int x = 0; x < n; x++) {
      newPoints[x] = points[x];
    }
    newPoints[n] = point;
    return newPoints;
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
    FastCollinearPoints collinear = new FastCollinearPoints(points);
    for (LineSegment segment : collinear.segments()) {
      segment.draw();
    }
    StdDraw.show();
  }
}
