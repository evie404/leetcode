package week3

func levelOrder(root *TreeNode) [][]int {
	values := [][]int{}

	if root == nil {
		return values
	}

	currentLevelNodes := []*TreeNode{root}

	for len(currentLevelNodes) > 0 {
		nextLevelNodes := make([]*TreeNode, 0, len(currentLevelNodes)*2)
		currentLevelValues := make([]int, 0, len(currentLevelNodes))

		for _, node := range currentLevelNodes {
			currentLevelValues = append(currentLevelValues, node.Val)
			if node.Left != nil {
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}

			if node.Right != nil {
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}

		values = append(values, currentLevelValues)
		currentLevelNodes = nextLevelNodes
	}

	return values
}
