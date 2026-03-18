// https://www.codewars.com/kata/5a6663e9fd56cb5ab800008b/train/go
package cat_years_dog_years

func CalculateYears(years int) (result [3]int) {
	if years == 1 {
		return [3]int{1, 15, 15}
	} else if years == 2 {
		return [3]int{2, 24, 24}
	} else if years > 2 {
		return [3]int{years, 24 + (years-2)*4, 24 + (years-2)*5}
	}
	return result
}
