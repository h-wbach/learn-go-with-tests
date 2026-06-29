package iteration

const repeatTotal = 5

func Repeat(letter string) string {
	var repeated string
	for i := 0; i < repeatTotal; i++ {
		repeated += letter
	}
	return repeated
}
