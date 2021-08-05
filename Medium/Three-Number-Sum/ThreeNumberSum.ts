// Give an unsorted array, return all triplet arrays with three elements that sum up to equal the targetSum number.
// Big O: O(n^2) Time | O(n) Space
const arr = [12, 3, 1, 2, -6, 5, -8, 6];

const threeNumberSum = (arr: number[], targetSum: number): number[][] => {
  arr.sort((a, b) => a - b);

  let left: number, right: number;
  const triplets: number[][] = [];

  for (let i = 0; i < arr.length - 2; i++) {
    left = i + 1;
    right = arr.length - 1;
    while (left < right) {
      const sum = arr[i] + arr[left] + arr[right];
      if (sum < targetSum) {
        left += 1;
      } else if (sum > targetSum) {
        right -= 1;
      } else {
        triplets.push([arr[i], arr[left], arr[right]]);
        left++;
        right--;
      }
    }
  }

  return triplets;
};

console.log(threeNumberSum(arr, 0));
