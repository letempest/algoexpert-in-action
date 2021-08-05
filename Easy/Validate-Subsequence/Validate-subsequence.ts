// To find out if one array is the subsequence of another array
// big O: O(n) Time O(1) space
const origSequence = [5, 1, 22, 25, 6, -1, 8, 10, 123];
const subSequence = [1, 6, -1, 10, 123];

// SOLUTION 1: recursion
const isSubSequence1 = (arr1: number[], arr2: number[]): boolean => {
  if (arr1.length >= 0 && arr2.length === 0) return true;

  let pointer = arr1.findIndex(val => val === arr2[0]);
  if (pointer === -1 || arr1.length - pointer < arr2.length) return false;

  return isSubSequence1(arr1.slice(pointer), arr2.slice(1));
};

// SOLUTION 2: pointer traversing
const isSubSequence2 = (array: number[], sequence: number[]): boolean => {
  let arrIdx = 0,
    seqIdx = 0;

  while (arrIdx < array.length && seqIdx < sequence.length) {
    if (array[arrIdx] === sequence[seqIdx]) seqIdx++;
    arrIdx++;
  }
  return seqIdx === sequence.length;
};

// SOLUTION 3
const isSubSequence = (array: number[], sequence: number[]): boolean => {
  let seqIdx = 0;

  for (const value of array) {
    if (seqIdx === sequence.length) break;
    if (value === sequence[seqIdx]) seqIdx++;
  }
  return seqIdx === sequence.length;
};

console.log(isSubSequence(origSequence, subSequence));
console.log(isSubSequence([1, 2, 3], [1, 2]));
console.log(isSubSequence([3, 7, 2, 1, 4], [7, 3, 1]));
console.log(isSubSequence([-5, 12, 3, 7, 2, 4, 0, -28, 4, 3], [12, 2, 0, 4]));
