import { GraphNode } from '../GraphNode';

class BfsNode extends GraphNode {
  children: BfsNode[] = [];

  // O(v + e) Time, where v = vertex, e = edge | O(v) Space (worst case when queue has length v-1)
  breadthFirstSearch(this: BfsNode, array: string[]): string[] {
    // simplify the implementation by using an array to act as a LIFO queue
    const queue: BfsNode[] = [this];
    let currentNode: BfsNode;
    while (queue.length > 0) {
      currentNode = queue.shift()!;
      for (const child of currentNode.children) {
        queue.push(child);
      }
      array.push(currentNode.name);
    }
    return array;
  }
}

const I = new BfsNode('I');
const J = new BfsNode('J');
const F = new BfsNode('F');
F.addChild(I);
F.addChild(J);
const E = new BfsNode('E');
const G = new BfsNode('G');
const H = new BfsNode('H');
const B = new BfsNode('B');
const C = new BfsNode('C');
const D = new BfsNode('D');
B.addChild(E);
B.addChild(F);
D.addChild(G);
D.addChild(H);
const graph = new BfsNode('A');
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
// bfs: [A, B, C, D, E, F, G, H, I, J]
console.log(graph.breadthFirstSearch([]));
