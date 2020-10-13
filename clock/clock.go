package clock

import "fmt"

const minutesInDay = 60 * 24

type Clock int

func New(hour, minute int) Clock {
	var c Clock
	return c.Add(hour * 60).Add(minute)
}

func (c Clock) String() string {
	i := int(c)
	hour := i / 60
	minute := i % 60
	return fmt.Sprintf("%02d:%02d", hour, minute)
}

func (c Clock) Add(minutes int) Clock {
	m := minutes % minutesInDay
	return Clock((int(c) + minutesInDay + m) % minutesInDay)
}

func (c Clock) Subtract(minutes int) Clock {
	m := minutes % minutesInDay
	return Clock((int(c) + minutesInDay - m) % minutesInDay)
}
