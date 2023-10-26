import edu.princeton.cs.algs4.Picture;
import edu.princeton.cs.algs4.Stack;
import edu.princeton.cs.algs4.StdOut;

import java.awt.Color;

public class SeamCarver {
  private Picture p;

  // create a seam carver object based on the given picture
  public SeamCarver(Picture picture) {
    if (picture == null) {
      throw new IllegalArgumentException();
    }
    p = new Picture(picture);
  }

  // current picture
  public Picture picture() {
    return new Picture(p);
  }

  // width of current picture
  public int width() {
    return p.width();
  }

  // height of current picture
  public int height() {
    return p.height();
  }

  // energy of pixel at column x and row y
  public double energy(int x, int y) {
    if (x < 0 || y < 0 || x > p.width() - 1 || y > p.height() - 1) {
      throw new IllegalArgumentException();
    }
    if (x == 0 || y == 0 || x == p.width() - 1 || y == p.height() - 1) {
      return 1000;
    }

    Color colorXPlus = p.get(x + 1, y);
    Color colorXMinus = p.get(x - 1, y);
    int rx = colorXPlus.getRed() - colorXMinus.getRed();
    int gx = colorXPlus.getGreen() - colorXMinus.getGreen();
    int bx = colorXPlus.getBlue() - colorXMinus.getBlue();

    Color colorYPlus = p.get(x, y + 1);
    Color colorYMinus = p.get(x, y - 1);
    int ry = colorYPlus.getRed() - colorYMinus.getRed();
    int gy = colorYPlus.getGreen() - colorYMinus.getGreen();
    int by = colorYPlus.getBlue() - colorYMinus.getBlue();

    double deltaXSq = Math.pow(rx, 2) + Math.pow(gx, 2) + Math.pow(bx, 2);
    double deltaYSq = Math.pow(ry, 2) + Math.pow(gy, 2) + Math.pow(by, 2);

    return Math.sqrt(deltaXSq + deltaYSq);
  }

  // sequence of indices for horizontal seam
  public int[] findHorizontalSeam() {
    if (p.width() < 1) {
      throw new IllegalArgumentException();
    }
    Pixel[][] horizontalPaths = new Pixel[p.width()][p.height()];
    for (int y = 0; y < p.height(); y++) {
      horizontalPaths[0][y] = new Pixel(0, y, energy(0, y), null);
    }
    for (int x = 1; x < p.width(); x++) {
      for (int y = 0; y < p.height(); y++) {
        Pixel p1 = pixelAt(horizontalPaths, x - 1, y);
        Pixel p2 = pixelAt(horizontalPaths, x - 1, y - 1);
        Pixel p3 = pixelAt(horizontalPaths, x - 1, y + 1);
        Pixel minPixel = getMinimumTotalEnergyPixel(p1, p2, p3);
        horizontalPaths[x][y] = new Pixel(x, y, minPixel.totalEnergy() + energy(x, y), minPixel);
      }
    }
    Pixel minTotalEnergyPixel = pixelAt(horizontalPaths, p.width() - 1, 0);
    for (int y = 1; y < p.height(); y++) {
      Pixel pixel = pixelAt(horizontalPaths, p.width() - 1, y);
      if (pixel.totalEnergy() < minTotalEnergyPixel.totalEnergy()) {
        minTotalEnergyPixel = pixel;
      }
    }

    Stack<Integer> st = new Stack<>();
    Pixel pixel = minTotalEnergyPixel;
    while (pixel != null) {
      st.push(pixel.y());
      pixel = pixel.prev();
    }
    int[] arr = new int[st.size()];
    int index = 0;
    for (int e : st) {
      arr[index++] = e;
    }
    return arr;
  }

  // sequence of indices for vertical seam
  public int[] findVerticalSeam() {
    if (p.height() < 1) {
      throw new IllegalArgumentException();
    }
    Pixel[][] verticalPaths = new Pixel[p.width()][p.height()];
    for (int x = 0; x < p.width(); x++) {
      verticalPaths[x][0] = new Pixel(x, 0, energy(x, 0), null);
    }
    for (int y = 1; y < p.height(); y++) {
      for (int x = 0; x < p.width(); x++) {
        Pixel p1 = pixelAt(verticalPaths, x, y - 1);
        Pixel p2 = pixelAt(verticalPaths, x - 1, y - 1);
        Pixel p3 = pixelAt(verticalPaths, x + 1, y - 1);
        Pixel minPixel = getMinimumTotalEnergyPixel(p1, p2, p3);
        verticalPaths[x][y] = new Pixel(x, y, minPixel.totalEnergy() + energy(x, y), minPixel);
      }
    }
    Pixel minTotalEnergyPixel = pixelAt(verticalPaths, 0, p.height() - 1);
    for (int x = 1; x < p.width(); x++) {
      Pixel pixel = pixelAt(verticalPaths, x, p.height() - 1);
      if (pixel.totalEnergy() < minTotalEnergyPixel.totalEnergy()) {
        minTotalEnergyPixel = pixel;
      }
    }

    Stack<Integer> st = new Stack<>();
    Pixel pixel = minTotalEnergyPixel;
    while (pixel != null) {
      st.push(pixel.x());
      pixel = pixel.prev();
    }
    int[] arr = new int[st.size()];
    int index = 0;
    for (int e : st) {
      arr[index++] = e;
    }
    return arr;
  }

