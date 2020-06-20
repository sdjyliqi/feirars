package utils

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
