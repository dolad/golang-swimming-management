package users

import (
	domainErrors "swimming-content-management/domain"
	authdomain "swimming-content-management/domain/authdomain"
	domain "swimming-content-management/domain/userdomain"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

const (
	createError = "Error in creating new User"

	loginError = "Error in login user"
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
	entity, err := toDbModel(user)
	if err != nil {
		appError := domainErrors.NewAppError(errors.Wrap(err, createError), domainErrors.RepositoryError)
		return nil, appError
	}
	existing, _ := GetUserByEmail(user, s.db)
	if existing != nil {
		return nil, errors.New("User already existing")
	}
	if err := s.db.Create(entity).Error; err != nil {
		appError := domainErrors.NewAppError(errors.Wrap(err, createError), domainErrors.RepositoryError)
		return nil, appError
	}

	return toDomainModel(entity), nil
}

func GetUserByEmail(user *domain.User, db *gorm.DB) (*domain.User, error) {
	result := &domain.User{}
	query := db.Where("email= ?", user.Email).First(result)
	if query.RecordNotFound() {
		appError := domainErrors.NewAppErrorWithType(domainErrors.NotFound)
		return nil, appError
	}

	if err := query.Error; err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, loginError), domainErrors.RepositoryError)
		return nil, appErr
	}

	return result, nil

}

func (s *Store) Login(user *domain.User) (*domain.AuthPayload, error) {
	result, err := GetUserByEmail(user, s.db)
	if err != nil {
		return nil, err
	}
	token, err := authdomain.GenerateAccessToken(result)
	if err != nil {
		return nil, err
	}

	authPayload := &domain.AuthPayload{
		Token: token,
		User:  result,
	}
	return authPayload, nil

}
