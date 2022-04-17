package seed

import (
	"github.com/jinzhu/gorm"
	"log"
	"swimming-content-management/constants"
	"swimming-content-management/data/permission"
	"swimming-content-management/data/role"
)

var permissions = []permission.Permission{
	permission.Permission{
		Name: constants.CanViewSwimmerPerformanceOnly,
	},

	permission.Permission{
		Name: constants.CanViewPersonalDetailsOnly,
	},

	permission.Permission{
		Name: constants.CanManagePersonalDetails,
	},
	permission.Permission{
		Name: constants.CanManageSwimmerPerformance,
	},
	permission.Permission{
		Name: constants.CanViewRaceDataOnly,
	},
	permission.Permission{
		Name: constants.CanManageRaceData,
	},
}

var roles = []role.Role{

	role.Role{
		Name: constants.RoleSwimmer,
	},

	role.Role{
		Name: constants.RoleParent,
	},

	role.Role{
		Name: constants.RoleClubAdministrator,
	},

	role.Role{
		Name: constants.RoleCoaches,
	},

	role.Role{
		Name: constants.NonAdultSwimmer,
	},

	role.Role{
		Name: constants.NonAdultSwimmer,
	},

	role.Role{
		Name: constants.AllAccess,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().AutoMigrate(&role.Role{}, &permission.Permission{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range permissions {
		err = db.Debug().Model(&permission.Permission{}).Create(&permissions[i]).Error
		if err != nil {
			log.Fatalf("cannot seed permission: %v", err)
		}
		db.Debug().Create(&roles[i])
		db.Model(&roles[i]).Association("Permissions").Append(&permissions[i])
		//roles[i].Permissions = []*permission.Permission{&permissions[i]}
		//err = db.Debug().Model(&role.Role{}).Create(&roles[i]).Error
		if err != nil {
			log.Fatalf("cannot seed role table %v", err)
		}

	}

}

func DropRoleAndPermissionTables(db *gorm.DB) {
	//db.Table("role_permissions").RemoveForeignKey("role_id", "roles(id)")
	//db.Table("role_permissions").RemoveForeignKey("permission_id", "permissions(id)")
	err := db.Debug().DropTableIfExists(&permission.Permission{}, &role.Role{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
}
