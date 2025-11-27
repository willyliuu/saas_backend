package organization

type CreateOrganizationDTO struct {
	Name string `json:"name" binding:"required"`
}

type UpdateOrganizationDTO struct {
	Name string `json:"name"`
}
