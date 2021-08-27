package utils 

import(
	"testing"
)

func TestGetMaxIdx(t *testing.T){
	var slice = [] int{1, 2, 3}
	if 2 != GetMaxIdx(slice) { t.Fatal("faied test") }

	slice = [] int{1}
	if 0 != GetMaxIdx(slice) { t.Fatal("faied test") }
}

func TestGetMinIdx(t *testing.T){
	var slice = [] int{1, 2, 3}
	if 0 != GetMinIdx(slice) { t.Fatal("faied test") }

	slice = [] int{3, 2, 1}
	if 2 != GetMinIdx(slice) { t.Fatal("faied test") }
}