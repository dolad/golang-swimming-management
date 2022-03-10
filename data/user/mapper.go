package users

import (
	domain "swimming-content-management/domain/userdomain"
)

func toDbModel(entity *domain.User) *User {
	return &User{
		Id:       entity.Id,
		Email:    entity.Email,
		Username: entity.Username,
		Password: entity.Password,
	}
}

func toDomainModel(entity *User) *domain.User {
	return &domain.User{
		Id:        entity.Id,
		Email:     entity.Email,
		Username:  entity.Username,
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
