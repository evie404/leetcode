package allintwobst

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	all1 := inOrder(root1)
	all2 := inOrder(root2)

	return mergeSort(all1, all2)
}

func mergeSort(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))

	aIdx := 0
	bIdx := 0

	for aIdx < len(a) && bIdx < len(b) {
		if a[aIdx] < b[bIdx] {
			result = append(result, a[aIdx])
			aIdx++
		} else {
			result = append(result, b[bIdx])
			bIdx++
		}
	}

	result = append(result, a[aIdx:]...)
	result = append(result, b[bIdx:]...)

	return result
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	results := []int{}

	results = append(results, inOrder(root.Left)...)
	results = append(results, root.Val)
	results = append(results, inOrder(root.Right)...)

	return results
}
