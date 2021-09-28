// Given a matrix of currency exchange rates, determine whether there's arbitrage
// by transfering from 1 unit of any currency to others, and eventually transfer back
// to an amount greater than 1 unit of the original currency.
// 2d matrix probleam -> graph problem, we can convert the exchange rates into
// edges of a fully-connected weighted graph, run bellman-ford algorithm on each node,
// bellman-ford runs in O(v*e) time, and for a fully-connected graph, v = e,
// after that, run bellman-ford once more to see whether the shortest path
// has any update, if updated, we know we got a negative weight cycle, which means
// arbitrage observed.

// Big O: O(n^3) time | O(n^2) space, creating a fully-connected directed graph with n vertices,
// such graph has n^2 edges, thus come O(n^2) space;
// as for time complexity, function relaxEdgesAndUpdateDistances() is Bellman-ford algorithm,
// which computes the shortest path from a source vertex to all of the other vertices in the graph,
// it requires to run n-1 times, within each iteration, there's 2 nested for loop over n vertices,
// so in total O(n^3) time
const detectArbitrage = (exchangeRates: number[][]): boolean => {
  const logExchangeRates = convertToLogMatrix(exchangeRates);

  return foundNegativeWeightCycle(logExchangeRates, 0);
};

const foundNegativeWeightCycle = (
  graph: number[][],
  startVertex: number
): boolean => {
  const distancesFromStart = new Array(graph.length)
    .fill(0)
    .map(() => Number.POSITIVE_INFINITY);
  distancesFromStart[startVertex] = 0;

  for (let i = 0; i < graph.length - 1; i++) {
    if (!relaxEdgesAndUpdateDistances(graph, distancesFromStart)) {
      return false;
    }
  }
  return relaxEdgesAndUpdateDistances(graph, distancesFromStart);
};

const relaxEdgesAndUpdateDistances = (
  graph: number[][],
  distances: number[]
): boolean => {
  let updated = false;
  for (let sourceIdx = 0; sourceIdx < graph.length; sourceIdx++) {
    const edges = graph[sourceIdx];
    for (
      let destinationIdx = 0;
      destinationIdx < edges.length;
      destinationIdx++
    ) {
      const edgeWeight = edges[destinationIdx];
      const newDistanceToDestination = distances[sourceIdx] + edgeWeight;
      if (newDistanceToDestination < distances[destinationIdx]) {
        distances[destinationIdx] = newDistanceToDestination;
        updated = true;
      }
    }
  }
  return updated;
};

// leveraging the fact that log(a*b) = log(a) + log(b)
// so to find a*b*c > 1, we need  -log(a*b*c) < -log(1) => -log(a) + -log(b) + -log(c) < 0
const convertToLogMatrix = (matrix: number[][]): number[][] => {
  const newMatrix: number[][] = [];
  for (let row = 0; row < matrix.length; row++) {
    newMatrix.push([]);
    for (const rate of matrix[row]) {
      newMatrix[row].push(-Math.log(rate));
    }
  }
  return newMatrix;
};

const exchangeRates = [
  [1.0, 0.8631, 0.5903],
  [1.1586, 1.0, 0.6849],
  [1.6939, 1.46, 1.0]
];
console.log(detectArbitrage(exchangeRates));
