package clock

import "fmt"

const testVersion = 3

//Clock is a constructor of int type
type Clock int

//New is a function that returns Clock - int type data. I can return only one
// value so here I calculated the given hour and minute into minute which I
// used later on to show time and add minute.
func New(hour, minute int) Clock {
	//minute can take more than 60 so have to work on it
	time := hour*60 + minute
	if time > 1439 {
		time = time % 1440
		return Clock(time)
	} else if time < 0 {
		time = time%1440 + 60*24
		return Clock(time)
	} else {
		return Clock(time)
	}

}

func (clk Clock) String() string {

	//return fmt.Sprintf("%.2d:%.2d", clk / 60, clk % 60)
	return fmt.Printf("%.2d:%.2d", clk/60, clk%60)

}

func (clk Clock) Add(minutes int) Clock {
	return New(0, int(clk)+minutes)

}
