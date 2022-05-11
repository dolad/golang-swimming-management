package users

import (
	users "swimming-content-management/data/user"
)

func ToResponseModel(entity *users.User) *UserResponse {

	return &UserResponse{
		Id:          entity.ID,
		Username:    entity.Username,
		Email:       entity.Email,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		Surname:     entity.Surname,
		FirstName:   entity.FirstName,
		DateofBirth: entity.DateofBirth,
		PhoneNumber: entity.PhoneNumber,
		Address:     entity.Address,
		Country:     entity.Country,
		State:       entity.State,
		PostCode:    entity.PostCode,
		Role:        entity.Role,
	}
}

func toAuthResponseModel(entity *users.AuthPayload) *AuthPayloadResponse {
	return &AuthPayloadResponse{
		Id:          entity.ID,
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
		Role:        entity.Role,
		RoleID:      entity.RoleID,
		Token:       entity.Token,
	}
}
