package permission

import domain "swimming-content-management/domain/permission"

func ToDbModel(entity *domain.Permission) *domain.Permission {

	return &domain.Permission{
		Name: entity.Name,
	}
}
