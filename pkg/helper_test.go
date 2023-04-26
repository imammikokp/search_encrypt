package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMedianRangeValue(t *testing.T){
	t.Run("1-5", func(t *testing.T){
		leftMin,leftMax,rightMin,rightMax:=MedianRangeValue(1,5)
		t.Log(leftMin," ",leftMax," ",)
		require.Equal(t,1, leftMin)
		require.Equal(t,3,leftMax)
		require.Equal(t,4,rightMin )
		require.Equal(t,5,rightMax)
	})
	t.Run("1-10", func(t *testing.T){
		leftMin,leftMax,rightMin,rightMax:=MedianRangeValue(1,10)
		t.Log(leftMin," ",leftMax," ",)
		require.Equal(t,1, leftMin)
		require.Equal(t,5,leftMax)
		require.Equal(t,6,rightMin )
		require.Equal(t,10,rightMax)
	})
	t.Run("1-3", func(t *testing.T){
			leftMin,leftMax,rightMin,rightMax:=MedianRangeValue(1,3)
		t.Log(leftMin," ",leftMax," ",)
		require.Equal(t,1, leftMin)
		require.Equal(t,2,leftMax)
		require.Equal(t,3,rightMin )
		require.Equal(t,3,rightMax)
	})
}