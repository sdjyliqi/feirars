package utils

import (
	"testing"
)

func Test_TwoSliceIntersect(t *testing.T) {
	s1 := []string{"1", "2", "3", "4"}
	s2 := []string{"2", "3", "4", "5"}

	s := TwoSliceIntersect(s1, s2)
	t.Log(s)
}
