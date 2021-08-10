export default class Stack<T> {
  stack: T[] = [];

  peak = (): T => {
    return this.stack[this.stack.length - 1];
  };

  push = (val: T): void => {
    this.stack.push(val);
  };

  pop = (): void => {
    this.stack.pop();
  };
}
