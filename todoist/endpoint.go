package todoist

const (
	baseURL = "https://todoist.com/api/v1"

	syncEndpoint = baseURL + "/sync"

	TaskEndpoint                          = baseURL + "/tasks"
	TaskQuickAddEndpoint                  = TaskEndpoint + "/quick"
	TaskCompletedByCompletionDateEndpoint = TaskEndpoint + "/completed/by_completion_date"
	TaskCompletedByDueDateEndpoint        = TaskEndpoint + "/completed/by_due_date"
)
