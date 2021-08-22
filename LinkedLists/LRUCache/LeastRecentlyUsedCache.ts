// Least Recently Used Cache with a capacity, require that 1) insertion, 2) lookup value,
// 3) get most recently used cache key all yields a O(1) space/time complexity
// cache implemented with hash table for constant time complexity
// for the cache key-value pair, value is of type DoublyLinkedListNode,
// so the feature least/most recently used could be achieved through updating the doubly linked list head/tail
class LRUCache {
  cache: { [key: string]: DoublyLinkedListNode } = {};
  currentSize: number = 0;
  listOfMostRecent: DoublyLinkedList;

  constructor(public readonly maxSize: number = 1) {
    this.listOfMostRecent = new DoublyLinkedList();
  }

  // O(1) time | O(1) space
  insertKeyValuePair = (key: string, value: any): void => {
    if (key in this.cache === false) {
      if (this.currentSize === this.maxSize) {
        this.evictLeastRecent();
      } else {
        this.currentSize += 1;
      }
      this.cache[key] = new DoublyLinkedListNode(key, value);
    } else {
      this.replaceKey(key, value);
    }
    this.updateMostRecent(this.cache[key]);
  };

  // O(1) time | O(1) space
  getValueFromKey = (key: string): any => {
    if (key in this.cache === false) {
      return;
    }
    const node = this.cache[key];
    this.updateMostRecent(node);
    return node.value;
  };

  // O(1) time | O(1) space
  getMostRecentKey = (): any => {
    return this.listOfMostRecent.head?.key;
  };

  private replaceKey = (key: string, value: any): void => {
    if (key in this.cache === false) {
      throw new Error('The provided key is not in the cache!');
    }
    this.cache[key].value = value;
  };

  private evictLeastRecent = (): void => {
    const keyToRemove = this.listOfMostRecent.tail!.key;
    delete this.cache[keyToRemove];
    this.listOfMostRecent.removeTail();
  };

  private updateMostRecent = (node: DoublyLinkedListNode): void => {
    this.listOfMostRecent.setHeadTo(node);
  };
}

class DoublyLinkedListNode {
  prev: DoublyLinkedListNode | null = null;
  next: DoublyLinkedListNode | null = null;
  constructor(public key: string, public value: any) {}

  removeBindings = (): void => {
    if (this.prev) {
      this.prev.next = this.next;
    }
    if (this.next) {
      this.next.prev = this.prev;
    }
    this.prev = null;
    this.next = null;
  };
}

class DoublyLinkedList {
  head: DoublyLinkedListNode | null = null;
  tail: DoublyLinkedListNode | null = null;

  setHeadTo = (node: DoublyLinkedListNode): void => {
    if (this.head === node) return;
    if (!this.head) {
      this.head = node;
      this.tail = node;
    } else if (this.head === this.tail) {
      this.tail.prev = node;
      this.head = node;
      this.head.next = this.tail;
    } else {
      if (this.tail === node) {
        this.removeTail();
      }
      node.removeBindings();
      this.head.prev = node;
      node.next = this.head;
      this.head = node;
    }
  };

  removeTail = (): void => {
    if (!this.tail) return;
    if (this.tail === this.head) {
      this.head = null;
      this.tail = null;
      return;
    }
    this.tail = this.tail.prev;
    this.tail!.next = null;
  };
}

const cache = new LRUCache(3);
console.log(cache.getMostRecentKey());

cache.insertKeyValuePair('a', 1);
cache.insertKeyValuePair('b', 2);
cache.insertKeyValuePair('c', 3);
console.log(cache.getMostRecentKey());

cache.insertKeyValuePair('x', 100);
console.log(cache.getMostRecentKey());
console.log(cache);
cache.insertKeyValuePair('b', 97);
console.log(cache.getMostRecentKey());
