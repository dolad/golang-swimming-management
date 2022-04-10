package users

import (
	userdomain "swimming-content-management/domain/userdomain"
)

func toResponseModel(entity *userdomain.User) *UserResponse {

	return &UserResponse{
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
		RoleName:    entity.RoleName,
	}
}

func toAuthResponseModel(entity *userdomain.AuthPayload) *AuthPayloadResponse {
	return &AuthPayloadResponse{
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
		Token:       entity.Token,
	}
}
