package users

import (
	userdomain "swimming-content-management/domain/userdomain"
)

func toResponseModel(entity *userdomain.User) *UserResponse {
	return &UserResponse{
		Id:        entity.Id,
		Username:  entity.Username,
		Email:     entity.Email,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
