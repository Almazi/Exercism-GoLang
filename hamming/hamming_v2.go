package hamming

import "errors"

const testVersion = 4

//Distance will take two string and see the lengths are equal or not. If not
// equal it will return count = -1 as there is no match and also an error. I
// am not quite sure what error is but I read abit from: https://gobyexample
// .com/errors. then I am checking every position/index of both string arrays
// and if there is any inequality i see i count 1 more than before. at the
// end i send total count of inequalities and a nil error
func Distance(a, b string) (int, error) {
	count := 0
	if len(a) != len(b) {
		return -1, errors.New("Inequal distance")
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				count++
			}
		}
	}

	return count, nil
}
