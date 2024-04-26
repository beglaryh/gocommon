package time

import "time"

type OffsetDateTime time.Time

func CurrentOffsetDateTime() OffsetDateTime {
	return OffsetDateTime(time.Now())
}

func ParseOffsetDateTime(s string) (OffsetDateTime, error) {
	t, err := time.Parse(time.RFC3339, s)
	return OffsetDateTime(t), err
}

func (t OffsetDateTime) String() string {
	return time.Time(t).Format(time.RFC3339)
}
