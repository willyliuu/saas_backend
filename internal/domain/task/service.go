package task

import (
	"gin/internal/models"
)

type TaskService struct {
	Repo *TaskRepository
}

func (s *TaskService) Create(projectID uint, dto CreateTaskDTO) (*models.Task, error) {
	task := &models.Task{
		Title:       dto.Title,
		Description: dto.Description,
		Status:      dto.Status,
		DueDate:     dto.DueDate,
		ProjectID:   projectID,
		AssigneeID:  dto.AssigneeID,
	}

	if err := s.Repo.Create(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) FindByProjectID(projectID uint) ([]*models.Task, error) {
	return s.Repo.FindByProjectID(projectID)
}

func (s *TaskService) FindByID(id uint) (*models.Task, error) {
	return s.Repo.FindByID(id)
}

func (s *TaskService) Update(id uint, dto UpdateTaskDTO) (*models.Task, error) {
	task, err := s.Repo.FindByID(id)
	if err != nil {
		return task, err
	}

	task.Title = dto.Title
	task.Description = dto.Description
	task.Status = dto.Status
	task.DueDate = dto.DueDate
	task.AssigneeID = dto.AssigneeID

	if err := s.Repo.Update(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) Delete(id uint) error {
	return s.Repo.Delete(id)
}
