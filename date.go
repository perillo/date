// Copyright 2015 Manlio Perillo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate stringer -type=Weekday,Month

// Package date implements support for Gregorian date, following ISO 8601
// standard.
package date

import (
	"time"
)

// A Duration represents the elapsed time between two dates as an int32 day count.
type Duration int32

const (
	Day  Duration = 1
	Week Duration = 7
)

// A Weekday specifies a day of the week, as per ISO 8601 (Monday = 1, ...).
type Weekday int

const (
	Monday Weekday = 1 + iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// A Month specifies a month of the year (January = 1, ...).
type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

// These are predefined layouts for use in Date.Format and Date.Parse.
// The reference time used in the layouts is the specific date:
//	Mon Jan 2 2006
const (
	ANSIC   = "Mon Jan _2 2006"
	RFC822  = "02 Jan 06"
	RFC850  = "Monday, 02-Jan-06"
	RFC1123 = "Mon, 02 Jan 2006"
	RFC3339 = "2006-01-02"
)

// A Date represents a Gregorian date.
type Date struct {
	// The implementation uses time.Time to keep code simple; it should be
	// int64.
	tm time.Time
}

// New returns the Date corresponding to yyyy-mm-dd.
func New(year int, month Month, day int) Date {
	tm := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	return Date{tm}
}

// newFromTime returns the Date in witch t occurs.
func newFromTime(t time.Time) Date {
	year, month, day := t.Date()

	return New(year, Month(month), day)
}

// Today returns the current date.
func Today() Date {
	tm := time.Now()

	return newFromTime(tm)
}

// Parse parses a formatted string and returns the date value it represents.
// The layout defines the format by showing how the reference date, defined to be
//  Mon Jan 2 2006
// would be interpreted if it were the value.
func Parse(layout, value string) (Date, error) {
	tm, err := time.Parse(layout, value)
	if err != nil {
		return Date{}, err
	}

	return newFromTime(tm), nil
}

// Format returns a textual representation of the date value formatted
// according to layout, which defines the format by showing how the reference
// date, defined to be
//  Mon Jan 2 2006
// would be displayed if it were the value.
func (d Date) Format(layout string) string {
	return d.tm.Format(layout)
}

// Time returns the Time when the midnight of d occurs, in UTC.
func (d Date) Time() time.Time {
	return d.tm
}

// String returns an RFC3339/ISO-8601 date string, of the form "2006-01-02".
func (d Date) String() string {
	return d.Format(RFC3339)
}

// After reports whether the date d is after u.
func (d Date) After(u Date) bool {
	return d.tm.After(u.tm)
}

// Before reports whether the date d is before u.
func (d Date) Before(u Date) bool {
	return d.tm.Before(u.tm)
}

// Equal reports whether d and u represent the same date.
func (d Date) Equal(u Date) bool {
	return d.tm.Equal(u.tm)
}

// IsZero reports whether d represents the zero date,
// January 1, year 1.
func (d Date) IsZero() bool {
	return d.tm.IsZero()
}

// Add returns the date d + dd.
func (d Date) Add(dd Duration) Date {
	tm := d.tm.Add(time.Duration(dd) * 24 * time.Hour)

	return newFromTime(tm)
}

// AddDate returns the date corresponding to adding the given number of years,
// months, and days to d.
func (d Date) AddDate(years int, months int, days int) Date {
	tm := d.tm.AddDate(years, months, days)

	return newFromTime(tm)
}

// Date returns the year, month, and day of d.
func (d Date) Date() (year int, month Month, day int) {
	yy, mm, dd := d.tm.Date()

	return yy, Month(mm), dd
}

// Year returns the year of d.
func (d Date) Year() int {
	return d.tm.Year()
}

// Month returns the month of the year specified by d.
func (d Date) Month() Month {
	return Month(d.tm.Month())
}

// Day returns the day of the month specified by d.
func (d Date) Day() int {
	return d.tm.Day()
}

// Weekday returns the day of the week specified by d.
func (d Date) Weekday() Weekday {
	return Weekday((d.tm.Weekday() + 6) % 7)
}

// Week returns the week number specified by d.
func (d Date) Week() int {
	_, week := d.tm.ISOWeek()

	return week
}
