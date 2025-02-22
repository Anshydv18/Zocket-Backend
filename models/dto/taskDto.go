package dto

type Task struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Deadline      string `json:"deadline"`
	Priority      string `json:"priority"`
	Progess       string `json:"progress"`
	AssigneeEmail string `json:"assigneeEmail"`
	CreatedBy     string `json:"created_by"`
}
