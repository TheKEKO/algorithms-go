// https://www.codewars.com/kata/55f9bca8ecaa9eac7100004a/train/go
// Beginner Series #2 Clock

package beginner_series_2_clock

func Past(h, m, s int) int {
	return (h*3600 + m*60 + s*1) * 1000
}
