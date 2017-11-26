package generated

import "time"

// Time ...
type Time struct {
	After           func(d time.Duration) <-chan time.Time
	Sleep           func(d time.Duration)
	Tick            func(d time.Duration) <-chan time.Time
	ParseDuration   func(s string) (time.Duration, error)
	Since           func(t time.Time) time.Duration
	Until           func(t time.Time) time.Duration
	FixedZone       func(name string, offset int) *time.Location
	LoadLocation    func(name string) (*time.Location, error)
	NewTicker       func(d time.Duration) *time.Ticker
	Date            func(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time
	Now             func() time.Time
	Parse           func(layout, value string) (time.Time, error)
	ParseInLocation func(layout, value string, loc *time.Location) (time.Time, error)
	Unix            func(sec int64, nsec int64) time.Time
	AfterFunc       func(d time.Duration, f func()) *time.Timer
	NewTimer        func(d time.Duration) *time.Timer
}

// NewTime ...
func NewTime() *Time {
	return &Time{
		After:           time.After,
		Sleep:           time.Sleep,
		Tick:            time.Tick,
		ParseDuration:   time.ParseDuration,
		Since:           time.Since,
		Until:           time.Until,
		FixedZone:       time.FixedZone,
		LoadLocation:    time.LoadLocation,
		NewTicker:       time.NewTicker,
		Date:            time.Date,
		Now:             time.Now,
		Parse:           time.Parse,
		ParseInLocation: time.ParseInLocation,
		Unix:            time.Unix,
		AfterFunc:       time.AfterFunc,
		NewTimer:        time.NewTimer,
	}
}
