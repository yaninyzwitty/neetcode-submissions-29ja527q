/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil // no trees, no levels to process
	}

	// result holds the levels, queue pushes bfs
	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// capture the level and levelsize
		levelSize := len(queue)
		level := []int{}

		for i := 0; i < levelSize; i++ {
			// capture the nodes
			node := queue[0]
			queue = queue[1:]

			// append the level
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
    
}
