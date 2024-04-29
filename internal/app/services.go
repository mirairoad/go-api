package application

import "github.com/mirairoad/goApi/internal/app/services/users"

type Services struct {
	users users.UserServiceImpl
}

func (a *App) loadServices() {
	a.services.users = users.UserServiceImpl{}
}
