package ts

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/nleeper/goment"
)

const (
	Second = iota
	Minute
	Hour
	Day
	Month
	Year
)

const (
	ModPlus  = '+'
	ModMinus = '-'
)

const TimeUnits = "smhdMy"

var (
	ErrNotChange    = errors.New("not a time change")
	ErrChangeSyntax = errors.New("invalid time change syntax")
)

type change struct {
	year      int
	month     int
	day       int
	hour      int
	minute    int
	second    int
	existFlag uint
	override  bool
}

func (chg *change) apply(t int64) int64 {
	if chg.existFlag == 0 {
		return t
	}
	res, err := goment.Unix(t)
	res = res.SetUTCOffset(0)
	if err != nil {
		//shouldn't happen. Ever.
		return t
	}

	if (chg.existFlag>>Second)&1 == 1 {
		if chg.override {
			res = res.SetSecond(chg.second)
		} else {
			res = res.Add(chg.second, "seconds")
		}
	}

	if (chg.existFlag>>Minute)&1 == 1 {
		if chg.override {
			res = res.SetMinute(chg.minute)
		} else {
			res = res.Add(chg.minute, "minutes")
		}
	}

	if (chg.existFlag>>Hour)&1 == 1 {
		if chg.override {
			res = res.SetHour(chg.hour)
		} else {
			res = res.Add(chg.hour, "hours")
		}
	}

	if (chg.existFlag>>Day)&1 == 1 {
		if chg.override {
			res = res.SetDate(chg.day)
		} else {
			res = res.Add(chg.day, "days")
		}
	}

	if (chg.existFlag>>Month)&1 == 1 {
		if chg.override {
			res = res.SetMonth(chg.month)
		} else {
			res = res.Add(chg.month, "months")
		}
	}

	if (chg.existFlag>>Year)&1 == 1 {
		if chg.override {
			res = res.SetYear(chg.year)
		} else {
			res = res.Add(chg.year, "years")
		}
	}

	return res.ToUnix()
}

func NewChange(s string) (*change, error) {
	if len(s) < 2 { //one for number, one for unit
		return nil, ErrNotChange
	}

	isOverride := true
	plus := true
	switch s[0] {
	case ModPlus:
		isOverride = false
		plus = true
	case ModMinus:
		isOverride = false
		plus = false
	}

	chg := change{}

	if isOverride {
		chg.override = true
	} else {
		s = s[1:]
	}

	for len(s) > 0 {
		unit := strings.IndexAny(s, TimeUnits)
		if unit == -1 {
			return nil, ErrChangeSyntax
		}
		quantity, err := strconv.Atoi(s[:unit])
		if err != nil {
			return nil, err
		}
		if !plus {
			quantity = -quantity
		}

		unitRune, runeLen := utf8.DecodeRuneInString(s[unit:])
		with(&chg, quantity, unitRune)
		s = s[unit+runeLen:]
	}

	return &chg, nil
}

func with(change *change, quantity int, unit rune) {
	intUnit := 0
	switch unit {
	case 's':
		change.second = quantity
		intUnit = Second
	case 'm':
		change.minute = quantity
		intUnit = Minute
	case 'h':
		change.hour = quantity
		intUnit = Hour
	case 'd':
		change.day = quantity
		intUnit = Day
	case 'M':
		change.month = quantity
		intUnit = Month
	case 'y':
		change.year = quantity
		intUnit = Year
	}

	for i := 0; i <= intUnit; i++ {
		change.existFlag |= 1 << i
	}
}

func NewSecondChange(s int, override bool) change {
	return change{
		second:    s,
		existFlag: 1 << Second,
		override:  override,
	}
}

func NewMinuteChange(m int, s int, override bool) change {

	return change{
		minute:    m,
		second:    s,
		existFlag: 1<<Second | 1<<Minute,
		override:  override,
	}
}

func NewHourChange(h int, m int, s int, override bool) change {

	return change{
		hour:      h,
		minute:    m,
		second:    s,
		existFlag: 1<<Second | 1<<Minute | 1<<Hour,
		override:  override,
	}
}

func NewDayChange(d int, h int, m int, s int, override bool) change {

	return change{
		day:       d,
		hour:      h,
		minute:    m,
		second:    s,
		existFlag: 1<<Second | 1<<Minute | 1<<Hour | 1<<Day,
		override:  override,
	}
}

func NewMonthChange(M int, d int, h int, m int, s int, override bool) change {
	return change{
		month:     M,
		day:       d,
		hour:      h,
		minute:    m,
		second:    s,
		existFlag: 1<<Second | 1<<Minute | 1<<Hour | 1<<Day | 1<<Month,
		override:  override,
	}
}

func NewYearChange(y int, M int, d int, h int, m int, s int, override bool) change {
	return change{
		year:      y,
		month:     M,
		day:       d,
		hour:      h,
		minute:    m,
		second:    s,
		existFlag: 1<<Second | 1<<Minute | 1<<Hour | 1<<Day | 1<<Month | 1<<Year,
		override:  override,
	}
}
