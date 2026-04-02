// https://www.codewars.com/kata/57f222ce69e09c3630000212/train/go
// Well of Ideas - Easy Version
package well_of_ideas_easy_version

func Well(x []string) string {
	good := 0

	for _, i := range x {
		if i == "good" {
			good++
		}
	}

	if good == 0 {
		return "Fail!"
	} else if good <= 2 {
		return "Publish!"
	}

	return "I smell a series!"
}
