// Package clause.
package gigasecond

import (
	"time"
)


// Constant declaration.
const testVersion = 4 // find the value in gigasecond_test.go

// API function.  It uses a type from the Go standard library.
const GIGASECOND = time.Duration(1e9) * time.Second

/*AddGigasecond adds one gigasecond (1E9 seconds) to the date.*/
func AddGigasecond(today time.Time) time.Time {
	return today.Add(GIGASECOND)
}

// Birthday is my birthday and is used to compute a giga-versery.
var Birthday, _ = time.Parse("2006-01-02", "1941-12-07")