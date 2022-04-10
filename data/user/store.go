package users

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"swimming-content-management/constants"
	domainErrors "swimming-content-management/domain"
	"swimming-content-management/domain/authdomain"
	"swimming-content-management/domain/role"
	domain "swimming-content-management/domain/userdomain"
)

const (
	createError = "Error in creating new User"

	loginError = "Error in login user"

	GettingRole = "Error in getting default role"

	UserNameEmailExistError = "User with this email or username already existing"
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
	tx := s.db.Begin()
	defer tx.Commit()
	if err != nil {
		appError := domainErrors.NewAppError(errors.Wrap(err, createError), domainErrors.RepositoryError)
		return nil, appError
	}
	existing, _ := GetUserByEmail(user, s.db)
	userNameExist, _ := GetUserByUserName(user, s.db)

	if existing != nil || userNameExist != nil {
		return nil, errors.New(UserNameEmailExistError)
	}

	//assign default role to user
	defaultRole, err := GetDefaultRole(s.db)

	if err != nil {

		return nil, errors.New(GettingRole)
	}
	entity.RoleID = defaultRole.ID
	roleName, err := s.GetRoleName(defaultRole.ID)
	if err != nil {
		return nil, errors.New(GettingRole)
	}

	if err := tx.Create(entity).Error; err != nil {
		return nil, errors.New(createError)
	}

	return toDomainModel(entity, roleName.Name), nil
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

func GetUserByUserName(user *domain.User, db *gorm.DB) (*domain.User, error) {
	result := &domain.User{}
	query := db.Where("username= ?", user.Username).First(result)
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

func (s *Store) GetUserById(userId uuid.UUID, db *gorm.DB) (*domain.User, error) {
	result := &domain.User{}
	query := s.db.First(result, userId)
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

func GetDefaultRole(db *gorm.DB) (*role.Role, error) {
	result := &role.Role{}
	query := db.Where("name = ?", constants.RoleSwimmer).First(result)
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

func (s *Store) GetRoleName(roleId uint) (*role.Role, error) {
	result := &role.Role{}
	query := s.db.First(result, roleId)
	if query.RecordNotFound() {
		appError := domainErrors.NewAppErrorWithType(domainErrors.NotFound)
		return nil, appError
	}

	if err := query.Error; err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, GettingRole), domainErrors.RepositoryError)
		return nil, appErr
	}
	return result, nil
}

func (s *Store) UpdateUserRole(roleId uint) (*domain.User, error) {
	user := &domain.User{}

	query := s.db.First(user, roleId)
	if query.RecordNotFound() {
		appError := domainErrors.NewAppErrorWithType(domainErrors.NotFound)
		return nil, appError
	}

	if err := query.Error; err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, GettingRole), domainErrors.RepositoryError)
		return nil, appErr
	}
	user.RoleID = roleId
	s.db.Save(user)

	return user, nil
}
