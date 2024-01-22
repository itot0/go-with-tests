package iteration

const (
	repeatCount = 5
)

func Repeat(character string, repeatCount int) string {
	var repeated string
	// initial/condition/after for loop
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
