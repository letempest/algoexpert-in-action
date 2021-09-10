class Heap {
  private heap: number[];
  length: number;

  constructor(
    private comparisonFunc: (a: number, b: number) => boolean,
    array: number[]
  ) {
    this.heap = this.buildHeap(array);
    this.comparisonFunc = comparisonFunc;
    this.length = this.heap.length;
  }

  private buildHeap = (array: number[]): number[] => {
    const firstParentIdx = Math.floor((array.length - 2) / 2);
    for (let currentIdx = firstParentIdx; currentIdx >= 0; currentIdx--) {
      this.siftDown(currentIdx, array.length - 1, array);
    }
    return array;
  };

  private siftDown = (
    currentIdx: number,
    endIdx: number,
    heap: number[]
  ): void => {
    let childOneIdx = currentIdx * 2 + 1;
    while (childOneIdx <= endIdx) {
      const childTwoIdx =
        currentIdx * 2 + 2 <= endIdx ? currentIdx * 2 + 2 : -1;
      let idxToSwap: number;
      if (childTwoIdx !== -1) {
        if (this.comparisonFunc(heap[childTwoIdx], heap[childOneIdx])) {
          idxToSwap = childTwoIdx;
        } else {
          idxToSwap = childOneIdx;
        }
      } else {
        idxToSwap = childOneIdx;
      }
      if (this.comparisonFunc(heap[idxToSwap], heap[currentIdx])) {
        this.swap(currentIdx, idxToSwap, heap);
        currentIdx = idxToSwap;
        childOneIdx = currentIdx * 2 + 1;
      } else {
        return;
      }
    }
  };

  private siftUp = (currentIdx: number, heap: number[]): void => {
    let parentIdx = Math.floor((currentIdx - 1) / 2);
    while (currentIdx > 0) {
      if (this.comparisonFunc(heap[currentIdx], heap[parentIdx])) {
        this.swap(currentIdx, parentIdx, heap);
        currentIdx = parentIdx;
        parentIdx = Math.floor((currentIdx - 1) / 2);
      } else {
        return;
      }
    }
  };

  private swap = (i: number, j: number, array: number[]): void => {
    [array[i], array[j]] = [array[j], array[i]];
  };

  peek = (): number => {
    return this.heap[0];
  };

  remove = (): number => {
    this.swap(0, this.length - 1, this.heap);
    const valueToRemove = this.heap.pop()!;
    this.length -= 1;
    this.siftDown(0, this.length - 1, this.heap);
    return valueToRemove;
  };

  insert = (value: number): void => {
    this.heap.push(value);
    this.length += 1;
    this.siftUp(this.length - 1, this.heap);
  };
}

const maxHeapFunc = (a: number, b: number): boolean => {
  return a > b;
};

const minHeapFunc = (a: number, b: number): boolean => {
  return a < b;
};

export { Heap, maxHeapFunc, minHeapFunc };
