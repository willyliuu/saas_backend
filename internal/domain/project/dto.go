package project

type CreateProjectDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Archieved   bool   `json:"archieved"`
}

type UpdateProjectDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Archieved   bool   `json:"archieved"`
}
