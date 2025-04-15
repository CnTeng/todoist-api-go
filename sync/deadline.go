package sync

import (
	"encoding/json"
	"time"
)

// Similar to due dates, deadlines can be set on tasks, and can be used to
// differentiate between when a task should be started, and when it must be done
// by.
//
// Unlike due dates, deadlines only support non-recurring dates with no time
// component.
//
// You can find our more information about [deadlines] in our Help Center.
//
// See [Deadlines] for more details.
//
// [deadlines]: https://www.todoist.com/help/articles/introduction-to-deadlines-uMqbSLM6U
// [Deadlines]: https://todoist.com/api/v1/docs#tag/Deadlines
type Deadline struct {
	// Deadline in the format of YYYY-MM-DD (RFC 3339).
	Date *time.Time `json:"date"`

	// Lang which has to be used to parse the content of the string attribute.
	// Used by clients and on the server side to properly process deadlines.
	//
	// Valid languages are:
	//   en, da, pl, zh, ko, de, pt, ja, it, fr, sv, ru, es, nl, fi, nb, tw.
	Lang *string `json:"lang"`
}

func (d *Deadline) UnmarshalJSON(data []byte) error {
	type alias Deadline
	aux := &struct {
		Date string `json:"date"`
		*alias
	}{alias: (*alias)(d)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	date, err := time.Parse(time.DateOnly, aux.Date)
	if err != nil {
		return err
	}
	d.Date = &date

	return nil
}

func (d *Deadline) MarshalJSON() ([]byte, error) {
	type alias Deadline
	aux := struct {
		Date string `json:"date"`
		*alias
	}{
		Date:  d.Date.Format(time.DateOnly),
		alias: (*alias)(d),
	}

	return json.Marshal(&aux)
}
