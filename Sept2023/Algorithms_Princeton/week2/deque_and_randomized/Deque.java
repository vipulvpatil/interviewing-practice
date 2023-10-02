import java.util.Iterator;

public class Deque<Item> implements Iterable<Item> {

  private Node head;
  private Node tail;
  private int count;

  private class Node {
    Node next;
    Node prev;
    Item value;

    public Node(Item item) {
      value = item;
    }
  }

  // construct an empty deque
  public Deque() {
    this.head = null;
    this.tail = null;
    this.count = 0;
  }

  // is the deque empty?
  public boolean isEmpty() {
    return this.count == 0;
  }

  // return the number of items on the deque
  public int size() {
    return this.count;
  }

  // add the item to the front
  public void addFirst(Item item) {
    if (item == null) {
      throw new IllegalArgumentException();
    }
    Node newNode = new Node(item);
    newNode.next = this.head;
    if (this.head != null) {
      this.head.prev = newNode;
    }
    this.head = newNode;
    if (this.count == 0) {
      this.tail = newNode;
    }
    this.count++;
  }

  // add the item to the back
  public void addLast(Item item) {
    if (item == null) {
      throw new IllegalArgumentException();
    }
    Node newNode = new Node(item);
    newNode.prev = this.tail;
    if (this.tail != null) {
      this.tail.next = newNode;
    }
    this.tail = newNode;
    if (this.count == 0) {
      this.head = newNode;
    }
    this.count++;
  }

  // remove and return the item from the front
  public Item removeFirst() {
    if (this.head == null) {
      throw new java.util.NoSuchElementException();
    }
    Item item = this.head.value;
    if (this.tail == this.head) {
      this.head = null;
      this.tail = null;
    } else {
      this.head = this.head.next;
      if (this.head != null) {
        this.head.prev = null;
      }
    }
    this.count--;
    return item;
  }

  // remove and return the item from the back
  public Item removeLast() {
    if (this.tail == null) {
      throw new java.util.NoSuchElementException();
    }
    Item item = this.tail.value;
    if (this.tail == this.head) {
      this.head = null;
      this.tail = null;
    } else {
      this.tail = this.tail.prev;
      if (this.tail != null) {
        this.tail.next = null;
      }
    }
    this.count--;
    return item;
  }

  // return an iterator over items in order from front to back
  public Iterator<Item> iterator() {
    return new ListIterator();
  }

  private class ListIterator implements Iterator<Item> {
    private Node current = head;

    public boolean hasNext() {
      return current != null;
    }

    public void remove() {
      throw new UnsupportedOperationException();
    }

    public Item next() {
      if (!hasNext()) {
        throw new java.util.NoSuchElementException();
      }
      Item item = current.value;
      current = current.next;
      return item;
    }
  }

  // unit testing (required)
  public static void main(String[] args) {
    // Deque<String> deque = new Deque<String>();
    // deque.addFirst("two");
    // deque.addFirst("one");
    // deque.addLast("three");

    // Iterator<String> dequeIterator = deque.iterator();
    // while (dequeIterator.hasNext()) {
    // System.out.println(dequeIterator.next());
    // }

    // System.out.println(deque.size());
    // System.out.println("---");

    // System.out.println(deque.removeFirst());
    // System.out.println(deque.removeLast());
    // System.out.println(deque.removeFirst());
    // System.out.println(deque.isEmpty());

    //
    Deque<Integer> deque = new Deque<Integer>();
    deque.addFirst(1);
    System.out.println(deque.iterator().hasNext());
    System.out.println(deque.removeLast());
    System.out.println(deque.iterator().hasNext());
  }

}
