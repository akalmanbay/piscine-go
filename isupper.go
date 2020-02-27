package piscine

func IsUpper(str string) bool {

	for i, _ := range str {
		if str[i] > 90 || str[i] < 65 {
			return false
		}
	}
	return true
}
