package handlers

import (
	"net/http"

	"github.com/baramulti/ticketing-system/backend/internal/services"
	"github.com/baramulti/ticketing-system/backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userSvc services.UserService
}

func NewUserHandler(userSvc services.UserService) *UserHandler {
	return &UserHandler{
		userSvc: userSvc,
	}
}

func (h *UserHandler) GetMe(c *gin.Context) {
	// Extract user ID from context (set by auth middleware)
	userID, _ := c.Get("user_id")
	// TODO: Check user is Exist?


	user, err := h.userSvc.GetByID(c.Request.Context(), userID.(string))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to fetch user")
		return
	}

	response.Success(c, http.StatusOK, user)
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