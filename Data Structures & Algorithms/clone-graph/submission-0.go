/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    
    clone := make(map[*Node]*Node)
    var dfs func(cur *Node) *Node
    dfs = func(cur *Node) *Node {
        if cNode, exist := clone[cur]; exist {
            return cNode
        }

        copyNode := &Node{
            Val: cur.Val,
            Neighbors: make([]*Node, 0, len(cur.Neighbors)),
        }

        clone[cur] = copyNode

        for _, neighbor := range cur.Neighbors {
            copyNeighbors := dfs(neighbor)
            copyNode.Neighbors = append(copyNode.Neighbors, copyNeighbors)
        }

        return copyNode
    }

    return dfs(node)
}
