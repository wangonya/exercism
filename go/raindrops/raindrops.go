package raindrops

import "fmt"

func rainDrop(number int, divisor int, sound string) string {
	if number%divisor == 0 {
		return sound
	}
	return ""
}
func Convert(number int) string {
	ans := ""
	ans += rainDrop(number, 3, "Pling")
	ans += rainDrop(number, 5, "Plang")
	ans += rainDrop(number, 7, "Plong")

	if ans == "" {
		ans = fmt.Sprint(number)
	}
	return ans
}
