import java.util.ArrayList;
import java.util.Set;

import edu.princeton.cs.algs4.Point2D;
import edu.princeton.cs.algs4.RectHV;
import edu.princeton.cs.algs4.StdDraw;

public class KdTree {
  private class Node {
    public Point2D point;
    public Node left;
    public Node right;

    public Node(Point2D p) {
      if (p == null) {
        throw new IllegalArgumentException();
      }
      point = p;
    }

    public int size() {
      if (this == null) {
        return 0;
      }
      return 1 + this.left.size() + this.right.size();
    }

  }

  private Node root;

  public KdTree() {
    this.root = null;
  }

  public boolean isEmpty() {
    return this.root == null;
  }

  public int size() {
    if (this.root == null) {
      return 0;
    }
    return this.root.size();
  }

  public void insert(Point2D p) {
    if (p == null) {
      throw new IllegalArgumentException();
    }
    this.root = add(this.root, p, true);
  }

  public boolean contains(Point2D p) {
    if (p == null) {
      throw new IllegalArgumentException();
    }
    return search(this.root, p, true);
  }

  public void draw() {
    draw(this.root, true, new RectHV(0, 0, 1, 1));
  }

  public Iterable<Point2D> range(RectHV rect) {
    if (rect == null) {
      throw new IllegalArgumentException();
    }
    return range(this.root, rect, true, new RectHV(0, 0, 1, 1));
  }

  public Point2D nearest(Point2D p) {
    if (p == null) {
      throw new IllegalArgumentException();
    }
    Point2D closest = this.root.point;
    return nearest(this.root, p, closest, true, new RectHV(0, 0, 1, 1));
  }

  public static void main(String[] args) {

  }

  private Node add(Node n, Point2D p, boolean xDim) {
    if (p == null) {
      throw new IllegalArgumentException();
    }
    if (n == null) {
      return new Node(p);
    }
    if (n.point.equals(p)) {
      return null;
    }
    boolean isLess;
    if (xDim) {
      if (p.x() < n.point.x()) {
        isLess = true;
      } else {
        isLess = false;
      }
    } else {
      if (p.y() < n.point.y()) {
        isLess = true;
      } else {
        isLess = false;
      }
    }
    if (isLess) {
      n.left = add(n.left, p, !xDim);
    } else {
      n.right = add(n.right, p, !xDim);
    }
    return n;
  }

  private boolean search(Node n, Point2D p, boolean xDim) {
    if (p == null) {
      throw new IllegalArgumentException();
    }
    if (n == null) {
      return false;
    }
    if (n.point.equals(p)) {
      return true;
    }
    boolean isLess;
    if (xDim) {
      if (p.x() < n.point.x()) {
        isLess = true;
      } else {
        isLess = false;
      }
    } else {
      if (p.y() < n.point.y()) {
        isLess = true;
      } else {
        isLess = false;
      }
    }
    if (isLess) {
      return search(n.left, p, !xDim);
    } else {
      return search(n.right, p, !xDim);
    }
  }

  private void draw(Node n, boolean xDim, RectHV bounds) {
    if (n == null) {
      return;
    }
    StdDraw.setPenColor(StdDraw.BLACK);
    StdDraw.setPenRadius(0.01);
    StdDraw.point(n.point.x(), n.point.y());
    if (xDim) {
      StdDraw.setPenRadius();
      StdDraw.setPenColor(StdDraw.BLUE);
      StdDraw.line(n.point.x(), bounds.ymin(), n.point.x(), bounds.ymax());
      RectHV leftBounds = new RectHV(bounds.xmin(), bounds.ymin(), n.point.x(), bounds.ymax());
      draw(n.left, !xDim, leftBounds);
      RectHV rightBounds = new RectHV(n.point.x(), bounds.ymin(), bounds.xmax(), bounds.ymax());
      draw(n.right, !xDim, rightBounds);
    } else {
      StdDraw.setPenRadius();
      StdDraw.setPenColor(StdDraw.RED);
      StdDraw.line(bounds.xmin(), n.point.y(), bounds.xmax(), n.point.y());
      RectHV leftBounds = new RectHV(bounds.xmin(), bounds.ymin(), bounds.xmax(), n.point.y());
      draw(n.left, !xDim, leftBounds);
      RectHV rightBounds = new RectHV(bounds.xmin(), n.point.y(), bounds.xmax(), bounds.ymax());
      draw(n.right, !xDim, rightBounds);
    }
  }

