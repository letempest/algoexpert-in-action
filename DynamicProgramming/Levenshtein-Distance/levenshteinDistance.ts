// Given two strings, e.g. 'abc' and 'yabd', decided the minimum amount of steps required to transform from string 1 to string 2
// actions available: 1)insertion  2)removal  3)substitution
// for the example, result should be 2, says insert 'y' in front and substitute 'c' with 'd' at the end

// Big O: O(nm) time | O(nm) space
// const levenshteinDistance = (str1: string, str2: string): number => {
//   // initialize a n by m matrix, where n, m are length of str1, str2 respectively
//   const edits: number[][] = new Array(str2.length + 1)
//     .fill(0)
//     .map(() => new Array(str1.length + 1).fill(0));

//   edits[0] = edits[0].map((_, i) => i);
//   for (let i = 1; i < str2.length + 1; i++) {
//     edits[i][0] = edits[i - 1][0] + 1;
//   }
//   // i -> row, j -> column
//   for (let i = 1; i < str2.length + 1; i++) {
//     for (let j = 1; j < str1.length + 1; j++) {
//       if (str2[i - 1] === str1[j - 1]) {
//         // equal to diagonal value
//         edits[i][j] = edits[i - 1][j - 1];
//       } else {
//         // equal to the minimum of three neighboring values plus 1
//         edits[i][j] =
//           1 + Math.min(edits[i][j - 1], edits[i - 1][j], edits[i - 1][j - 1]);
//       }
//     }
//   }

//   return edits[str2.length][str1.length];
// };

// Improved Space Complexity Solution
// O(nm) time | O(min(n, m)) space (no need to maintain a n by m matrix, at any point of time only a 2 x min(n, m) matrix is necessary)
const levenshteinDistance = (str1: string, str2: string): number => {
  const short = str1.length < str2.length ? str1 : str2;
  const long = str1.length >= str2.length ? str1 : str2;
  // initialize two min(n, m) length array, where n, m are length of str1, str2 respectively
  const evenEdits = new Array(short.length + 1).fill(0).map((_, i) => i);
  const oddEdits = new Array(short.length + 1).fill(0);

  let currentEdits: number[] = oddEdits,
    previousEdits: number[] = evenEdits;
  for (let i = 1; i < long.length + 1; i++) {
    if (i % 2 === 1) {
      currentEdits = oddEdits;
      previousEdits = evenEdits;
    } else {
      currentEdits = evenEdits;
      previousEdits = oddEdits;
    }
    currentEdits[0] = i;
    for (let j = 1; j < short.length + 1; j++) {
      if (long[i - 1] === short[j - 1]) {
        currentEdits[j] = previousEdits[j - 1];
      } else {
        currentEdits[j] =
          1 +
          Math.min(currentEdits[j - 1], previousEdits[j], previousEdits[j - 1]);
      }
    }
  }
  return currentEdits[currentEdits.length - 1];
};

console.log(levenshteinDistance('abc', 'yabd'));
