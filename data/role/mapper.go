package role

import (
	"swimming-content-management/domain/role"
)

func ToDbRoleModel(entity *role.Role) *role.Role {

	if entity.Name == "" {
		return &role.Role{
			Name: "swimmer",
		}
	}

	return &role.Role{
		Name: entity.Name,
	}
}
