package handlers

import "auth/pkg/models"

type AuthHandlers interface{}

type Handlers struct {
	u *models.User
}

func NewAuth(u *models.User) *Handlers {
	return &Handlers{
		u: u,
	}
}

func (h *Handlers) Test() {}
