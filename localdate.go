package gocommon

import "time"

type LocalDate time.Time

var day time.Duration = time.Second * 60 * 60 * 24

func NewLocalDate(year int, month time.Month, day int) LocalDate {
	return LocalDate(time.Date(year, month, day, 0, 0, 0, 0, time.UTC))
}

func ParseIsoDate(iso string) (LocalDate, error) {
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
	return time.Time(d).Sub(time.Time(o)) < 0
}

func (d LocalDate) isAfter(o LocalDate) bool {
	return time.Time(d).Sub(time.Time(o)) > 0
}

func (d LocalDate) PlusMonths(months int) LocalDate {
	m1 := d.GetMonth()
	m2 := int(m1) + months
	numYears := m2 / 12
	numMonths := m2 % 12

	if numYears != 0 {
		return d.PlusYears(numYears).PlusMonths(numMonths)
	}
	d2 := NewLocalDate(d.GetYear(), time.Month(m2), d.GetDay())
	if d2.GetMonth() != time.Month(m2) {
		d2 = d2.MinusDays(1)
	}
	return d2
}

func (d LocalDate) MinusMonths(months int) LocalDate {
	return d.PlusMonths(-1 * months)
}

func (d LocalDate) PlusYears(years int) LocalDate {
	return NewLocalDate(d.GetYear()+years, d.GetMonth(), d.GetYear())
}

func (d LocalDate) Equals(o LocalDate) bool {
	return time.Time(d).Sub(time.Time(o)) == 0
}
func (d LocalDate) String() string {
	return time.Time(d).Format(time.DateOnly)
}
