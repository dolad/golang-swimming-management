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

		Id:          entity.Id,
		Username:    entity.Username,
		Email:       entity.Email,
		Password:    hashedPassword,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		Surname:     entity.Surname,
		FirstName:   entity.FirstName,
		DateofBirth: entity.DateofBirth,
		PhoneNumber: entity.PhoneNumber,
		Address:     entity.Address,
		PostCode:    entity.PostCode,
	}, nil
}

func toDomainModel(entity *User, roleName string) *domain.User {
	return &domain.User{
		Id:          entity.Id,
		Username:    entity.Username,
		Email:       entity.Email,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		Surname:     entity.Surname,
		FirstName:   entity.FirstName,
		DateofBirth: entity.DateofBirth,
		PhoneNumber: entity.PhoneNumber,
		Address:     entity.Address,
		PostCode:    entity.PostCode,
		RoleID:      entity.RoleID,
		RoleName:    roleName,
	}
}
