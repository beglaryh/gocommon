package offsetdatetime

import (
	"encoding/json"
	"time"

	"github.com/beglaryh/gocommon"
)

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

func (t OffsetDateTime) ToString() gocommon.String {
	return gocommon.String(t.String())
}

func (t OffsetDateTime) IsBefore(other OffsetDateTime) bool {
	return time.Time(t).Before(time.Time(other))
}

func (t OffsetDateTime) IsAfter(other OffsetDateTime) bool {
	return time.Time(t).After(time.Time(other))
}

func (d OffsetDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *OffsetDateTime) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	date, err := Parse(str)
	*d = date
	return err
}
