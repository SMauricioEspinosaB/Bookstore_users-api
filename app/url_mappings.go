package app

import (
	"github.com/SMauricioEspinosaB/Bookstore_users-api/controllers/ping"
	"github.com/SMauricioEspinosaB/Bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	//param route
	router.GET("/users/:user_id", users.Get)
	router.POST("/users", users.Create)
	router.PUT("/users:user_id", users.Update)
	router.PATCH("/users:user_id", users.Update)
	router.DELETE("/users:user_id", users.Delete)

	// param query
	router.GET("internal/users/search", users.Search)
}