  // remove horizontal seam from current picture
  public void removeHorizontalSeam(int[] seam) {
    if (seam == null) {
      throw new IllegalArgumentException();
    }
    if (p.height() == 1) {
      throw new IllegalArgumentException();
    }
    if (seam.length != p.width()) {
      throw new IllegalArgumentException();
    }
    for (int i = 0; i < seam.length - 1; i++) {
      if (seam[i] > p.height() - 1 || seam[i] < 0) {
        throw new IllegalArgumentException();
      }
      if (Math.abs(seam[i] - seam[i + 1]) > 1) {
        throw new IllegalArgumentException();
      }
    }
    if (seam[seam.length - 1] > p.height() - 1 || seam[seam.length - 1] < 0) {
      throw new IllegalArgumentException();
    }
    Picture newP = new Picture(p.width(), p.height() - 1);
    for (int x = 0; x < p.width(); x++) {
      for (int y = 0; y < p.height() - 1; y++) {
        if (y < seam[x]) {
          newP.setRGB(x, y, p.getRGB(x, y));
        } else {
          newP.setRGB(x, y, p.getRGB(x, y + 1));
        }
      }
    }
    this.p = newP;
  }

  // remove vertical seam from current picture
  public void removeVerticalSeam(int[] seam) {
    if (seam == null) {
      throw new IllegalArgumentException();
    }
    if (p.width() == 1) {
      throw new IllegalArgumentException();
    }
    if (seam.length != p.height()) {
      throw new IllegalArgumentException();
    }
    for (int i = 0; i < seam.length - 1; i++) {
      if (seam[i] > p.width() - 1 || seam[i] < 0) {
        throw new IllegalArgumentException();
      }
      if (Math.abs(seam[i] - seam[i + 1]) > 1) {
        throw new IllegalArgumentException();
      }
    }
    if (seam[seam.length - 1] > p.width() - 1 || seam[seam.length - 1] < 0) {
      throw new IllegalArgumentException();
    }

    Picture newP = new Picture(p.width() - 1, p.height());
    for (int y = 0; y < p.height(); y++) {
      for (int x = 0; x < p.width() - 1; x++) {
        if (x < seam[y]) {
          newP.setRGB(x, y, p.getRGB(x, y));
        } else {
          newP.setRGB(x, y, p.getRGB(x + 1, y));
        }
      }
    }
    this.p = newP;
  }

  // unit testing (optional)
  public static void main(String[] args) {
    Picture picture = new Picture(args[0]);
    SeamCarver sm = new SeamCarver(picture);
    int[] hSeam = sm.findHorizontalSeam();
    int[] vSeam = sm.findVerticalSeam();
    for (int i = 0; i < vSeam.length; i++) {
      StdOut.printf("%d, ", vSeam[i]);
    }
    StdOut.println();
    for (int i = 0; i < hSeam.length; i++) {
      StdOut.printf("%d, ", hSeam[i]);
    }
  }

  private class Pixel {
    private final int x;
    private final int y;
    private final double totalEnergy;
    private final Pixel prev;

    public Pixel(int x, int y, double totalEnergy, Pixel prev) {
      this.x = x;
      this.y = y;
      this.totalEnergy = totalEnergy;
      this.prev = prev;
    }

    public int x() {
      return x;
    }

    public int y() {
      return y;
    }

    public double totalEnergy() {
      return this.totalEnergy;
    }

    public Pixel prev() {
      return this.prev;
    }
  }

  private static Pixel pixelAt(Pixel[][] paths, int x, int y) {
    if (x < 0 || y < 0 || x > paths.length - 1 || y > paths[0].length - 1) {
      return null;
    }
    return paths[x][y];
  }

  private static Pixel getMinimumTotalEnergyPixel(Pixel p1, Pixel p2, Pixel p3) {
    Pixel minPixel = p1;
    double minEnergy = totalEnergyOf(p1);
    if (totalEnergyOf(p2) < minEnergy) {
      minPixel = p2;
      minEnergy = totalEnergyOf(p2);
    }
    if (totalEnergyOf(p3) < minEnergy) {
      minPixel = p3;
    }
    return minPixel;
  }

  private static double totalEnergyOf(Pixel p) {
    if (p == null) {
      return Double.POSITIVE_INFINITY;
    }
    return p.totalEnergy();
  }
}
