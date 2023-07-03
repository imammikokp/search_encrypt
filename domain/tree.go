package domain

import (
	"fmt"
	"math"
)

type Root struct {
	Root *Node
}

type Node struct {
	Min   int
	Max   int
	Left  *Node
	Right *Node
}

func NewRoot(min, max int) *Root {
	return &Root{
		Root: &Node{
			Min: min,
			Max: max,
		},
	}
}

// membagi dua value min max
func (t *Node) DevidedEqually() {
	MedianValue := math.RoundToEven(float64(t.Max-t.Min) / 2)
	leftMin := t.Min
	leftMax := int(MedianValue) + t.Min
	rightMin := int(MedianValue) + t.Min + 1
	rightMax := t.Max

	t.Left = &Node{
		Min: leftMin,
		Max: leftMax,
	}

	t.Right = &Node{
		Min: rightMin,
		Max: rightMax,
	}
}

func DevidedEqually(t *Node) {
	MedianValue := math.RoundToEven(float64(t.Max-t.Min) / 2)
	leftMin := t.Min
	leftMax := int(MedianValue) + t.Min
	rightMin := int(MedianValue) + t.Min + 1
	rightMax := t.Max

	t.Left = &Node{
		Min: leftMin,
		Max: leftMax,
	}

	t.Right = &Node{
		Min: rightMin,
		Max: rightMax,
	}
}

func RemoveEqualByMinMax(parentN *Node, childN *Node, min int, max int) bool {

	if childN.Left != nil {
		if min >= childN.Left.Min && max <= childN.Left.Max {
			return RemoveEqualByMinMax(childN, childN.Left, min, max)
		}
	}

	if childN.Right != nil {
		if max >= childN.Right.Min && max <= childN.Right.Max {
			return RemoveEqualByMinMax(childN, childN.Right, min, max)
		}
	}
	if parentN != nil {

		if parentN.Left != nil {
			if parentN.Left.Min == min && parentN.Left.Max == max {
				parentN.Left = nil
				return true
			}
		}

		if parentN.Right != nil {
			if parentN.Right.Min == min && parentN.Right.Max == max {
				parentN.Right = nil
				return true
			}
		}
	}

	return false

}

func SearchEqualByMinMax(parentN *Node, childN *Node, min int, max int) (parent, child *Node) {
	// if
	if childN.Min == min && childN.Max == max {
		parent = parentN
		child = childN
		return
	}

	if childN.Left != nil {
		if min >= childN.Left.Min && max <= childN.Left.Max {
			return SearchEqualByMinMax(childN, childN.Left, min, max)
		}
	}

	if childN.Right != nil {
		if max >= childN.Right.Min && max <= childN.Right.Max {
			return SearchEqualByMinMax(childN, childN.Right, min, max)
		}
	}

	return
}

// search node terkiri dan mengembalikan parrent dan node sekarang
func NearestSearch(parentN, childN *Node) (parent, child *Node) {
	if childN == nil {
		return
	}

	if childN.Left == nil && childN.Right == nil {
		parent = parentN
		child = childN
		return
	}

	if childN.Left != nil {
		return NearestSearch(childN, childN.Left)
	}

	if childN.Right != nil {
		return NearestSearch(childN, childN.Right)
	}
	return
}

func Stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		Stringify(n.Left, level)
		fmt.Printf(format+" min: %d max: %d \n", n.Min, n.Max)
		Stringify(n.Right, level)
	}
}
