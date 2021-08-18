// Given an array of number representing building height, and two possible viewing direction: West / East
// return an array sorted by original building index, in which each element can see the sunset without being blocked by higher building

import Stack from '../Stack';

enum Direction {
  WEST = 'WEST',
  EAST = 'EAST'
}

// Solution 1
// Big O: O(n) time | O(n) space
// const sunsetViews = (buildings: number[], direction: Direction): number[] => {
//   let buildingsWithSunsetViews: number[] = [];
//   let runningMaxHeight = 0;

//   const startIdx = direction === Direction.WEST ? 0 : buildings.length - 1;
//   const step = direction === Direction.WEST ? 1 : -1;
//   let idx = startIdx;
//   while (idx >= 0 && idx < buildings.length) {
//     const buildingHeight = buildings[idx];
//     if (buildingHeight > runningMaxHeight) {
//       buildingsWithSunsetViews.push(idx);
//     }
//     runningMaxHeight = Math.max(buildingHeight, runningMaxHeight);
//     idx += step;
//   }
//   return direction === Direction.WEST
//     ? buildingsWithSunsetViews
//     : buildingsWithSunsetViews.reverse();
// };

// Solution 2
// Big O: O(n) time | O(n) space
const sunsetViews = (buildings: number[], direction: Direction): number[] => {
  const candidateBuildings = new Stack<number>();

  const startIdx = direction === Direction.EAST ? 0 : buildings.length - 1;
  const step = direction === Direction.EAST ? 1 : -1;
  let idx = startIdx;
  while (idx >= 0 && idx < buildings.length) {
    const buildingHeight = buildings[idx];

    // keep poping element in the stack if current building is higher than the last building in the stack (LIFO!)
    while (
      candidateBuildings.stack.length > 0 &&
      buildingHeight >= buildings[candidateBuildings.peak()]
    ) {
      candidateBuildings.pop();
    }

    candidateBuildings.push(idx);
    idx += step;
  }

  return direction === Direction.EAST
    ? candidateBuildings.stack
    : candidateBuildings.stack.reverse();
};

//
//          |
//          |   |   |
//      |   |   |   |   |       |
//      |   |   |   |   |       |   |
//      |   |   |   |   |   |   |   |
//      .------------------------------->
//      0   1   2   3   4   5   6   7
//   West           --->             East

console.log(sunsetViews([3, 5, 4, 4, 3, 1, 3, 2], Direction.EAST));
console.log(sunsetViews([3, 5, 4, 4, 3, 1, 3, 2], Direction.WEST));
