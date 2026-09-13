package users

import (
	"errors"
	"net/http"
	"strconv"

	uc "ddd-structure/internal/application/v1/users_use_case"
	domain "ddd-structure/internal/domain/v1/users"
	dto "ddd-structure/internal/interface/dto/v1/users"
	"ddd-structure/package/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *uc.Service
}

func NewHandler(service *uc.Service) *Handler {
	return &Handler{service: service}
}

// CreateUser godoc
// @Summary Create user
// @Tags v1/users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body dto.CreateUserRequest true "body"
// @Success 200 {object} response.APIResponse
// @Router /users [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}

	user, err := h.service.Create(c.Request.Context(), domain.CreateUserCmd{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		writeUserError(c, err)
		return
	}
	response.Success(c, dto.MapUserToDTO(user))
}

// GetUser godoc
// @Summary Get user by id
// @Tags v1/users
// @Security BearerAuth
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.APIResponse
// @Router /users/{id} [get]
func (h *Handler) GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, "INVALID_REQUEST", "id must be a number")
		return
	}

	user, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		writeUserError(c, err)
		return
	}
	response.Success(c, dto.MapUserToDTO(user))
}

// ListUsers godoc
// @Summary List users
// @Tags v1/users
// @Security BearerAuth
// @Produce json
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {object} response.APIResponse
// @Router /users [get]
func (h *Handler) ListUsers(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	users, total, err := h.service.List(c.Request.Context(), domain.ListParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		_ = c.Error(err)
		response.Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "failed to list users")
		return
	}
	response.SuccessPaginated(c, dto.MapUsersToDTO(users), total, limit, offset)
}

// UpdateUser godoc
// @Summary Update user
// @Tags v1/users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param body body dto.UpdateUserRequest true "body"
// @Success 200 {object} response.APIResponse
// @Router /users/{id} [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, "INVALID_REQUEST", "id must be a number")
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}

	user, err := h.service.Update(c.Request.Context(), id, domain.UpdateUserCmd{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		IsActive:  req.IsActive,
	})
	if err != nil {
		writeUserError(c, err)
		return
	}
	response.Success(c, dto.MapUserToDTO(user))
}

// DeleteUser godoc
// @Summary Delete user
// @Tags v1/users
// @Security BearerAuth
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.APIResponse
// @Router /users/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, "INVALID_REQUEST", "id must be a number")
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		writeUserError(c, err)
		return
	}
	response.Success(c, gin.H{"deleted": true})
}

func writeUserError(c *gin.Context, err error) {
	_ = c.Error(err)

	var notFound domain.ErrNotFound
	var exists domain.ErrAlreadyExists
	var invalid domain.ErrInvalidInput

	switch {
	case errors.As(err, &notFound):
		response.Fail(c, http.StatusNotFound, "USER_NOT_FOUND", "user not found")
	case errors.As(err, &exists):
		response.Fail(c, http.StatusConflict, "USER_ALREADY_EXISTS", "user with this email already exists")
	case errors.As(err, &invalid):
		response.Fail(c, http.StatusBadRequest, "INVALID_INPUT", invalid.Error())
	default:
		response.Fail(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "unexpected error")
	}
}
