package utils

import "fmt"

func ConvertToString(v int) string {
	return fmt.Sprintf("%d", v)
}

func ConvertToPercent(v float64) string {
	return fmt.Sprintf("%.2f%%", v*100)
}

func TwoSliceIntersect(s1, s2 []string) []string {
	dest := make([]string, 0)
	m := map[string]bool{}
	for _, v := range s1 {
		m[v] = true
	}

	for _, v := range s2 {
		if m[v] {
			dest = append(dest, v)
		}
	}
	return dest
}

func SliceUnique(s []string) []string {
	dest := make([]string, 0)
	m := map[string]bool{}
	for _, v := range s {
		m[v] = true
	}

	for k := range m {
		dest = append(dest, k)
	}
	return dest
}
