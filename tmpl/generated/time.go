package generated

import "time"

type Time struct{}

func NewTime() *Time         { return &Time{} }
func (_ Time) ANSIC() string { return time.ANSIC }

func (_ *Time) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}

func (_ *Time) AfterFunc(d time.Duration, f func()) *time.Timer {
	return time.AfterFunc(d, f)
}
func (_ Time) April() time.Month  { return time.April }
func (_ Time) August() time.Month { return time.August }

func (_ *Time) Date(year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location) time.Time {
	return time.Date(year, month, day, hour, min, sec, nsec, loc)
}
func (_ Time) December() time.Month { return time.December }
func (_ Time) February() time.Month { return time.February }

func (_ *Time) FixedZone(name string, offset int) *time.Location {
	return time.FixedZone(name, offset)
}
func (_ Time) Friday() time.Weekday { return time.Friday }
func (_ Time) Hour() time.Duration  { return time.Hour }
func (_ Time) January() time.Month  { return time.January }
func (_ Time) July() time.Month     { return time.July }
func (_ Time) June() time.Month     { return time.June }
func (_ Time) Kitchen() string      { return time.Kitchen }

func (_ *Time) LoadLocation(name string) (*time.Location, error) {
	return time.LoadLocation(name)
}

func (_ *Time) LoadLocationFromTZData(name string, data []byte) (*time.Location, error) {
	return time.LoadLocationFromTZData(name, data)
}
func (_ Time) Local() *time.Location      { return time.Local }
func (_ Time) March() time.Month          { return time.March }
func (_ Time) May() time.Month            { return time.May }
func (_ Time) Microsecond() time.Duration { return time.Microsecond }
func (_ Time) Millisecond() time.Duration { return time.Millisecond }
func (_ Time) Minute() time.Duration      { return time.Minute }
func (_ Time) Monday() time.Weekday       { return time.Monday }
func (_ Time) Nanosecond() time.Duration  { return time.Nanosecond }

func (_ *Time) NewTicker(d time.Duration) *time.Ticker {
	return time.NewTicker(d)
}

func (_ *Time) NewTimer(d time.Duration) *time.Timer {
	return time.NewTimer(d)
}
func (_ Time) November() time.Month { return time.November }

func (_ *Time) Now() time.Time {
	return time.Now()
}
func (_ Time) October() time.Month { return time.October }

func (_ *Time) Parse(layout string, value string) (time.Time, error) {
	return time.Parse(layout, value)
}

func (_ *Time) ParseDuration(s string) (time.Duration, error) {
	return time.ParseDuration(s)
}

func (_ *Time) ParseInLocation(layout string, value string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(layout, value, loc)
}
func (_ Time) RFC1123() string        { return time.RFC1123 }
func (_ Time) RFC1123Z() string       { return time.RFC1123Z }
func (_ Time) RFC3339() string        { return time.RFC3339 }
func (_ Time) RFC3339Nano() string    { return time.RFC3339Nano }
func (_ Time) RFC822() string         { return time.RFC822 }
func (_ Time) RFC822Z() string        { return time.RFC822Z }
func (_ Time) RFC850() string         { return time.RFC850 }
func (_ Time) RubyDate() string       { return time.RubyDate }
func (_ Time) Saturday() time.Weekday { return time.Saturday }
func (_ Time) Second() time.Duration  { return time.Second }
func (_ Time) September() time.Month  { return time.September }

func (_ *Time) Since(t time.Time) time.Duration {
	return time.Since(t)
}

func (_ *Time) Sleep(d time.Duration) {
	time.Sleep(d)
}
func (_ Time) Stamp() string          { return time.Stamp }
func (_ Time) StampMicro() string     { return time.StampMicro }
func (_ Time) StampMilli() string     { return time.StampMilli }
func (_ Time) StampNano() string      { return time.StampNano }
func (_ Time) Sunday() time.Weekday   { return time.Sunday }
func (_ Time) Thursday() time.Weekday { return time.Thursday }

func (_ *Time) Tick(d time.Duration) <-chan time.Time {
	return time.Tick(d)
}
func (_ Time) Tuesday() time.Weekday { return time.Tuesday }
func (_ Time) UTC() *time.Location   { return time.UTC }

func (_ *Time) Unix(sec int64, nsec int64) time.Time {
	return time.Unix(sec, nsec)
}
func (_ Time) UnixDate() string { return time.UnixDate }

func (_ *Time) Until(t time.Time) time.Duration {
	return time.Until(t)
}
func (_ Time) Wednesday() time.Weekday { return time.Wednesday }
