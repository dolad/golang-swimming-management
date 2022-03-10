package users

import (
	domainErrors "swimming-content-management/domain"
	domain "swimming-content-management/domain/userdomain"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

const (
	createError = "Error in creating new User"
)

type Store struct {
	db *gorm.DB
}

// New creates a new Store struct
func New(db *gorm.DB) *Store {
	// migrate schema
	db.AutoMigrate(&domain.User{})

	return &Store{
		db: db,
	}
}

func (s *Store) SignUp(user *domain.User) (*domain.User, error) {
	entity := toDbModel(user)
	if err := s.db.Create(entity).Error; err != nil {
		appError := domainErrors.NewAppError(errors.Wrap(err, createError), domainErrors.RepositoryError)
		return nil, appError
	}

	return toDomainModel(entity), nil
}
