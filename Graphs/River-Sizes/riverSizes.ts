// My depth-first search Recursive Solution
// Big O: O(wh) time, where w, h is the dimensions of the matrix, = O(n) where n is the total number of nodes
// O(wh) space for the auxiliary visited matrix; also the worst case scenario O(wh) recursive calls on the callstack
const riverSizes = (matrix: number[][]): number[] => {
  const visited: boolean[][] = matrix.map(arr =>
    new Array(arr.length).fill(false)
  );
  const sizes: number[] = [];

  for (let i = 0; i < matrix.length; i++) {
    for (let j = 0; j < matrix[i].length; j++) {
      // Every call iterative call below is dealing with a new river, if current matrix value is 1, this is checked in helper func
      const riverLength = traverseNode(matrix, visited, 0, i, j);
      if (riverLength > 0) {
        sizes.push(riverLength);
      }
    }
  }
  return sizes.sort((a, b) => a - b);
};

//
const traverseNode = (
  matrix: number[][],
  visited: boolean[][],
  riverLength: number,
  i: number,
  j: number
): number => {
  if (visited[i][j]) {
    return riverLength;
  }

  visited[i][j] = true;
  if (matrix[i][j] === 0) {
    return riverLength;
  } else {
    riverLength += 1;
  }

  if (i > 0) {
    riverLength = traverseNode(matrix, visited, riverLength, i - 1, j);
  }
  if (i < matrix.length - 1) {
    riverLength = traverseNode(matrix, visited, riverLength, i + 1, j);
  }
  if (j > 0) {
    riverLength = traverseNode(matrix, visited, riverLength, i, j - 1);
  }
  if (j < matrix[0].length - 1) {
    riverLength = traverseNode(matrix, visited, riverLength, i, j + 1);
  }
  return riverLength;
};

// ================================================================
// Expert's depth-first search Iterative Solution,
// theoretically better space complexity due to its iterative nature, so no additional calls awaiting on the callstack
// Big O: O(wh) time | O(wh) space
// const riverSizes = (matrix: number[][]): number[] => {
//   const sizes: number[] = [];
//   const visited: boolean[][] = matrix.map(arr =>
//     new Array(arr.length).fill(false)
//   );

//   for (let i = 0; i < matrix.length; i++) {
//     for (let j = 0; j < matrix[i].length; j++) {
//       if (visited[i][j]) {
//         continue;
//       }
//       traverseNode(i, j, matrix, visited, sizes);
//     }
//   }
//   return sizes.sort((a, b) => a - b);
// };

// const traverseNode = (
//   i: number,
//   j: number,
//   matrix: number[][],
//   visited: boolean[][],
//   sizes: number[]
// ): void => {
//   let currentRiverSize = 0;
//   // a stack, same node could be pushed into it multiple time due to graph connection, so not guarantee to be unvisited
//   let nodesToExplore = [[i, j]];
//   while (nodesToExplore.length > 0) {
//     const currentNode = nodesToExplore.pop()!;
//     const i = currentNode[0];
//     const j = currentNode[1];
//     if (visited[i][j]) {
//       continue;
//     }
//     visited[i][j] = true;
//     if (matrix[i][j] === 0) {
//       continue;
//     }
//     if (matrix[i][j] === 1) {
//       currentRiverSize += 1;
//     }
//     const unvisitedNeighbors = getUnvisitedNeighbors(i, j, matrix, visited);
//     nodesToExplore = [...nodesToExplore, ...unvisitedNeighbors];
//   }
//   if (currentRiverSize > 0) {
//     sizes.push(currentRiverSize);
//   }
// };

// const getUnvisitedNeighbors = (
//   i: number,
//   j: number,
//   matrix: number[][],
//   visited: boolean[][]
// ): number[][] => {
//   const unvisitedNeighbors: number[][] = [];
//   if (i - 1 >= 0 && !visited[i - 1][j]) {
//     unvisitedNeighbors.push([i - 1, j]);
//   }
//   if (i < matrix.length - 1 && !visited[i + 1][j]) {
//     unvisitedNeighbors.push([i + 1, j]);
//   }
//   if (j - 1 >= 0 && !visited[i][j - 1]) {
//     unvisitedNeighbors.push([i, j - 1]);
//   }
//   if (j < matrix[i].length - 1 && !visited[i][j + 1]) {
//     unvisitedNeighbors.push([i, j + 1]);
//   }
//   return unvisitedNeighbors;
// };

const matrix = new Array(5).fill(0).map(() => new Array(5).fill(0));
matrix[0] = [1, 0, 0, 1, 0];
matrix[1] = [1, 0, 1, 0, 0];
matrix[2] = [0, 0, 1, 0, 1];
matrix[3] = [1, 0, 1, 0, 1];
matrix[4] = [1, 0, 1, 1, 0];

const rivers = riverSizes(matrix);
console.log(rivers);
