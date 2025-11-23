package organization

type CreateOrganizationDTO struct {
	Name string `json:"name" binding:"required"`
}
