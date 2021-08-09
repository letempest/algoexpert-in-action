class GraphNode {
  children: GraphNode[] = [];
  constructor(public name: string) {}

  addChild(this: GraphNode, node: GraphNode): void {
    this.children.push(node);
  }

  // O(V+E) Time, where V = vertex, E = edge | O(V) Space
  depthFirstSearch(this: GraphNode, array: string[]): string[] {
    array.push(this.name);
    for (const child of this.children) {
      child.depthFirstSearch(array);
    }
    return array;
  }
}

const I = new GraphNode('I');
const J = new GraphNode('J');
const F = new GraphNode('F');
F.addChild(I);
F.addChild(J);
const E = new GraphNode('E');
const G = new GraphNode('G');
const H = new GraphNode('H');
const B = new GraphNode('B');
const C = new GraphNode('C');
const D = new GraphNode('D');
B.addChild(E);
B.addChild(F);
D.addChild(G);
D.addChild(H);
const graph = new GraphNode('A');
graph.addChild(B);
graph.addChild(C);
graph.addChild(D);

//            A
//          / | \
//         B  C   D
//       /  \    / \
//      E   F   G   H
//         / \
//        I   J

console.log(graph);
// dfs: [A, B, E, F, I, J, C, D, G, H]
console.log(graph.depthFirstSearch([]));
