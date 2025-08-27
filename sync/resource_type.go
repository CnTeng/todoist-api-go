package sync

// ResourceType represents a type of resource that can be fetched from the
// server.
type ResourceType string

const (
	UserData             ResourceType = "user"
	Projects             ResourceType = "projects"
	Tasks                ResourceType = "items"
	TaskComments         ResourceType = "notes"
	ProjectComments      ResourceType = "project_notes"
	Sections             ResourceType = "sections"
	Labels               ResourceType = "labels"
	Filters              ResourceType = "filters"
	Reminders            ResourceType = "reminders"
	RemindersLocation    ResourceType = "reminders_location"
	Locations            ResourceType = "locations"
	Collaborators        ResourceType = "collaborators"
	CollaboratorStates   ResourceType = "collaborator_states"
	LiveNotifications    ResourceType = "live_notifications"
	UserSettings         ResourceType = "user_settings"
	NotificationSettings ResourceType = "notification_settings"
	UserPlanLimits       ResourceType = "user_plan_limits"
	Workspaces           ResourceType = "workspaces"
	WorkspaceUsers       ResourceType = "workspace_users"
	CompletedInfo        ResourceType = "completed_info"
	Stats                ResourceType = "stats"
	All                  ResourceType = "all"

	NoUserData             ResourceType = "-user"
	NoProjects             ResourceType = "-projects"
	NoTasks                ResourceType = "-items"
	NoTaskComments         ResourceType = "-notes"
	NoProjectComments      ResourceType = "-project_notes"
	NoSections             ResourceType = "-sections"
	NoLabels               ResourceType = "-labels"
	NoFilters              ResourceType = "-filters"
	NoReminders            ResourceType = "-reminders"
	NoRemindersLocation    ResourceType = "-reminders_location"
	NoLocations            ResourceType = "-locations"
	NoCollaborators        ResourceType = "-collaborators"
	NoCollaboratorStates   ResourceType = "-collaborator_states"
	NoLiveNotifications    ResourceType = "-live_notifications"
	NoUserSettings         ResourceType = "-user_settings"
	NoNotificationSettings ResourceType = "-notification_settings"
	NoUserPlanLimits       ResourceType = "-user_plan_limits"
	NoWorkspaces           ResourceType = "-workspaces"
	NoWorkspaceUsers       ResourceType = "-workspace_users"
	NoCompletedInfo        ResourceType = "-completed_info"
	NoStats                ResourceType = "-stats"
)

// ResourceTypes represents resources that can be fetched from the server.
type ResourceTypes []ResourceType

// DefaultResourceTypes is the default set of resource to fetch from the
// server.
var DefaultResourceTypes = ResourceTypes{
	All,
	NoTaskComments,
	NoProjectComments,
	NoCollaborators,
	NoCollaboratorStates,
	NoNotificationSettings,
	NoUserPlanLimits,
	NoWorkspaces,
	NoWorkspaceUsers,
	NoCompletedInfo,
	NoStats,
}
