/**
 * Definition for undirected graph.
 * function UndirectedGraphNode(label) {
 *     this.label = label;
 *     this.neighbors = [];   // Array of UndirectedGraphNode
 * }
 */

/**
 * @param {UndirectedGraphNode} graph
 * @return {UndirectedGraphNode}
 */
var cloneGraph = function(node) {
    const map = {};
    return cloneDeep(node);
  
    function cloneDeep(node) {
      if (!node) return node;
      if (!map[node.label]) {
        map[node.label] = new UndirectedGraphNode(node.label);
        map[node.label].neighbors = node.neighbors.map(cloneDeep);
      }
      return map[node.label];
    }
  };