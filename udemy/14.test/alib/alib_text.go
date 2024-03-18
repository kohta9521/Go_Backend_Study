package alib

import "testing"

var Debug bool = false

func TestAverage(t *testing.T) {
	if Debug {
		t.Skip("スキップさせる")
	}
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Errorf("%vが返されました", v)
	}
}