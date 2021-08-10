import Stack from '../Stack';

// Big O: O(n) time | O(n) space
const isBalanced = (str: string): boolean => {
  const stack = new Stack<string>();
  const matchingBrackets: { [key: string]: string } = {
    ')': '(',
    ']': '[',
    '}': '{'
  };

  for (const char of str) {
    if ('([{'.includes(char)) {
      // opening curly, paren, square bracket
      stack.push(char);
    } else if (')]}'.includes(char)) {
      // closing
      if (stack.stack.length === 0) return false;
      if (stack.peak() !== matchingBrackets[char]) return false;
      stack.pop();
    }
  }
  // after iterating the whole string, check if the stack is empty
  return stack.stack.length === 0;
};

const testStr = '(([]()()){})';

console.log(isBalanced(testStr));
console.log(isBalanced('){}[]()'));
console.log(isBalanced('([)]'));
