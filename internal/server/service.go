package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sesudhanshu/Go_Microservice/internal/dberrors"
	"github.com/sesudhanshu/Go_Microservice/internal/models"
)

func (s *EchoServer) GetService(ctx echo.Context) error {
	services, err := s.DB.GetService(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, services)
}

func (s *EchoServer) AddService(ctx echo.Context) error {
	service := new(models.Services)
	if err := ctx.Bind(service); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	service, err := s.DB.AddService(ctx.Request().Context(), service)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, service)
}

func (s *EchoServer) GetServiceByID(ctx echo.Context) error {
	ID := ctx.Param("id")
	service, err := s.DB.GetServiceByID(ctx.Request().Context(), ID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, service)
}

func (s *EchoServer) UpdateService(ctx echo.Context) error {
	ID := ctx.Param("id")
	service := new(models.Services)
	if err := ctx.Bind(service); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	if ID != service.ServiceId {
		return ctx.JSON(http.StatusBadRequest, "id passed in param not matched")
	}
	service, err := s.DB.UpdateService(ctx.Request().Context(), service)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, service)

}

func (s *EchoServer) DeleteService(ctx echo.Context) error {
	ID := ctx.Param("id")
	err := s.DB.DeleteService(ctx.Request().Context(), ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusNoContent)
}
