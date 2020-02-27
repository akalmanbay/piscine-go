package piscine

func IsUpper(str string) bool {

	for i, _ := range str {
		if str[i] > 122 || str[i] < 97 {
			return false
		}
	}
	return true
}
