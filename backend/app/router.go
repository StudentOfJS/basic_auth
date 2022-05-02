package app

import "basic_auth/backend/controller/users"

func mapUrls() {
	router.POST("/api/register", users.Register)
	// router.POST("/api/login", users.Login)
}
