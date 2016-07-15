package leap

const testVersion = 2

// IsLeapYear is finding leap year from input integer by following the
// condition: it is a leap year if it is divisible by 4 and either divisible by 400 or not divisible by 100 //
func IsLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}
