package trees

type NodeAVL struct {
	left   *NodeAVL
	right  *NodeAVL
	value  int
	height int
}

func GetHeightAVL(root *NodeAVL) int {
	if root != nil {
		return root.height
	}

	return 0
}

func GetMaxHeightAVL(root *NodeAVL) int {
	return max(GetMaxHeightAVL(root.left), GetMaxHeightAVL(root.right)) + 1
}

func GetBalanceAVL(root *NodeAVL) int {
	if root != nil {
		return GetHeightAVL(root.right) - GetHeightAVL(root.left)
	}

	return 0
}

func LeftRotateAVL(root *NodeAVL) *NodeAVL {
	nodeCopy := root.right
	root.right = nodeCopy.left
	nodeCopy.left = root
	root.height = GetMaxHeightAVL(root)
	nodeCopy.height = GetMaxHeightAVL(nodeCopy)

	return nodeCopy
}

func RightRotateAVL(root *NodeAVL) *NodeAVL {
	nodeCopy := root.left
	root.left = nodeCopy.right
	nodeCopy.right = root
	root.height = GetMaxHeightAVL(root)
	nodeCopy.height = GetMaxHeightAVL(root)

	return nodeCopy
}

func BalanceAVL(root *NodeAVL) *NodeAVL {
	balance := GetBalanceAVL(root)
	if balance == -2 {
		if GetBalanceAVL(root.left) == 1 {
			root.left = LeftRotateAVL(root.left)
		}

		return RightRotateAVL(root)
	} else if balance == 2 {
		if GetBalanceAVL(root.right) == -1 {
			root.right = RightRotateAVL(root.right)
		}

		return LeftRotateAVL(root)
	}

	return root
}

func InsertAVL(root *NodeAVL, value int) *NodeAVL {
	if root == nil {
		return &NodeAVL{nil, nil, value, 1}
	}

	if root.value > value {
		root.left = InsertAVL(root.left, value)
	} else if root.value < value {
		root.right = InsertAVL(root.right, value)
	} else {
		return root
	}

	root.height = GetMaxHeightAVL(root)

	return BalanceAVL(root)
}

func GetMaxNodeAVL(root *NodeAVL) *NodeAVL {
	for root.right != nil {
		root = root.right
	}

	return root
}

func DeleteAVL(root *NodeAVL, value int) *NodeAVL {
	if root == nil {
		return root
	}

	if root.value > value {
		root.left = DeleteAVL(root.left, value)
	} else if root.value < value {
		root.right = DeleteAVL(root.right, value)
	} else {
		if root.left == nil || root.right == nil {
			var node *NodeAVL
			if root.left != nil {
				node = root.left
			} else {
				node = root.right
			}

			if node == nil {
				return nil
			}

			*root = *node
		} else {
			node := GetMaxNodeAVL(root.left)
			root.value = node.value
			root.left = DeleteAVL(root.left, node.value)
		}
	}

	root.height = GetMaxHeightAVL(root)

	return BalanceAVL(root)
}
