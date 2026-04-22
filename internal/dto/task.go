package dto

type CreateTask struct {
	Name        string
	Description *string
}

type UpdateTask struct {
	ID     int
	Status string
}

type RemoveTask struct {
	ID int
}

type GetTaskList struct {
	Status string
	SortBy string
	Desc   bool
}
