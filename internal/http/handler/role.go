package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/go-echo-template/internal/http/dto"
	"github.com/sherwin-77/go-echo-template/internal/service"
	"github.com/sherwin-77/go-echo-template/pkg/response"
	"net/http"
)

type RoleHandler struct {
	RoleService service.RoleService
}

func NewRoleHandler(roleService service.RoleService) RoleHandler {
	return RoleHandler{roleService}
}

func (h *RoleHandler) GetRoles(ctx echo.Context) error {
	roles, err := h.RoleService.GetRoles(ctx.Request().Context())

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", roles, nil))
}

func (h *RoleHandler) GetRoleByID(ctx echo.Context) error {
	roleID := ctx.Param("id")
	if roleID == "" {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	role, err := h.RoleService.GetRoleByID(ctx.Request().Context(), roleID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Success", role, nil))
}

func (h *RoleHandler) CreateRole(ctx echo.Context) error {
	var req dto.RoleRequest

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if err := ctx.Validate(req); err != nil {
		return err
	}

	role, err := h.RoleService.CreateRole(ctx.Request().Context(), req)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, response.NewResponse(http.StatusCreated, "Role Created", role, nil))
}

func (h *RoleHandler) UpdateRole(ctx echo.Context) error {
	var req dto.UpdateRoleRequest

	if err := ctx.Bind(&req); err != nil {
		return err
	}

	if err := ctx.Validate(req); err != nil {
		return err
	}

	role, err := h.RoleService.UpdateRole(ctx.Request().Context(), req)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Role Updated", role, nil))
}

func (h *RoleHandler) DeleteRole(ctx echo.Context) error {
	roleID := ctx.Param("id")
	if roleID == "" {
		return echo.NewHTTPError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	if err := h.RoleService.DeleteRole(ctx.Request().Context(), roleID); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.NewResponse(http.StatusOK, "Role Deleted", nil, nil))
}
