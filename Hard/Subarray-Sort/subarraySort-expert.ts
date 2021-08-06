// See how the expert solve this problem elegantly�

const subArraySort = (array: number[]): [number, number] => {
  let minOutOfOrder = Number.POSITIVE_INFINITY;
  let maxOutOfOrder = Number.NEGATIVE_INFINITY;
  for (let i = 0; i < array.length; i++) {
    const num = array[i];
    if (isOutOfOrder(i, num, array)) {
      minOutOfOrder = Math.min(minOutOfOrder, num);
      maxOutOfOrder = Math.max(maxOutOfOrder, num);
    }
  }
  if (minOutOfOrder === Number.POSITIVE_INFINITY) return [-1, -1];
  let subarrayLeftIdx = 0;
  while (minOutOfOrder >= array[subarrayLeftIdx]) {
    subarrayLeftIdx += 1;
  }
  let subarrayRightIdx = array.length - 1;
  while (maxOutOfOrder <= array[subarrayRightIdx]) {
    subarrayRightIdx -= 1;
  }
  return [subarrayLeftIdx, subarrayRightIdx];
};

const isOutOfOrder = (i: number, num: number, array: number[]): boolean => {
  if (i === 0) return num > array[i + 1];
  if (i === array.length - 1) return num < array[i - 1];
  return num > array[i + 1] || num < array[i - 1];
};

const input1 = [1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19];
const res1 = subArraySort(input1);
console.log(res1);

console.log(subArraySort([9, 8, 7, 6, 5, 4, 3, 2, 1]));
console.log(subArraySort([1, 2, 3, 4, 5, 6, 7, 8, 9]));
