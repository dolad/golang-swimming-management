package main

import (
	"net/http"
	"swimming-content-management/config"
	database "swimming-content-management/data/database"
	userStore "swimming-content-management/data/user"
	userDomain "swimming-content-management/domain/userdomain"
	router "swimming-content-management/router/http"
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

	userRepository := userStore.New(db)
	userServices := userDomain.NewService(userRepository)

	httpRouter := router.NewHTTPHandler(userServices)
	err = http.ListenAndServe(":"+configuration.Port, httpRouter)
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
