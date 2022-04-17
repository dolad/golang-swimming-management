package main

import (
	"net/http"
	"swimming-content-management/config"
	database "swimming-content-management/data/database"
	permissionStore "swimming-content-management/data/permission"
	roleStore "swimming-content-management/data/role"
	userStore "swimming-content-management/data/user"
	permissionDomain "swimming-content-management/domain/permission"
	roleDomain "swimming-content-management/domain/role"
	userDomain "swimming-content-management/domain/userdomain"
	router "swimming-content-management/router/http"
	"swimming-content-management/seed"
)

func main() {
	// get configuration stucts via .env file
	configuration, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	// establish DB connection
	db, err := database.Connect(configuration.Database)
	if err != nil {
		panic(err)
	}
	seed.DropRoleAndPermissionTables(db)
	seed.Load(db)
	//user repo and routes
	userRepository := userStore.New(db)
	userServices := userDomain.NewService(userRepository)

	// permission repo and routes
	permissionRepository := permissionStore.New(db)

	permissionServices := permissionDomain.NewService(permissionRepository)

	// role repo and routes
	roleRepository := roleStore.New(db)
	
	roleServices := roleDomain.NewService(roleRepository)

	httpRouter := router.NewHTTPHandler(userServices, permissionServices, roleServices)

	err = http.ListenAndServe(":"+configuration.Port, httpRouter)
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
