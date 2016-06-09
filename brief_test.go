package bulletin

import (
	"testing"
	"time"
)

func TestKST(t *testing.T) {
	UTC := time.Now().UTC()
	KST := time.Date(
		UTC.Year(), UTC.Month(), UTC.Day(),
		UTC.Hour(), UTC.Minute(), UTC.Second(),
		UTC.Nanosecond(), adequateTimezone,
	)

	if UTC.Sub(KST) != 9*time.Hour {
		t.Error("KST is not 9 hours ahead of UTC")
		t.Fail()
	}
}

func TestLoadKST(t *testing.T) {
	_, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Error("cannot find Asia/Tokyo timezone in zoneinfo.zip")
		t.Fail()
	}
}
