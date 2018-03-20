package generated

import "time"

type Time struct{}

func NewTime() *Time       { return &Time{} }
func (Time) ANSIC() string { return time.ANSIC }

func (*Time) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}

func (*Time) AfterFunc(d time.Duration, f func()) *time.Timer {
	return time.AfterFunc(d, f)
}
func (Time) April() time.Month  { return time.April }
func (Time) August() time.Month { return time.August }

func (*Time) Date(year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location) time.Time {
	return time.Date(year, month, day, hour, min, sec, nsec, loc)
}
func (Time) December() time.Month { return time.December }
func (Time) February() time.Month { return time.February }

func (*Time) FixedZone(name string, offset int) *time.Location {
	return time.FixedZone(name, offset)
}
func (Time) Friday() time.Weekday { return time.Friday }
func (Time) Hour() time.Duration  { return time.Hour }
func (Time) January() time.Month  { return time.January }
func (Time) July() time.Month     { return time.July }
func (Time) June() time.Month     { return time.June }
func (Time) Kitchen() string      { return time.Kitchen }

func (*Time) LoadLocation(name string) (*time.Location, error) {
	return time.LoadLocation(name)
}

func (*Time) LoadLocationFromTZData(name string, data []byte) (*time.Location, error) {
	return time.LoadLocationFromTZData(name, data)
}
func (Time) Local() *time.Location      { return time.Local }
func (Time) March() time.Month          { return time.March }
func (Time) May() time.Month            { return time.May }
func (Time) Microsecond() time.Duration { return time.Microsecond }
func (Time) Millisecond() time.Duration { return time.Millisecond }
func (Time) Minute() time.Duration      { return time.Minute }
func (Time) Monday() time.Weekday       { return time.Monday }
func (Time) Nanosecond() time.Duration  { return time.Nanosecond }

func (*Time) NewTicker(d time.Duration) *time.Ticker {
	return time.NewTicker(d)
}

func (*Time) NewTimer(d time.Duration) *time.Timer {
	return time.NewTimer(d)
}
func (Time) November() time.Month { return time.November }

func (*Time) Now() time.Time {
	return time.Now()
}
func (Time) October() time.Month { return time.October }

func (*Time) Parse(layout string, value string) (time.Time, error) {
	return time.Parse(layout, value)
}

func (*Time) ParseDuration(s string) (time.Duration, error) {
	return time.ParseDuration(s)
}

func (*Time) ParseInLocation(layout string, value string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(layout, value, loc)
}
func (Time) RFC1123() string        { return time.RFC1123 }
func (Time) RFC1123Z() string       { return time.RFC1123Z }
func (Time) RFC3339() string        { return time.RFC3339 }
func (Time) RFC3339Nano() string    { return time.RFC3339Nano }
func (Time) RFC822() string         { return time.RFC822 }
func (Time) RFC822Z() string        { return time.RFC822Z }
func (Time) RFC850() string         { return time.RFC850 }
func (Time) RubyDate() string       { return time.RubyDate }
func (Time) Saturday() time.Weekday { return time.Saturday }
func (Time) Second() time.Duration  { return time.Second }
func (Time) September() time.Month  { return time.September }

func (*Time) Since(t time.Time) time.Duration {
	return time.Since(t)
}

func (*Time) Sleep(d time.Duration) {
	time.Sleep(d)
}
func (Time) Stamp() string          { return time.Stamp }
func (Time) StampMicro() string     { return time.StampMicro }
func (Time) StampMilli() string     { return time.StampMilli }
func (Time) StampNano() string      { return time.StampNano }
func (Time) Sunday() time.Weekday   { return time.Sunday }
func (Time) Thursday() time.Weekday { return time.Thursday }

func (*Time) Tick(d time.Duration) <-chan time.Time {
	return time.Tick(d)
}
func (Time) Tuesday() time.Weekday { return time.Tuesday }
func (Time) UTC() *time.Location   { return time.UTC }

func (*Time) Unix(sec int64, nsec int64) time.Time {
	return time.Unix(sec, nsec)
}
func (Time) UnixDate() string { return time.UnixDate }

func (*Time) Until(t time.Time) time.Duration {
	return time.Until(t)
}
func (Time) Wednesday() time.Weekday { return time.Wednesday }
