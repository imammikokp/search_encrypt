package domain

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	root := NewRoot(0, 500)

	t.Run("divide", func(t *testing.T) {
		DevidedEqually(root.Root)
	})

	t.Run("NearestSearc", func(t *testing.T) {
		fmt.Println("\n  ------- \n NearrestSearch")
		parent,node := NearestSearch(nil,root.Root)
		Stringify(root.Root, 0)
		fmt.Println("parent ",parent.Min," ",parent.Max,"node",node.Min, " ",node.Max)
		// Remove(node)
		fmt.Println(node.Min, node.Max)
		Stringify(root.Root, 0)
	})

	
	t.Run("SearchEqualByMinMax", func(t *testing.T){
		fmt.Println("\n ----- \n search Equal By min max" )
		parent,node := SearchEqualByMinMax(nil, root.Root,251,500 )
		fmt.Println("---graph tree---")
		if parent != nil && node != nil{
			fmt.Println("parent ",parent.Min," ",parent.Max,"node",node.Min, " ",node.Max)
			fmt.Println(node.Min, node.Max)
		}
		Stringify(root.Root, 0)
	})

	t.Run("Remove by min max", func(t *testing.T){
			fmt.Println("\n ----- \n search Equal By min max" )
			result:=RemoveEqualByMinMax(nil,root.Root,251,500)
			fmt.Println("---graph tree---")
			fmt.Println(result)
			Stringify(root.Root, 0)
	})

	// t.Run("")
}
