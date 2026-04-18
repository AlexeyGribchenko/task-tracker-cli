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

type GetTasksSorted struct {
	ColumnSorted string
}

type GetTasksFiltered struct {
	ColumnFiltered string
	FilterValue    string
}
