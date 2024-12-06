package localdate

import (
	"encoding/json"
	"time"

	"github.com/beglaryh/gocommon"
)

type LocalDate time.Time

var day time.Duration = time.Second * 60 * 60 * 24

func New(year int, month time.Month, day int) LocalDate {
	return LocalDate(time.Date(year, month, day, 0, 0, 0, 0, time.UTC))
}

func Now() LocalDate {
	t := time.Now()
	return New(t.Year(), t.Month(), t.Day())
}

func Parse(iso string) (LocalDate, error) {
	d, err := time.Parse(time.DateOnly, iso)
	return LocalDate(d), err
}

func (d LocalDate) GetMonth() time.Month {
	return time.Time(d).Month()
}

func (d LocalDate) GetYear() int {
	return time.Time(d).Year()
}

func (d LocalDate) GetDay() int {
	return time.Time(d).Day()
}

func (d LocalDate) PlusDays(days int) LocalDate {
	numDays := time.Duration(days) * day
	return LocalDate(time.Time(d).Add(numDays))
}

func (d LocalDate) MinusDays(days int) LocalDate {
	return d.PlusDays(-1 * days)
}

func (d LocalDate) IsBefore(o LocalDate) bool {
	return time.Time(d).Before(time.Time(o))
}

func (d LocalDate) IsAfter(o LocalDate) bool {
	return time.Time(d).After(time.Time(o))
}

func (d LocalDate) PlusMonths(months int) LocalDate {
	m1 := d.GetMonth()
	m2 := int(m1) + months
	numYears := m2 / 12
	numMonths := m2 % 12

	if numYears != 0 {
		return d.PlusYears(numYears).PlusMonths(numMonths)
	}
	d2 := New(d.GetYear(), time.Month(m2), d.GetDay())
	if d2.GetMonth() != time.Month(m2) {
		d2 = d2.MinusDays(1)
	}
	return d2
}

func (d LocalDate) MinusMonths(months int) LocalDate {
	return d.PlusMonths(-1 * months)
}

func (d LocalDate) PlusYears(years int) LocalDate {
	return New(d.GetYear()+years, d.GetMonth(), d.GetYear())
}

func (d LocalDate) Equals(o LocalDate) bool {
	return time.Time(d).Sub(time.Time(o)) == 0
}

func (d LocalDate) String() string {
	return time.Time(d).Format(time.DateOnly)
}

func (d LocalDate) ToString() gocommon.String {
	return gocommon.String(d.String())
}

func (d LocalDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *LocalDate) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	date, err := Parse(str)
	*d = date
	return err
}
