// https://www.codewars.com/kata/5a023c426975981341000014/train/go
package main

import "fmt"

func main() {
	fmt.Println(OtherAngle(60, 60))
}

func OtherAngle(a int, b int) int {

	return 180 - a - b
}
