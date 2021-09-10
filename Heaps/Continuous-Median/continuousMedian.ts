// With an array of continuously incoming numbers, return the running median of the current array.
// optimal solution: at any point of new value appending, insert the new value to one of two auxiliary heaps,
// a) a Max heap for lower half value of current array; b) a Min heap for greater half value of current array.
// with the property of min/max heap, we can keep track of two values closest to middle index, which will either
// be the running median itself (if current array lenth is an odd number); or the average of these two numbers
// is the running median (if current array length is even), always be sure that difference in length of the two
// heaps is <= 1, if difference reaches 2, the two heaps have to be rebalanced to maintain somewhat roughly
// equal length; rebalancing two heaps means removing top value from the larger heap, and insert the same value
// into the smaller heap, two O(log(n)) operation

import { Heap, minHeapFunc, maxHeapFunc } from './Heap';

export class ContinuousMedianHandler {
  private lowers: Heap; // max heap
  private greaters: Heap; // min heap
  private _median?: number;

  constructor() {
    this.lowers = new Heap(maxHeapFunc, []);
    this.greaters = new Heap(minHeapFunc, []);
  }

  // O(log(n)) time | O(n) space for storing all numbers in two heaps
  insert = (value: number): void => {
    const lowerHalfMax = this.lowers.peek();
    if (this.lowers.length === 0 || value < lowerHalfMax) {
      this.lowers.insert(value);
    } else {
      this.greaters.insert(value);
    }

    this.rebalanceHeaps();
    this.updateMedian();
  };

  private rebalanceHeaps = (): void => {
    if (this.greaters.length - this.lowers.length === 2) {
      this.lowers.insert(this.greaters.remove());
    } else if (this.lowers.length - this.greaters.length === 2) {
      this.greaters.insert(this.lowers.remove());
    }
  };

  private updateMedian = (): void => {
    if (this.greaters.length === this.lowers.length) {
      this._median = (this.greaters.peek() + this.lowers.peek()) / 2;
    } else if (this.greaters.length > this.lowers.length) {
      this._median = this.greaters.peek();
    } else {
      this._median = this.lowers.peek();
    }
  };

  get median(): number | undefined {
    return this._median;
  }
}

const continuousMedianHandler = new ContinuousMedianHandler();

// Should print out 5, 7.5, 10, 55, 10, 11.5, 13, 13.5
for (const val of [5, 10, 100, 200, 6, 13, 14, 1000]) {
  continuousMedianHandler.insert(val);
  console.log(continuousMedianHandler.median);
}
