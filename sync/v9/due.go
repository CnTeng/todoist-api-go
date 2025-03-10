package sync

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"time"
)

const floatingDateTime = "2006-01-02T15:04:05.000000"

// Due dates for tasks and reminders is one of the core concepts of Todoist.
// It's very powerful and quite complex, because it has to embrace different
// use-cases of Todoist users.
// Todoist supports three types of due dates:
//
//  1. Full-day dates (like "1 January 2018" or "tomorrow")
//
//     - Date is [time.DateOnly].
//
//     - Timezone is null.
//
//  2. Floating due dates with time (like "1 January 2018 at 12:00" or "tomorrow
//     at 10am")
//
//     - Date is "2006-01-02T15:04:05.000000".
//
//     - Timezone is null.
//
//  3. Due dates with time and fixed timezone (like "1 January 2018 at 12:00
//     America/Chicago" or "tomorrow at 10am Asia/Jakarta")
//
//     - Date is [time.RFC3339].
//
//     - Timezone is not null.
//
// Unless specified explicitly, dates with time are created as floating.
//
// In addition, any of these due dates can be set to recurring or not, depending
// on the date string, provided by the client.
//
// See [Due dates] for more details.
//
// [Due dates]: https://developer.todoist.com/sync/v9#due-dates
type Due struct {
	// Due date.
	Date time.Time `json:"date"`

	// Timezone of the due instance.
	//
	// Used to recalculate properly the next iteration for a recurring due date.
	Timezone *time.Location `json:"timezone"`

	// Human-readable representation of due date. String always represents the due
	// object in user's timezone. See
	// https://todoist.com/help/articles/introduction-to-dates-and-time-q7VobO for
	// more details.
	String string `json:"string"`

	// Lang which has to be used to parse the content of the string attribute.
	// Used by clients and on the server side to properly process due dates when
	// date object is not set, and when dealing with recurring tasks.
	//
	// Valid languages are:
	//   en, da, pl, zh, ko, de, pt, ja, it, fr, sv, ru, es, nl, fi, nb, tw.
	Lang string `json:"lang"`

	// Boolean flag which is set to true if the due object represents a recurring
	// due date.
	IsRecurring bool `json:"is_recurring"`
}

func (d *Due) UnmarshalJSON(data []byte) error {
	type alias Due
	aux := &struct {
		Date     string  `json:"date"`
		Timezone *string `json:"timezone"`
		*alias
	}{alias: (*alias)(d)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	date, err := parseDate(aux.Date)
	if err != nil {
		return nil
	}
	d.Date = date

	if aux.Timezone != nil {
		tz, err := parseTimezone(*aux.Timezone)
		if err != nil {
			return err
		}
		d.Timezone = tz
	}

	return nil
}

func (d *Due) MarshalJSON() ([]byte, error) {
	type alias Due
	aux := struct {
		Date     string  `json:"date"`
		Timezone *string `json:"timezone"`
		*alias
	}{alias: (*alias)(d)}

	if d.Timezone == nil {
		if d.Date.Hour() == 0 && d.Date.Minute() == 0 && d.Date.Second() == 0 {
			aux.Date = d.Date.Format(time.DateOnly)
		} else {
			aux.Date = d.Date.Format(floatingDateTime)
		}
	} else {
		aux.Date = d.Date.Format(time.RFC3339)

		tz := d.Timezone.String()
		aux.Timezone = &tz
	}

	return json.Marshal(&aux)
}

func parseDate(date string) (time.Time, error) {
	dateFormats := []string{
		time.DateOnly,
		floatingDateTime,
		time.RFC3339,
	}

	for _, format := range dateFormats {
		t, err := time.Parse(format, date)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, errors.New("invalid date format")
}

func parseTimezone(timezone string) (*time.Location, error) {
	loc, err := time.LoadLocation(timezone)
	if err == nil {
		return loc, nil
	}

	re := regexp.MustCompile(`GMT ([+-]\d{1,2}):00`)
	matches := re.FindStringSubmatch(timezone)

	if len(matches) != 2 {
		return nil, errors.New("invalid GMT offset format")
	}

	offset, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil, err
	}

	return time.FixedZone(timezone, offset*60*60), nil
}
