package piscine 

func IsNumeric(str string) bool {
for i, _ := range str {
		if str[i] >= 57 || str[i] <= 49 {
			return false
		}
	}
	return true


}
