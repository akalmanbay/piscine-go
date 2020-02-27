package piscine 

func isPrintable(str string) bool {
	for i, _ := range str {
		if str[i] >= 00 || str[i] <= 127 {
			return true
		}
	}
	return false
}
