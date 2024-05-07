package utils

import (
	"context"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// origin from https://github.com/robfig/cron/tree/v3.0.1, base on parser.go and spec.go

// six columns mean：
//       second：0-59
//       minute：0-59
//       hour：1-23
//       day：1-31
//       month：1-12
//       week：0-6（0 means Sunday）
// SetCron some signals：
//       *： any time
//       ,：　 separate signal
//　　    －：duration
//       /n : do as n times of time duration
/////////////////////////////////////////////////////////
//	0/30 * * * * *                        every 30s
//	0 43 21 * * *                         21:43
//	0 15 05 * * * 　　                     05:15
//	0 0 17 * * *                          17:00
//	0 0 17 * * 1                           17:00 in every Monday
//	0 0,10 17 * * 0,2,3                   17:00 and 17:10 in every Sunday, Tuesday and Wednesday
//	0 0-10 17 1 * *                       17:00 to 17:10 in 1 min duration each time on the first day of month
//	0 0 0 1,15 * 1                        0:00 on the 1st day and 15th day of month
//	0 42 4 1 * * 　 　                     4:42 on the 1st day of month
//	0 0 21 * * 1-6　　                     21:00 from Monday to Saturday
//	0 0,10,20,30,40,50 * * * *　           every 10 min duration
//	0 */10 * * * * 　　　　　　              every 10 min duration
//	0 * 1 * * *　　　　　　　　               1:00 to 1:59 in 1 min duration each time
//	0 0 1 * * *　　　　　　　　               1:00
//	0 0 */1 * * *　　　　　　　               0 min of hour in 1 hour duration
//	0 0 * * * *　　　　　　　　               0 min of hour in 1 hour duration
//	0 2 8-20/3 * * *　　　　　　             8:02, 11:02, 14:02, 17:02, 20:02
//	0 30 5 1,15 * *　　　　　　              5:30 on the 1st day and 15th day of month

const _starBit = 1 << 63 // Set the top bit if a star was included in the expression.
const _yearLimit = 5

// bounds provides a range of acceptable values (plus a map of name to value).
type bounds struct {
	min, max uint
	names    map[string]uint
}

// The bounds for each field.
var (
	_seconds = bounds{0, 59, nil}
	_minutes = bounds{0, 59, nil}
	_hours   = bounds{0, 23, nil}
	_days    = bounds{1, 31, nil}
	_months  = bounds{1, 12, map[string]uint{
		"jan": 1,
		"feb": 2,
		"mar": 3,
		"apr": 4,
		"may": 5,
		"jun": 6,
		"jul": 7,
		"aug": 8,
		"sep": 9,
		"oct": 10,
		"nov": 11,
		"dec": 12,
	}}
	_weeks = bounds{0, 6, map[string]uint{
		"sun": 0,
		"mon": 1,
		"tue": 2,
		"wed": 3,
		"thu": 4,
		"fri": 5,
		"sat": 6,
	}}
)

// NextTime returns the next time base on spec(linux cron), start from param start
func NextTime(spec string, start time.Time) (next time.Time, err error) {

	var second, minute, hour, dom, month, dow uint64 // dom day of month, dow day of week
	if second, minute, hour, dom, month, dow, err = parseSpec(spec); err != nil {
		return
	}
	next, err = nextTime(second, minute, hour, dom, month, dow, start)
	return
}

// TickerRunWithContext will block current goroutine, usage: go TickerRunWithContext(...)
func TickerRunWithContext(ctx context.Context, spec string, start time.Time, run func(errTickerRun error, now time.Time)) {

	if ctx == nil {
		run(errors.New("func TickerRunWithContext param ctx is nil"), time.Now())
		return
	}
	var second, minute, hour, dom, month, dow uint64 // dom day of month, dow day of week
	var err error
	if second, minute, hour, dom, month, dow, err = parseSpec(spec); err != nil {
		run(err, time.Now())
		return
	}
	var next, now time.Time
	var chanRun = make(chan time.Time, 1)
	for {
		if next, err = nextTime(second, minute, hour, dom, month, dow, start); err != nil {
			run(err, time.Now())
			return
		}
		go func(next, start time.Time) {
			time.Sleep(next.Sub(start))
			chanRun <- next
		}(next, start)
		start = next
		select {
		case <-ctx.Done():
			run(ctx.Err(), time.Now())
			return
		case now = <-chanRun:
			go run(nil, now)
		}
	}
}

func TickerRun(spec string, start time.Time, run func(errTickerRun error, now time.Time)) {

	var second, minute, hour, dom, month, dow uint64 // dom day of month, dow day of week
	var err error
	if second, minute, hour, dom, month, dow, err = parseSpec(spec); err != nil {
		run(err, time.Now())
		return
	}
	var next time.Time
	for {
		if next, err = nextTime(second, minute, hour, dom, month, dow, start); err != nil {
			run(err, time.Now())
			return
		}
		time.Sleep(next.Sub(start))
		start = next
		go run(nil, next)
	}
}

// parseSpec dom day of month, dow day of week
func parseSpec(spec string) (second, minute, hour, dom, month, dow uint64, err error) {

	// Split on whitespace.  We require 5 or 6 fields.
	// (second) (minute) (hour) (day of month) (month) (day of week, optional)
	var fields = strings.Fields(spec)
	var lenFields = len(fields)
	if lenFields != 5 && lenFields != 6 {
		err = errors.New("spec field count must be 5 or 6, format (second) (minute) (hour) (day of month) (month) (day of week, optional)")
		return
	}
	// If a sixth field is not provided (DayOfWeek), then it is equivalent to star.
	if lenFields == 5 {
		fields = append(fields, "*")
	}
	if second, err = getField(fields[0], _seconds); err != nil {
		return
	}
	if minute, err = getField(fields[1], _minutes); err != nil {
		return
	}
	if hour, err = getField(fields[2], _hours); err != nil {
		return
	}
	if dom, err = getField(fields[3], _days); err != nil {
		return
	}
	if month, err = getField(fields[4], _months); err != nil {
		return
	}
	dow, err = getField(fields[5], _weeks)
	return
}

func nextTime(second, minute, hour, dom, month, dow uint64, start time.Time) (next time.Time, err error) {

	// General approach
	//
	// For Month, Day, Hour, Minute, Second:
	// Check if the time value matches.  If yes, continue to the next field.
	// If the field doesn't match the schedule, then increment the field until it matches.
	// While incrementing the field, a wrap-around brings it back to the beginning
	// of the field list (since it is necessary to re-verify previous field
	// values)

	// Note that without a time zone, start.Location() are treated
	// as local to the time provided.

	// Start at the earliest possible time (the upcoming second).
	next = start.Add(1*time.Second - time.Duration(start.Nanosecond())*time.Nanosecond)
	// This flag indicates whether a field has been incremented.
	var added = false
	var yearLimit = next.Year() + _yearLimit

WRAP:
	if next.Year() > yearLimit {
		err = errors.Errorf("year over limit, limit is %d", _yearLimit)
		return
	}
	// Find the first applicable month.
	// If it's this month, then do nothing.
	for 1<<uint(next.Month())&month == 0 {
		// If we have to add a month, reset the other parts to 0.
		if !added {
			added = true
			// Otherwise, set the date at the beginning (since the current time is irrelevant).
			next = time.Date(next.Year(), next.Month(), 1, 0, 0, 0, 0, next.Location())
		}
		next = next.AddDate(0, 1, 0)
		// Wrapped around.
		if next.Month() == time.January {
			goto WRAP
		}
	}
	// Now get a day in that month.
	//
	// NOTE: This causes issues for daylight savings regimes where midnight does
	// not exist.  For example: Sao Paulo has DST that transforms midnight on
	// 11/3 into 1am. Handle that by noticing when the Hour ends up != 0.
	for !dayMatches(dom, dow, next) {
		if !added {
			added = true
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
		}
		next = next.AddDate(0, 0, 1)
		// Notice if the hour is no longer midnight due to DST.
		// Add an hour if it's 23, subtract an hour if it's 1.
		if next.Hour() != 0 {
			if next.Hour() > 12 {
				next = next.Add(time.Duration(24-next.Hour()) * time.Hour)
			} else {
				next = next.Add(time.Duration(-next.Hour()) * time.Hour)
			}
		}
		if next.Day() == 1 {
			goto WRAP
		}
	}

	for 1<<uint(next.Hour())&hour == 0 {
		if !added {
			added = true
			next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), 0, 0, 0, next.Location())
		}
		next = next.Add(1 * time.Hour)
		if next.Hour() == 0 {
			goto WRAP
		}
	}

	for 1<<uint(next.Minute())&minute == 0 {
		if !added {
			added = true
			next = next.Truncate(time.Minute)
		}
		next = next.Add(1 * time.Minute)
		if next.Minute() == 0 {
			goto WRAP
		}
	}

	for 1<<uint(next.Second())&second == 0 {
		if !added {
			added = true
			next = next.Truncate(time.Second)
		}
		next = next.Add(1 * time.Second)
		if next.Second() == 0 {
			goto WRAP
		}
	}
	return
}

