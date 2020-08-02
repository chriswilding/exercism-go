package gigasecond

import "time"

const g = time.Duration(1_000_000_000 * time.Second)

// AddGigasecond returns the time 1,000,000,000 after a given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(g)
}
