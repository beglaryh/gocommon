package time

import (
	"testing"

	"github.com/beglaryh/gocommon/time/offsetdatetime"
)

func TestOffsetDateTime(t *testing.T) {
	utc := "2024-04-03T06:14:36Z"
	ot, err := offsetdatetime.Parse(utc)
	if err != nil || ot.String() != utc {
		t.Fail()
	}

	nutc := "2007-12-03T10:15:30+01:00"
	ot, err = offsetdatetime.Parse(nutc)
	if err != nil || ot.String() != nutc {
		t.Fail()
	}
}
