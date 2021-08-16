// Given a string and a POSITIVE integer key, shift every character forward key time, e.g. 'abc' w/ key=3  => 'def'
// limitation: only consider lower case alphabetical string from a to z now, with a huge integer, the string should wrap back and stay between a-z

// Solution 1: with charcode, a: 97, b: 98, ... , y: 121, z: 122
// Big O: O(n) time | O(n) space
const caesarCipherEncryptor = (srcStr: string, key: number): string => {
  const newKey = key % 26;
  let newChars: string[] = [];
  for (const char of srcStr) {
    const newChar = getNewChar(char, newKey);
    newChars.push(newChar);
  }
  return newChars.join('');
};

const getNewChar = (char: string, key: number): string => {
  let newCharCode = char.charCodeAt(0) + key;
  if (newCharCode > 122) {
    newCharCode = 96 + (newCharCode % 122);
  }
  return String.fromCharCode(newCharCode);
};

// Solution 2, using index range 0-25 instead of charcode
// Big O: O(n) time | O(n) space
// const caesarCipherEncryptor = (srcStr: string, key: number): string => {
//   const newKey = key % 26;
//   let newChars: string[] = [];
//   const alphabet = 'abcdefghijklmnopqrstuvwxyz'.split('');
//   for (const char of srcStr) {
//     const newChar = getNewChar(char, newKey, alphabet);
//     newChars.push(newChar);
//   }
//   return newChars.join('');
// };

// const getNewChar = (char: string, key: number, alphabet: string[]): string => {
//   let newCharIndex = alphabet.indexOf(char) + key;
//   if (newCharIndex > 25) {
//     newCharIndex = -1 + (newCharIndex % 25);
//   }
//   return alphabet[newCharIndex];
// };

console.log(caesarCipherEncryptor('xyz', 2));
