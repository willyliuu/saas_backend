package task

import (
	"gin/internal/database"
	"gin/internal/models"
)

type TaskRepository struct{}

func (r *TaskRepository) Create(task *models.Task) error {
	return database.DB.Create(task).Error
}

func (r *TaskRepository) FindByProjectID(projectID uint) ([]*models.Task, error) {
	var tasks []*models.Task
	err := database.DB.Find(&tasks, projectID).Error

	return tasks, err
}

func (r *TaskRepository) FindByID(id uint) (*models.Task, error) {
	var task *models.Task
	err := database.DB.First(&task).Error

	return task, err
}

func (r *TaskRepository) Update(task *models.Task) error {
	return database.DB.Save(task).Error
}

func (r *TaskRepository) Delete(id uint) error {
	return database.DB.Delete(&models.Task{}, id).Error
}