  private ArrayList<Point2D> range(Node n, RectHV rect, boolean xDim, RectHV bounds) {
    if (n == null) {
      return new ArrayList<Point2D>();
    }
    ArrayList<Point2D> totalPoints = new ArrayList<>();
    if (rect.contains(n.point)) {
      totalPoints.add(n.point);
    }
    if (xDim) {
      RectHV leftBounds = new RectHV(bounds.xmin(), bounds.ymin(), n.point.x(), bounds.ymax());
      ArrayList<Point2D> leftPoints = new ArrayList<>();
      if (leftBounds.intersects(rect)) {
        leftPoints = range(n.left, rect, !xDim, leftBounds);
      }
      RectHV rightBounds = new RectHV(n.point.x(), bounds.ymin(), bounds.xmax(), bounds.ymax());
      ArrayList<Point2D> rightPoints = new ArrayList<>();
      if (rightBounds.intersects(rect)) {
        rightPoints = range(n.right, rect, !xDim, rightBounds);
      }
      totalPoints.addAll(leftPoints);
      totalPoints.addAll(rightPoints);
    } else {
      RectHV leftBounds = new RectHV(bounds.xmin(), bounds.ymin(), bounds.xmax(), n.point.y());
      ArrayList<Point2D> leftPoints = new ArrayList<>();
      if (leftBounds.intersects(rect)) {
        leftPoints = range(n.left, rect, !xDim, leftBounds);
      }
      RectHV rightBounds = new RectHV(bounds.xmin(), n.point.y(), bounds.xmax(), bounds.ymax());
      ArrayList<Point2D> rightPoints = new ArrayList<>();
      if (rightBounds.intersects(rect)) {
        rightPoints = range(n.right, rect, !xDim, rightBounds);
      }
      totalPoints.addAll(leftPoints);
      totalPoints.addAll(rightPoints);
    }
    return totalPoints;
  }

  private Point2D nearest(Node n, Point2D p, Point2D closest, boolean xDim, RectHV bounds) {
    if (n == null) {
      return closest;
    }
    Point2D newClosest = closest;
    double currDistSqr = p.distanceSquaredTo(n.point);
    double minDistSq = p.distanceSquaredTo(closest);
    if (currDistSqr < minDistSq) {
      newClosest = n.point;
      minDistSq = currDistSqr;
    }

    if (xDim) {
      RectHV leftBounds = new RectHV(bounds.xmin(), bounds.ymin(), n.point.x(), bounds.ymax());
      RectHV rightBounds = new RectHV(n.point.x(), bounds.ymin(), bounds.xmax(), bounds.ymax());
      if (p.x() < n.point.x()) {
        newClosest = nearest(n.left, p, newClosest, !xDim, leftBounds);
        double minDist = p.distanceTo(newClosest);
        if (minDist > Math.abs(p.x() - n.point.x())) {
          newClosest = nearest(n.right, p, newClosest, !xDim, rightBounds);
        }
      } else {
        newClosest = nearest(n.right, p, newClosest, !xDim, rightBounds);
        double minDist = p.distanceTo(newClosest);
        if (minDist > Math.abs(p.x() - n.point.x())) {
          newClosest = nearest(n.left, p, newClosest, !xDim, leftBounds);
        }
      }
    } else {
      RectHV leftBounds = new RectHV(bounds.xmin(), bounds.ymin(), bounds.xmax(), n.point.y());
      RectHV rightBounds = new RectHV(bounds.xmin(), n.point.y(), bounds.xmax(), bounds.ymax());
      if (p.y() < n.point.y()) {
        newClosest = nearest(n.left, p, newClosest, !xDim, leftBounds);
        double minDist = p.distanceTo(newClosest);
        if (minDist > Math.abs(p.y() - n.point.y())) {
          newClosest = nearest(n.right, p, newClosest, !xDim, rightBounds);
        }
      } else {
        newClosest = nearest(n.right, p, newClosest, !xDim, rightBounds);
        double minDist = p.distanceTo(newClosest);
        if (minDist > Math.abs(p.y() - n.point.y())) {
          newClosest = nearest(n.left, p, newClosest, !xDim, leftBounds);
        }
      }
    }
    return newClosest;
  }
}
