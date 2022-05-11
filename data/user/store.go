package users

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"log"
	domainErrors "swimming-content-management/domain"
	"swimming-content-management/domain/role"
)

const (
	createError             = "Error in creating new User"
	loginError              = "Error in login user"
	ROLENOTFOUND            = "Role Not found"
	GettingRole             = "Error in getting default role"
	UserNameEmailExistError = "User with this email or username already existing"
	UserNotFound            = "User with this email not found"
	WrongPassword           = "Password does not match"
)

type Store struct {
	db *gorm.DB
}

type AuthPayload struct {
	Token string
	*User
}

// New creates a new Store struct
func New(db *gorm.DB) *Store {
	// migrate schema
	db.AutoMigrate(&User{})

	return &Store{
		db: db,
	}
}

func (s *Store) SignUp(user *User) (*User, error) {
	entity, err := toDbModel(user)
	tx := s.db.Begin()
	defer tx.Commit()
	if err != nil {
		appError := domainErrors.NewAppError(errors.Wrap(err, createError), domainErrors.RepositoryError)
		return nil, appError
	}

	defaultRole, err := s.FindRoleById(entity.RoleID)
	fmt.Println(defaultRole)
	if defaultRole.Name == "" {
		return nil, errors.New(err.Error())
	}

	existing, _ := GetUserByEmail(user.Email, s.db)
	userNameExist, _ := GetUserByUserName(user, s.db)

	if existing != nil || userNameExist != nil {
		return nil, errors.New(UserNameEmailExistError)
	}

	//skip the creation of swimming data when creating user
	if err := tx.Create(&entity).Error; err != nil {
		return nil, errors.New(createError)
	}

	return entity, nil
}

func GetUserByEmail(email string, db *gorm.DB) (*User, error) {
	var result User

	query := db.Preload("Role").Where("email= ?", email).First(&result)
	if query.RecordNotFound() {
		appError := domainErrors.NewAppErrorWithType(domainErrors.NotFound)
		return nil, appError
	}

	if err := query.Error; err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, loginError), domainErrors.RepositoryError)
		return nil, appErr
	}

	return &result, nil

}

func GetUserById(userId *uuid.UUID) (*User, error) {
	var db *gorm.DB
	var result User
	query := db.Preload("Role").Where("id= ?", userId).First(&result)
	if query.RecordNotFound() {
		appError := domainErrors.NewAppErrorWithType(domainErrors.NotFound)
		return nil, appError
	}

	if err := query.Error; err != nil {
		appErr := domainErrors.NewAppError(errors.Wrap(err, loginError), domainErrors.RepositoryError)
		return nil, appErr
	}

	return &result, nil

}

func GetUserByUserName(user *User, db *gorm.DB) (*User, error) {
	result := &User{}
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

func (s *Store) GetUserById(userId uuid.UUID) (*User, error) {
	result := &User{}
	fmt.Println((userId))
	query := s.db.Preload("Role").Where("id= ?", userId).First(&result)
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

func (s *Store) Login(user *User) (*AuthPayload, error) {
	result, err := GetUserByEmail(user.Email, s.db)

	if err != nil {
		return nil, errors.New(UserNotFound)
	}
	isPasswordMatch := VerifyHash(result.Password, user.Password)

	if isPasswordMatch != true {
		return nil, errors.New(WrongPassword)
	}

	if err != nil {
		return nil, err
	}
	token, err := GenerateAccessToken(result)
	if err != nil {
		return nil, err
	}
	authPayload := &AuthPayload{
		Token: token,
		User:  result,
	}
	return authPayload, nil

}

func (s *Store) FindRoleById(roleId uint32) (role.Role, error) {
	var result role.Role
	query := s.db.Where("id= ?", roleId).First(&result)

	if query.RecordNotFound() {
		return role.Role{}, errors.New(ROLENOTFOUND)
	}

	if err := query.Error; err != nil {
		return role.Role{}, errors.New(err.Error())
	}

	return result, nil

}

func (s *Store) GetRoleName(roleId uint32) (*role.Role, error) {
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

func (s *Store) UpdateUserRole(roleId uint32) (*User, error) {
	user := &User{}

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

func (s *Store) GetUsers() ([]User, error) {
	var result []User
	if err := s.db.Preload("Role").Omit("Password").Find(&result).Error; err != nil {
		log.Fatalf(err.Error())
		return nil, errors.New(err.Error())
	}

	for i, _ := range result {
		result[i].Password = ""
	}

	return result, nil

}

func (s *Store) UpdateUserProfile(user *User, userId uuid.UUID) (*User, error) {
	var result User

	updatePayloadStruct := &User{
		Country:     user.Country,
		State:       user.State,
		Address:     user.Address,
		FirstName:   user.FirstName,
		Surname:     user.Surname,
		PhoneNumber: user.PhoneNumber,
	}
	if err := s.db.Model(&result).Where("id= ?", userId).Updates(updatePayloadStruct).Error; err != nil {
		log.Fatalf(err.Error())
		return nil, errors.New(err.Error())
	}

	return &result, nil

}
