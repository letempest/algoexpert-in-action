// Given a string with digits, find out all valid ip addresses that could be formed
// by separating the string with dot delimiter
// O(1) Time | O(1) Space; Reason: both time and space of possible ip address has a upper bound O(2^32),
// regardless of what input is provided
const ip = '1921680';

// helper func to check for 2 invalid cases:
// 1. whether the part (substring) is non-zero numbers padded with leading zero
// 2. whether the integer is greater than 255
const isValidPart = (part: string): boolean => {
  const numPart = parseInt(part);
  if (numPart > 255) return false;

  return numPart.toString().length === part.length;
};

const getValidIpList = (ipStr: string): string[] => {
  const ipList: string[] = [];
  let currentIP = '';
  let firstIdx: number = 1,
    secondIdx: number,
    thirdIdx: number;
  let first: string, second: string, third: string, fourth: string;

  if (ipStr.startsWith('0')) return [];

  // while (firstIdx < 4 && firstIdx < ipStr.length - 2) {
  while (firstIdx < Math.min(4, ipStr.length - 2)) {
    first = ipStr.slice(0, firstIdx);
    if (!isValidPart(first)) break;

    secondIdx = firstIdx + 1;
    thirdIdx = secondIdx + 1;
    // while (secondIdx < 7 && secondIdx < ipStr.length - 1) {
    while (secondIdx < Math.min(7, ipStr.length - 1)) {
      second = ipStr.slice(firstIdx, secondIdx);
      if (!isValidPart(second)) break;

      thirdIdx = secondIdx + 1;
      while (thirdIdx < ipStr.length) {
        third = ipStr.slice(secondIdx, thirdIdx);
        fourth = ipStr.slice(thirdIdx);

        if (isValidPart(third) && isValidPart(fourth)) {
          currentIP = [first, second, third, fourth].join('.');
          ipList.push(currentIP);
        }
        thirdIdx++;
      }
      secondIdx++;
    }
    firstIdx++;
  }

  return ipList;
};

console.log(getValidIpList(ip));
console.log(getValidIpList('2125074'));
console.log(getValidIpList('11112'));
console.log(getValidIpList('2553072'));
console.log(getValidIpList('035345'));
