package users

func toDbModel(entity *User) (*User, error) {

	hashedPassword, err := Hash(entity.Password)
	if err != nil {
		return nil, err
	}
	return &User{

		ID:          entity.ID,
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
		Country:     entity.Country,
		State:       entity.State,
		PostCode:    entity.PostCode,
		RoleID:      entity.RoleID,
		SquadID:     entity.SquadID,
	}, nil
}

func toDomainModel(entity User) User {
	return User{
		ID:          entity.ID,
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
		SquadID:     entity.SquadID,
	}
}
