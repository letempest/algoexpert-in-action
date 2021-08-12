// Given a tree where every node can track back to its ancestor (but has no child attribute),
// with any two node in the tree, find out their youngest common ancestor in the tree (farthest away one from the topancestor)

class YCANode {
  constructor(public name: string, public ancestor: YCANode | null = null) {}
}

const getDescendantDepth = (node: YCANode, topAncestor: YCANode): number => {
  let depth = 0;
  while (node !== topAncestor) {
    depth += 1;
    node = node.ancestor!;
  }
  return depth;
};

// Big O: O(d) time where d denotes depth of the lower descendant in the tree | O(1) space
const getYoungestCommonAncestor = (
  topAncestor: YCANode,
  descendantOne: YCANode,
  descendantTwo: YCANode
): YCANode => {
  const depthOne = getDescendantDepth(descendantOne, topAncestor);
  const depthTwo = getDescendantDepth(descendantTwo, topAncestor);
  let depthDiff = depthOne - depthTwo;
  if (depthDiff > 0) {
    return backtrackAncestralTree(descendantOne, descendantTwo, depthDiff);
  } else {
    return backtrackAncestralTree(descendantTwo, descendantOne, -1 * depthDiff);
  }
};

const backtrackAncestralTree = (
  lowerDescendant: YCANode,
  higherDescendant: YCANode,
  diff: number
): YCANode => {
  while (diff > 0) {
    lowerDescendant = lowerDescendant.ancestor!;
    diff--;
  }
  while (lowerDescendant !== higherDescendant) {
    lowerDescendant = lowerDescendant.ancestor!;
    higherDescendant = higherDescendant.ancestor!;
  }
  return higherDescendant;
};

const root = new YCANode('A');
const b = new YCANode('B', root);
const g = new YCANode('G', b);
const h = new YCANode('H', b);
const o = new YCANode('O', h);
const p = new YCANode('P', h);
const t = new YCANode('T', p);
const u = new YCANode('U', p);
const q = new YCANode('Q', h);
const r = new YCANode('R', h);
const v = new YCANode('V', r);
const w = new YCANode('W', v);
const x = new YCANode('X', v);
const z = new YCANode('Z', x);
const y = new YCANode('Y', v);
const i = new YCANode('I', b);
const c = new YCANode('C', root);
const j = new YCANode('J', c);
const d = new YCANode('D', root);
const k = new YCANode('K', d);
const l = new YCANode('L', d);
const e = new YCANode('E', root);
const f = new YCANode('F', root);
const m = new YCANode('M', f);
const n = new YCANode('N', f);

//                                          A
//                      B            C       D       E       F
//                   /  |   \        |      / \             / \
//                G     H     I      J      K  L            M  N
//                  O  P  Q  R                                                (o, p, q, r are all child of h)
//                    /\      \
//                   T  U      V
//                           / | \
//                          W  X  Y
//                             |
//                             Z

console.log(getYoungestCommonAncestor(root, t, i));
console.log(getYoungestCommonAncestor(root, u, v));
