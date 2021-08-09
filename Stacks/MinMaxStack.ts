// Big O: O(1) Time | O(1) Space for all methods (peak, push, pop, getMin, getMax)
class MinMaxStack {
  private stack: number[] = [];
  private minMaxStack: { min: number; max: number }[] = [];

  peak(): number {
    return this.stack[this.stack.length - 1];
  }

  push(val: number): void {
    this.stack.push(val);
    if (this.minMaxStack.length === 0) {
      this.minMaxStack.push({ min: val, max: val });
      return;
    }
    const lastMinMax = this.minMaxStack[this.minMaxStack.length - 1];
    const min = Math.min(lastMinMax.min, val);
    const max = Math.max(lastMinMax.max, val);
    this.minMaxStack.push({ min, max });
  }

  pop(): number | void {
    this.minMaxStack.pop();
    return this.stack.pop();
  }

  getMin = (): number | void => {
    if (this.stack.length === 0) return;
    return this.minMaxStack[this.minMaxStack.length - 1].min;
  };

  getMax(): number | void {
    if (this.stack.length === 0) return;
    return this.minMaxStack[this.minMaxStack.length - 1].max;
  }
}
