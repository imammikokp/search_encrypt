package domain

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	root := NewRoot(0, 500)

	t.Run("divide", func(t *testing.T) {
		DevidedEqually(root.root)
	})

	t.Run("NearestSearc", func(t *testing.T) {
		fmt.Println("\n  ------- \n NearrestSearch")
		parent,node := NearestSearch(nil,root.root)
		Stringify(root.root, 0)
		fmt.Println("parent ",parent.Min," ",parent.Max,"node",node.Min, " ",node.Max)
		// Remove(node)
		fmt.Println(node.Min, node.Max)
		Stringify(root.root, 0)
	})

	
	t.Run("SearchEqualByMinMax", func(t *testing.T){
		fmt.Println("\n ----- \n search Equal By min max" )
		parent,node := SearchEqualByMinMax(nil, root.root,251,500 )
		fmt.Println("---graph tree---")
		if parent != nil && node != nil{
			fmt.Println("parent ",parent.Min," ",parent.Max,"node",node.Min, " ",node.Max)
			fmt.Println(node.Min, node.Max)
		}
		Stringify(root.root, 0)
	})

	t.Run("Remove by min max", func(t *testing.T){
			fmt.Println("\n ----- \n search Equal By min max" )
			result:=RemoveEqualByMinMax(nil,root.root,251,500)
			fmt.Println("---graph tree---")
			fmt.Println(result)
			Stringify(root.root, 0)
	})

	// t.Run("")
}
