package offsetdatetime

import "time"

type OffsetDateTime time.Time

func New(year int, month time.Month, day int, hour int, minute int, second int, nanosecond int, location *time.Location) OffsetDateTime {
	return OffsetDateTime(time.Date(year, month, day, hour, minute, second, nanosecond, location))
}

func Now() OffsetDateTime {
	return OffsetDateTime(time.Now())
}

func Parse(s string) (OffsetDateTime, error) {
	t, err := time.Parse(time.RFC3339, s)
	return OffsetDateTime(t), err
}

func (t OffsetDateTime) String() string {
	return time.Time(t).Format(time.RFC3339)
}
