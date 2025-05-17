package sync

type TaskMoveArgs struct {
	// Required.
	// The ID of the task.
	ID string `json:"id"`

	// Optional.
	// ID of the destination parent task. The task becomes the last child task of
	// the parent task.
	ParentID *string `json:"parent_id,omitempty"`

	// Optional.
	// ID of the destination section. The task becomes the last root task of the
	// section.
	SectionID *string `json:"section_id,omitempty"`

	// Optional.
	// ID of the destination project. The task becomes the last root task of the
	// project.
	ProjectID *string `json:"project_id,omitempty"`
}

func (args *TaskMoveArgs) command() string {
	return "item_move"
}
