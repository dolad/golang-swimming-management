package permission

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

const PERMISSIONNOTFOUND = "PERMISSIONNOTFOUND"
const CANNOTCREATEPERMISSION = "CANNOTCREATEPERMISSION"

type Store struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Store {
	// migrate schema
	return &Store{
		db: db,
	}
}

func (s *Store) CreatePermission(permissionEntity *Permission) (*Permission, error) {
	// check if permission exist
	permissionExist, _ := s.FindByName(permissionEntity.Name)
	if permissionExist != nil {
		//	create a new role
		if err := s.db.Create(permissionExist).Error; err != nil {
			return nil, errors.New(CANNOTCREATEPERMISSION)
		}
	}
	return permissionEntity, nil
}

func (s *Store) FindByName(permissionName string) (*Permission, error) {
	result := &Permission{}
	query := s.db.Where("name= ?", permissionName).First(result)

	if query.RecordNotFound() {
		return nil, errors.New(PERMISSIONNOTFOUND)
	}

	if err := query.Error; err != nil {

		return nil, errors.New(err.Error())
	}

	return result, nil

}

func (s *Store) FindById(id uint32) (*Permission, error) {
	result := &Permission{}
	query := s.db.Where("id= ?", id).First(result)

	if query.RecordNotFound() {
		return nil, errors.New(PERMISSIONNOTFOUND)
	}

	if err := query.Error; err != nil {
		return nil, errors.New(err.Error())
	}

	return result, nil

}

func (s *Store) FindAll() ([]Permission, error) {
	var result []Permission
	if err := s.db.Find(&result).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	return result, nil

}
