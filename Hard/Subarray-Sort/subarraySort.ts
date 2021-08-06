// Given an unsorted array, find out the longest subarray which if this subarray is sorted in place, the original
// array would also be sorted. the algorithm should return a tuple [subarrayStartIdx, subarrayEndIdx]
// e.g. for the below input, the result should be [3, 9], i.e. subarray = originArr.slice(3, 9+1)
const input = [1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19];

// my first try, w/out helper func,it works but focusing on edge cases too much, the logic is obscure�
const subarraySort = (arr: number[]): [number, number] => {
  let cursor: number = 0;
  let left: number | null = null,
    right: number | null = null;

  while (cursor < arr.length) {
    // base case: left
    if (cursor === 0) {
      if (arr[cursor] <= arr[cursor + 1]) {
        cursor++;
      } else {
        left = cursor;
        cursor += 2;
      }
      continue;
    }

    // base case: right
    if (cursor === arr.length - 1) {
      if (arr[cursor] < arr[cursor - 1]) {
        right = cursor;
      }
      break;
    }

    // cursor > 0
    if (arr[cursor] >= arr[cursor - 1] && arr[cursor] <= arr[cursor + 1]) {
      cursor++;
      continue;
    } else {
      if (left === null) left = cursor;
      if (right === null || right < cursor) right = cursor + 1;
      cursor += 2;
    }
  }
  if (left === null || right === null) return [-1, -1];

  const min = Math.min(...arr.slice(left, right + 1));
  const max = Math.max(...arr.slice(left, right + 1));

  let start = arr.findIndex((v, idx) => v < min && min < arr[idx + 1]) + 1;
  let end: number;
  if (arr[right + 1] >= max) {
    end = right;
  } else {
    end =
      arr
        .slice(right + 1)
        .findIndex((v, idx, a) => v < max && max <= a[idx + 1]) +
      right +
      1;
  }

  return [start, end];
};

const result = subarraySort(input);
console.log(result);
console.log('---------------------');
// total reverse edge case
const input2 = [9, 8, 7, 6, 5, 4, 3, 2, 1];
const result2 = subarraySort(input2);
console.log(result2);
console.log('---------------------');
// already sorted array
const input3 = [1, 2, 3, 4, 5, 6, 7, 8, 9];
const result3 = subarraySort(input3);
console.log(result3);
