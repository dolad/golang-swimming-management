package role

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"log"
	"swimming-content-management/data/permission"
)

const ROLENOTFOUND = "ROLENOTFOUND"
const CANTCREATEROLE = "CANTCREATEROLE"

type Store struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Store {
	// migrate schema
	db.AutoMigrate(&Role{}, &permission.Permission{})

	return &Store{
		db: db,
	}
}

func (s *Store) CreateRole(roleEntity *Role) (*Role, error) {
	// check if role exist
	roleExist, _ := s.FindByName(roleEntity.Name)
	if roleExist.Name == "" {
		if err := s.db.Create(roleEntity).Error; err != nil {
			return nil, errors.New(CANTCREATEROLE)
		}
	}
	return roleEntity, nil
}

func (s *Store) FindByName(roleName string) (Role, error) {
	var result Role

	query := s.db.Preload("Permissions").Where("name= ?", roleName).First(&result)
	if query.RecordNotFound() {
		return Role{}, errors.New(ROLENOTFOUND)
	}
	if err := query.Error; err != nil {
		return Role{}, errors.New(err.Error())
	}

	return result, nil

}

func (s *Store) FindById(roleId uint32) (Role, error) {
	var result Role

	query := s.db.Preload("Permissions").Where("id= ?", roleId).First(&result)

	if query.RecordNotFound() {
		return Role{}, errors.New(ROLENOTFOUND)
	}

	if err := query.Error; err != nil {

		return Role{}, errors.New(err.Error())
	}

	return result, nil

}

func (s *Store) FindAll() ([]Role, error) {
	var result []Role
	if err := s.db.Preload("Permissions").Find(&result).Error; err != nil {
		log.Fatalf(err.Error())
		return nil, errors.New(err.Error())
	}

	return result, nil

}

//func (s *Store) addPermsissionToRole(roleId string, permissionId string) (*Role, error) {
//	permission, _ permissionRepo:
//	role, _ := s.findById(roleId)
//
//}
