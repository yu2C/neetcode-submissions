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

    clone := map[*Node]*Node{
        node: {
            Val: node.Val,
            Neighbors: make([]*Node, 0, len(node.Neighbors)),
        },
    }

    queue := []*Node{node}

    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]

        for _, neighbor := range cur.Neighbors {
            if _, exist := clone[neighbor]; !exist {
                clone[neighbor] = &Node{
                    Val: neighbor.Val,
                    Neighbors: make([]*Node, 0, len(neighbor.Neighbors)),
                }
                queue = append(queue, neighbor)
            }

            clone[cur].Neighbors = append(
                clone[cur].Neighbors,
                clone[neighbor],
            )
        }
    }
    return clone[node]
}
