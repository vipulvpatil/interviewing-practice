import java.util.ArrayList;
import edu.princeton.cs.algs4.Point2D;
import edu.princeton.cs.algs4.RectHV;
import edu.princeton.cs.algs4.SET;

public class PointSET {
  private SET<Point2D> points;

  public PointSET() {
    points = new SET<Point2D>();
  }

  public boolean isEmpty() {
    return points.isEmpty();
  }

  public int size() {
    return points.size();
  }

  public void insert(Point2D p) {
    if (p == null) {
      throw new IllegalArgumentException();
    }
    points.add(p);
  }

  public boolean contains(Point2D p) {
    if (p == null) {
      throw new IllegalArgumentException();
    }
    return points.contains(p);
  }

  public void draw() {
    for (Point2D p : points) {
      p.draw();
    }
  }

  public Iterable<Point2D> range(RectHV rect) {
    if (rect == null) {
      throw new IllegalArgumentException();
    }
    ArrayList<Point2D> contained = new ArrayList<Point2D>();
    for (Point2D p : points) {
      if (rect.contains(p)) {
        contained.add(p);
      }
    }
    return contained;
  }

  public Point2D nearest(Point2D p) {
    if (p == null) {
      throw new IllegalArgumentException();
    }
    if (points.isEmpty()) {
      return null;
    }
    Point2D closest = points.min();
    double minDistSq = Double.MAX_VALUE;
    for (Point2D point : points) {
      double distSq = p.distanceSquaredTo(point);
      if (distSq < minDistSq) {
        closest = point;
        minDistSq = distSq;
      }
    }

    return closest;
  }

  public static void main(String[] args) {

  }
}
