package users

func toDbModel(entity *User) (*User, error) {

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
		RoleID:      entity.RoleID,
	}, nil
}

func toDomainModel(entity *User) *User {
	return &User{
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
		Role:        entity.Role,
	}
}
