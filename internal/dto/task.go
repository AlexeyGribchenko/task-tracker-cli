package dto

type CreateTask struct {
	Name        string
	Description *string
}

type UpdateTask struct {
	ID     int
	Status string
}
