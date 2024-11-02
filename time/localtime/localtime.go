package localtime

import (
	"encoding/json"
	"time"

	"github.com/beglaryh/gocommon"
)

type LocalTime time.Time

func New(hour, minute, second, nanosecond int) LocalTime {
	return LocalTime(time.Date(0, 0, 0, hour, minute, second, nanosecond, time.UTC))
}

func Parse(str string) (LocalTime, error) {
	t, err := time.Parse(str, time.TimeOnly)
	return LocalTime(t), err
}

func (t LocalTime) String() string {
	return time.Time(t).Format(time.TimeOnly)
}

func (t LocalTime) ToString() gocommon.String {
	return gocommon.String(time.Time(t).Format(time.TimeOnly))
}

func (d LocalTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *LocalTime) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	date, err := Parse(str)
	*d = date
	return err
}
