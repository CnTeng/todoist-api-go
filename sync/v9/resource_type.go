package sync

import (
	"encoding/json"
	"net/url"
)

type ResourceType string

const (
	Labels               ResourceType = "labels"
	Projects             ResourceType = "projects"
	Items                ResourceType = "items"
	Notes                ResourceType = "notes"
	Sections             ResourceType = "sections"
	Filters              ResourceType = "filters"
	Reminders            ResourceType = "reminders"
	RemindersLocation    ResourceType = "reminders_location"
	Locations            ResourceType = "locations"
	User                 ResourceType = "user"
	LiveNotifications    ResourceType = "live_notifications"
	Collaborators        ResourceType = "collaborators"
	UserSettings         ResourceType = "user_settings"
	NotificationSettings ResourceType = "notification_settings"
	UserPlanLimits       ResourceType = "user_plan_limits"
	CompletedInfo        ResourceType = "completed_info"
	Stats                ResourceType = "stats"
	All                  ResourceType = "all"
)

type ResourceTypes []ResourceType

func (rts *ResourceTypes) Add(rt ...ResourceType) {
	*rts = append(*rts, rt...)
}

func (rts *ResourceTypes) Del(rt ...ResourceType) {
	for _, r := range rt {
		*rts = append(*rts, "-"+r)
	}
}

func (rts *ResourceTypes) EncodeValues(key string, v *url.Values) error {
	b, err := json.Marshal(rts)
	if err != nil {
		return err
	}

	v.Add(key, string(b))
	return nil
}
