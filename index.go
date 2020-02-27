package main

func Index(s string, toFind string) int {
	lf := 0
	for range toFind {
		lf++
	}

	for i := range s {
		if s[i] == toFind[0] {
			slice := s[i : i+lf]
			if slice == toFind {
				return i
			}
			
		}
	}
	return -1
}
