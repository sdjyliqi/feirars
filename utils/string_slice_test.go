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

func Test_SliceUnique(t *testing.T) {
	s := []string{"1", "2", "3", "4", "1"}
	dest := SliceUnique(s)
	t.Log(dest)
}

func Test_ConvertToPercent(t *testing.T) {
	a, b := int(5), int(12)

	t.Log(float64(a) / float64(b))

	v := float64(0.1234)
	res := ConvertToPercent(v)
	t.Log(res)
}
