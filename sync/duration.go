package sync

import "fmt"

// Duration represents a task's duration.
type Duration struct {
	// Time the task will take.
	Amount int `json:"amount"`

	// The amount represents which must be either minute or day.
	Unit string `json:"unit"`
}

func (d Duration) String() string {
	return fmt.Sprintf("%d %s", d.Amount, d.Unit)
}

func ParseDuration(s string) (*Duration, error) {
	var d Duration
	if _, err := fmt.Sscanf(s, "%d %s", &d.Amount, &d.Unit); err != nil {
		return nil, err
	}

	if d.Unit != "minute" && d.Unit != "day" {
		return nil, fmt.Errorf("invalid unit %q", d.Unit)
	}

	return &d, nil
}
