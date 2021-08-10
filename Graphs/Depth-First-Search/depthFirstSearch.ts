import { GraphNode } from '../GraphNode';

class DfsNode extends GraphNode {
  children: DfsNode[] = [];

  // O(V+E) Time, where V = vertex, E = edge | O(V) Space
  depthFirstSearch(this: DfsNode, array: string[]): string[] {
    array.push(this.name);
    for (const child of this.children) {
      child.depthFirstSearch(array);
    }
    return array;
  }
}

const I = new DfsNode('I');
const J = new DfsNode('J');
const F = new DfsNode('F');
F.addChild(I);
F.addChild(J);
const E = new DfsNode('E');
const G = new DfsNode('G');
const H = new DfsNode('H');
const B = new DfsNode('B');
const C = new DfsNode('C');
const D = new DfsNode('D');
B.addChild(E);
B.addChild(F);
D.addChild(G);
D.addChild(H);
const graph = new DfsNode('A');
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
