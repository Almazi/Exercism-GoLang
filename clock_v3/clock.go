package clock

import "fmt"

const testVersion = 3

//Clock is a constructor of int type
type Clock struct{
	minute, hour int
}

//New is a function that returns Clock - int type data. I can return only one
// value so here I calculated the given hour and minute into minute which I
// used later on to show time and add minute.
func New(hour, minute int) Clock {
	minute = (hour*60 + minute) % 1440
	if minute < 0 {
		minute += 1440 // As the minute is less than 0 suppose -5
		// that means the clock will show the hour before that and
		// minute will show 50-5 = 55. Using 1440 cause later on we
		// find hour by dividing minute by 60. So if the hour is -5
		// and minute is -5 clock has to show 19:55  neans 1140
		// minutes and 55 minutes. thats why have to
		// add 1440 so that later on we can divide it by 1440 and
		// find the proper hour
	}
	hour = minute / 60
	minute = minute % 60
	return Clock{minute, hour}
}

func (clk Clock) String() string {

	return fmt.Sprintf("%02d:%02d", clk.hour, clk.minute)
}

func (clk Clock) Add(minutes int) Clock {
	minutes = clk.minute + minutes
	return New(clk.hour, minutes)
}