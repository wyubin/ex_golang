package easyother

type TreeLevel struct {
	Level int
	Node  *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	arrTemp := []*TreeLevel{{Level: 1, Node: root}}
	arrLen := len(arrTemp)
	var treeLevelCurr *TreeLevel
	levelFinal := 1
	for arrLen > 0 {
		// pop arrTemp
		arrTemp, treeLevelCurr = arrTemp[0:arrLen-1], arrTemp[arrLen-1]
		currNode, level := treeLevelCurr.Node, treeLevelCurr.Level
		// append left and right
		if currNode.Left != nil {
			arrTemp = append(arrTemp, &TreeLevel{Node: currNode.Left, Level: level + 1})
		}
		if currNode.Right != nil {
			arrTemp = append(arrTemp, &TreeLevel{Node: currNode.Right, Level: level + 1})
		}
		// prepare next
		arrLen = len(arrTemp)
		if level > levelFinal {
			levelFinal = level
		}
	}
	return levelFinal
}
