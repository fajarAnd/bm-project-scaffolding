package handlers

import (
	"net/http"

	"github.com/baramulti/ticketing-system/backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	// TODO: add user service when needed
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetMe(c *gin.Context) {
	// TODO: extract user from context
	// TODO: return user info
	response.Error(c, http.StatusNotImplemented, "not implemented")
}

func (h *UserHandler) Create(c *gin.Context) {
	// TODO: admin only - create new user
	response.Error(c, http.StatusNotImplemented, "not implemented")
}

func (h *UserHandler) Update(c *gin.Context) {
	// TODO: update user info
	response.Error(c, http.StatusNotImplemented, "not implemented")
}

func (h *UserHandler) Delete(c *gin.Context) {
	// TODO: admin only - delete user
	response.Error(c, http.StatusNotImplemented, "not implemented")
}