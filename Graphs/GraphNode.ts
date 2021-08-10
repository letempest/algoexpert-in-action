export class GraphNode {
  children: GraphNode[] = [];
  constructor(public name: string) {}

  addChild(this: GraphNode, node: GraphNode): void {
    this.children.push(node);
  }
}
