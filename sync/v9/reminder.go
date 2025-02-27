package sync

type Reminder struct {
	// ID is the ID of the reminder.
	ID string `json:"id"`

	// NotifyUID is the user ID which should be notified of the reminder.
	NotifyUID string `json:"notify_uid"`

	// ItemID is the item ID for which the reminder is about.
	ItemID string `json:"item_id"`

	// Type is the type of the reminder: relative, absolute, or location.
	Type string `json:"type"`

	// Due is the due date of the reminder.
	Due Due `json:"due"`

	// MinuteOffset is the relative time in minutes before the due date of the item when the reminder should be triggered.
	MinuteOffset int `json:"minute_offset"`

	// Name is an alias name for the location.
	Name *string `json:"name,omitempty"`

	// LocLat is the location latitude.
	LocLat *string `json:"loc_lat,omitempty"`

	// LocLong is the location longitude.
	LocLong *string `json:"loc_long,omitempty"`

	// LocTrigger is what should trigger the reminder: on_enter or on_leave.
	LocTrigger *string `json:"loc_trigger,omitempty"`

	// Radius is the radius around the location that is still considered as part of the location (in meters).
	Radius *int `json:"radius,omitempty"`

	// IsDeleted indicates whether the reminder is marked as deleted.
	IsDeleted bool `json:"is_deleted"`
}
