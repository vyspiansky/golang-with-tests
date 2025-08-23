package iteration

import "strings"

// const repeatCount = 5

// func Repeat(character string) string {
// 	var repeated string
// 	for i := 0; i < repeatCount; i++ {
// 		repeated += character
// 	}
// 	return repeated
// }

// func Repeat(character string, repeatCount int) string {
// 	var repeated strings.Builder
// 	for i := 0; i < repeatCount; i++ {
// 		repeated.WriteString(character)
// 	}
// 	return repeated.String()
// }

func Repeat(character string, repeatCount int) string {
	return strings.Repeat(character, repeatCount)
}