// dayMatches returns true if the next's day-of-week and day-of-month
// restrictions are satisfied by the given time.
func dayMatches(dom, dow uint64, next time.Time) (isMatch bool) {

	var domMatch = 1<<uint(next.Day())&dom > 0
	var dowMatch = 1<<uint(next.Weekday())&dow > 0
	if dom&_starBit > 0 || dow&_starBit > 0 {
		isMatch = domMatch && dowMatch
	} else {
		isMatch = domMatch || dowMatch
	}
	return
}

// getField returns an Int with the bits set representing all of the times that
// the field represents or error parsing field value.  A "field" is a comma-separated
// list of "ranges".
func getField(field string, r bounds) (bits uint64, err error) {

	var ranges = strings.FieldsFunc(field, func(r rune) bool { return r == ',' })
	var bit uint64
	var expr string
	for _, expr = range ranges {
		if bit, err = getRange(expr, r); err != nil {
			return
		}
		bits |= bit
	}
	return
}

// getRange returns the bits indicated by the given expression:
//
//	number | number "-" number [ "/" number ]
//
// or error parsing range.
func getRange(expr string, r bounds) (bits uint64, err error) {

	var start, end, step uint
	var rangeAndStep = strings.Split(expr, "/")
	var lowAndHigh = strings.Split(rangeAndStep[0], "-")
	var lenLowAndHigh = len(lowAndHigh)
	var singleDigit = lenLowAndHigh == 1
	var extra uint64

	if lowAndHigh[0] == "*" || lowAndHigh[0] == "?" {
		start, end, extra = r.min, r.max, _starBit
	} else {
		if start, err = parseIntOrName(lowAndHigh[0], r.names); err != nil {
			return
		}
		switch lenLowAndHigh {
		case 1:
			end = start
		case 2:
			if end, err = parseIntOrName(lowAndHigh[1], r.names); err != nil {
				return
			}
		default:
			err = errors.Errorf("too many hyphens: %s", expr)
			return
		}
	}

	switch len(rangeAndStep) {
	case 1:
		step = 1
	case 2:
		if step, err = mustParseInt(rangeAndStep[1]); err != nil {
			return
		}
		// Special handling: "N/step" means "N-max/step".
		if singleDigit {
			end = r.max
		}
		if step > 1 {
			extra = 0
		}
	default:
		err = errors.Errorf("too many slashes: %s", expr)
		return
	}

	if start < r.min {
		err = errors.Errorf("beginning of range (%d) below minimum (%d): %s", start, r.min, expr)
		return
	}
	if end > r.max {
		err = errors.Errorf("end of range (%d) above maximum (%d): %s", end, r.max, expr)
		return
	}
	if start > end {
		err = errors.Errorf("beginning of range (%d) beyond end of range (%d): %s", start, end, expr)
		return
	}
	if step == 0 {
		err = errors.Errorf("step of range should be a positive number: %s", expr)
		return
	}
	bits = getBits(start, end, step) | extra
	return
}

// parseIntOrName returns the (possibly-named) integer contained in expr.
func parseIntOrName(expr string, names map[string]uint) (namedInt uint, err error) {

	if names != nil {
		var ok bool
		if namedInt, ok = names[strings.ToLower(expr)]; ok {
			return
		}
	}
	namedInt, err = mustParseInt(expr)
	return
}

// mustParseInt parses the given expression as an int or returns an error.
func mustParseInt(expr string) (uiNum uint, err error) {

	var num int
	if num, err = strconv.Atoi(expr); err != nil {
		err = errors.Wrapf(err, "failed to parse int from %s", expr)
		return
	}
	if num < 0 {
		err = errors.Errorf("negative number (%d) not allowed: %s", num, expr)
		return
	}
	uiNum = uint(num)
	return
}

// getBits sets all bits in the range [min, max], modulo the given step size.
func getBits(min, max, step uint) (bits uint64) {

	// If step is 1, use shifts, else, use a simple loop.
	if step == 1 {
		bits = ^(math.MaxUint64 << (max + 1)) & (math.MaxUint64 << min)
		return
	} else {
		for i := min; i <= max; i += step {
			bits |= 1 << i
		}
	}
	return
}
