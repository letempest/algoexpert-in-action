// Given an array (and possibly specified a starting index), jumping foward/backward with the value at the current index
// if finish iterating over every element exactly one time and stop at the starting index, then this is a single cycle
// Big O: O(n) time | O(1) space
const hasSingleCycle = (array: number[]): boolean => {
  const n = array.length;
  const STARTING_IDX = 0;
  let currentIdx = STARTING_IDX;
  for (let i = 0; i < array.length; i++) {
    // in the middle of iteration, if we're revisiting starting index already, then fail
    if (i !== 0 && currentIdx === STARTING_IDX) return false;
    currentIdx += array[currentIdx];
    if (currentIdx % n >= 0) {
      currentIdx = currentIdx % n;
    } else {
      // currentIdx is NEGATIVE, after taking modulo with n,
      // result could be ranging from 0,|  -1,  -2,..., -(n-2), -(n-1)
      //                to map into 0(n),| n-1, n-2,...,   2,      1
      // the fomula would be added with n, as follow
      currentIdx = (currentIdx % n) + n;
    }
    // console.log(i, currentIdx);
  }
  // after iterating all elements, we should be back to index 0(the starting point)
  return currentIdx === STARTING_IDX;
};

// true, 2 -> 1 -> -4 -> 2 -> 3 -> -1 -> 2
console.log(hasSingleCycle([2, 3, 1, -4, -4, 2]));
console.log(hasSingleCycle([1, -1, -1, 1]));
