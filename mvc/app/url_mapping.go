package app

import "github.com/Go-Lang/mvc/controllers"

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
