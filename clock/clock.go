package clock

import "fmt"

const testVersion = 3

//Clock is a constructor of int type
type Clock int

func New(hour, minute int) Clock {
	//minute can take more than 60 so have to work on it
	time := (hour*60 + minute) % (60*24)
	if time < 0 {
		time += 60*24
	}
	return Clock(time)

}

func (clk Clock) String() string {

	return fmt.Sprintf("%.2d:%.2d", clk/60, clk%60)

}

func (clk Clock) Add(minutes int) Clock {
	return New(0, int(clk)+minutes)

}
