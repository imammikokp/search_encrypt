package pkg

import (
	"fmt"
	"testing"
)

func TestGenerateUUID(t *testing.T){
	id:=GenerateUId()
	fmt.Println(id)
	t.Log(id)
}