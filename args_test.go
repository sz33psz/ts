package ts

import (
	"testing"
	"time"
)

func TestChangeParsing(t *testing.T) {
	testCases := []struct {
		desc    string
		str     string
		correct Change
	}{
		{
			desc:    "only seconds, override",
			str:     "45s",
			correct: newSecondChange(45, true),
		},
		{
			desc:    "only seconds, adjust, negative",
			str:     "-45s",
			correct: newSecondChange(-45, false),
		},
		{
			desc:    "seconds and minutes, override",
			str:     "15m10s",
			correct: newMinuteChange(15, 10, true),
		},
		{
			desc:    "seconds and minutes, adjust, negative",
			str:     "-15m10s",
			correct: newMinuteChange(-15, -10, false),
		},
		{
			desc:    "hours, adjust",
			str:     "+3h",
			correct: newHourChange(3, 0, 0, false),
		},
		{
			desc:    "hours, seconds and minutes, adjust, negative",
			str:     "-1h15m10s",
			correct: newHourChange(-1, -15, -10, false),
		},
		{
			desc:    "months and seconds, adjust",
			str:     "+4M10s",
			correct: newMonthChange(4, 0, 0, 0, 10, false),
		},
		{
			desc:    "years and minutes, override",
			str:     "3y10m",
			correct: newYearChange(3, 0, 0, 0, 10, 0, true),
		},
		{
			desc:    "year, override",
			str:     "1970y",
			correct: newYearChange(1970, 0, 0, 0, 0, 0, true),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			chg, err := NewChange(tC.str)
			if err != nil {
				t.Errorf("Error while parsing %v: %v", tC.str, err)
			}
			compareChanges(t, &tC.correct, chg)
		})
	}
}

func compareChanges(t *testing.T, expected *Change, got *Change) {
	if expected.existFlag != got.existFlag {
		t.Errorf("Invalid exists flag. Expected: %v, got: %v", expected.existFlag, got.existFlag)
	}
	if expected.second != got.second {
		t.Errorf("Invalid second. Expected: %v, got: %v", expected.second, got.second)
	}
	if expected.minute != got.minute {
		t.Errorf("Invalid minute. Expected: %v, got: %v", expected.minute, got.minute)
	}
	if expected.hour != got.hour {
		t.Errorf("Invalid hour. Expected: %v, got: %v", expected.hour, got.hour)
	}
	if expected.day != got.day {
		t.Errorf("Invalid day. Expected: %v, got: %v", expected.day, got.day)
	}
	if expected.month != got.month {
		t.Errorf("Invalid month. Expected: %v, got: %v", expected.month, got.month)
	}
	if expected.year != got.year {
		t.Errorf("Invalid year. Expected: %v, got: %v", expected.year, got.year)
	}
	if expected.override != got.override {
		t.Errorf("Invalid override flag. Expected: %v, got: %v", expected.override, got.override)
	}
}

func TestOverrides(t *testing.T) {
	var timestamp int64 = time.Date(2020, 8, 20, 12, 34, 45, 0, time.UTC).Unix() //2020-08-20 12:34:45
	testCases := []struct {
		desc     string
		override Change
		correct  int64
	}{
		{
			desc:     "Override seconds",
			override: newSecondChange(10, true),
			correct:  time.Date(2020, 8, 20, 12, 34, 10, 0, time.UTC).Unix(),
		},
		{
			desc:     "Override minutes",
			override: newMinuteChange(10, 30, true),
			correct:  time.Date(2020, 8, 20, 12, 10, 30, 0, time.UTC).Unix(),
		},
		{
			desc:     "Override hours",
			override: newHourChange(8, 15, 55, true),
			correct:  time.Date(2020, 8, 20, 8, 15, 55, 0, time.UTC).Unix(),
		},
		{
			desc:     "Override days",
			override: newDayChange(12, 22, 45, 10, true),
			correct:  time.Date(2020, 8, 12, 22, 45, 10, 0, time.UTC).Unix(),
		},
		{
			desc:     "Override months",
			override: newMonthChange(3, 1, 1, 2, 3, true),
			correct:  time.Date(2020, 3, 1, 1, 2, 3, 0, time.UTC).Unix(),
		},
		{
			desc:     "Override years",
			override: newYearChange(2019, 12, 25, 17, 1, 45, true),
			correct:  time.Date(2019, 12, 25, 17, 1, 45, 0, time.UTC).Unix(),
		},
		{
			desc:     "Override years",
			override: newYearChange(1970, 0, 0, 0, 0, 0, true),
			correct:  0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			calculated := tC.override.Apply(timestamp)
			if tC.correct != calculated {
				t.Errorf("Error in %v. Expected %v, got %v", tC.desc, tC.correct, calculated)
			}
		})
	}
}

func TestAdjust(t *testing.T) {
	var timestamp int64 = time.Date(2020, 8, 20, 12, 34, 45, 0, time.UTC).Unix() //2020-08-20 12:34:45
	testCases := []struct {
		desc     string
		override Change
		correct  int64
	}{
		{
			desc:     "Adjust seconds",
			override: newSecondChange(-10, false),
			correct:  time.Date(2020, 8, 20, 12, 34, 35, 0, time.UTC).Unix(),
		},
		{
			desc:     "Adjust minutes",
			override: newMinuteChange(10, 30, false),
			correct:  time.Date(2020, 8, 20, 12, 45, 15, 0, time.UTC).Unix(),
		},
		{
			desc:     "Adjust hours",
			override: newHourChange(-1, -10, -10, false),
			correct:  time.Date(2020, 8, 20, 11, 24, 35, 0, time.UTC).Unix(),
		},
		{
			desc:     "Adjust days",
			override: newDayChange(1, 5, 5, 5, false),
			correct:  time.Date(2020, 8, 21, 17, 39, 50, 0, time.UTC).Unix(),
		},
		{
			desc:     "Adjust months",
			override: newMonthChange(-1, -1, -1, -2, -3, false),
			correct:  time.Date(2020, 7, 19, 11, 32, 42, 0, time.UTC).Unix(),
		},
		{
			desc:     "Adjust years",
			override: newYearChange(-5, 1, 1, 1, 1, 1, false),
			correct:  time.Date(2015, 9, 21, 13, 35, 46, 0, time.UTC).Unix(),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			calculated := tC.override.Apply(timestamp)
			if tC.correct != calculated {
				t.Errorf("Error in %v. Expected %v, got %v", tC.desc, tC.correct, calculated)
			}
		})
	}
}
