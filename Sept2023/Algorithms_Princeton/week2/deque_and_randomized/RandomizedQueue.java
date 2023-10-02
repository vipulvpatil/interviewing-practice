import java.util.Iterator;
import edu.princeton.cs.algs4.StdRandom;

public class RandomizedQueue<Item> implements Iterable<Item> {

  private Item[] items;
  private int capacity;
  private int length;

  // construct an empty randomized queue
  public RandomizedQueue() {
    this.items = (Item[]) new Object[1];
    this.capacity = 1;
    this.length = 0;
  }

  // is the randomized queue empty?
  public boolean isEmpty() {
    return this.length == 0;
  }

  // return the number of items on the randomized queue
  public int size() {
    return this.length;
  }

  // add the item
  public void enqueue(Item item) {
    if (item == null) {
      throw new IllegalArgumentException();
    }
    if (this.capacity <= this.length) {
      resize(this.capacity * 2);
    }
    this.items[this.length++] = item;
  }

  // remove and return a random item
  public Item dequeue() {
    if (this.length == 0) {
      throw new java.util.NoSuchElementException();
    }
    int randomIndex = StdRandom.uniformInt(length);
    Item item = this.items[randomIndex];
    for (int i = randomIndex; i < this.length - 1; i++) {
      this.items[i] = this.items[i + 1];
    }
    this.items[--this.length] = null;
    if (this.capacity >= this.length * 4) {
      resize(this.capacity / 2);
    }
    return item;
  }

  // return a random item (but do not remove it)
  public Item sample() {
    if (this.length == 0) {
      throw new java.util.NoSuchElementException();
    }
    int randomIndex = StdRandom.uniformInt(length);
    return this.items[randomIndex];
  }

  // return an independent iterator over items in random order
  public Iterator<Item> iterator() {
    return new MyIterator(this.items, this.length);
  }

  private class MyIterator implements Iterator<Item> {
    RandomizedQueue<Item> copy;

    public MyIterator(Item[] items, int length) {
      this.copy = new RandomizedQueue<Item>();
      for (int i = 0; i < length; i++) {
        this.copy.enqueue(items[i]);
      }
    }

    public boolean hasNext() {
      return !this.copy.isEmpty();
    }

    public void remove() {
      throw new UnsupportedOperationException();
    }

    public Item next() {
      if (hasNext()) {
        return this.copy.dequeue();
      }
      throw new java.util.NoSuchElementException();
    }
  }

  // unit testing (required)
  public static void main(String[] args) {
    RandomizedQueue<Integer> list = new RandomizedQueue<Integer>();
    list.enqueue(1);
    list.enqueue(2);
    list.enqueue(3);
    list.enqueue(4);
    list.enqueue(5);

    Iterator<Integer> iter = list.iterator();
    while (iter.hasNext()) {
      System.out.println(iter.next());
    }

    System.out.println(list.dequeue());
    System.out.println(list.dequeue());
    System.out.println(list.dequeue());
    System.out.println(list.dequeue());
    System.out.println(list.dequeue());
    System.out.println(list.dequeue());
  }

  private void resize(int newCapacity) {
    if (newCapacity == 0) {
      return;
    }
    Item[] copy = (Item[]) new Object[newCapacity];
    for (int i = 0; i < newCapacity; i++) {
      if (i < length) {
        copy[i] = items[i];
      }
    }
    capacity = newCapacity;
    items = copy;
  }
}
