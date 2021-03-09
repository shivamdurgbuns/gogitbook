package iteration

const repeatCount = 5

/*
Repeat function will repeat the input n number of times withour space.
*/
func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
