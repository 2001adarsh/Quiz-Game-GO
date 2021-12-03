package Play

import "strconv"

func convertToInt(a string) int {
	ans, err := strconv.Atoi(a)
	if err != nil {
		return -1
	}
	return ans
}
