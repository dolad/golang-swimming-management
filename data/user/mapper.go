package users

import (
	domain "swimming-content-management/domain/userdomain"
)

func toDbModel(entity *domain.User) (*User, error) {

	hashedPassword, err := Hash(entity.Password)
	if err != nil {
		return nil, err
	}
	return &User{
		Id:       entity.Id,
		Email:    entity.Email,
		Username: entity.Username,
		Password: hashedPassword,
	}, nil
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
