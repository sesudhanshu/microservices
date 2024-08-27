package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sesudhanshu/Go_Microservice/internal/dberrors"
	"github.com/sesudhanshu/Go_Microservice/internal/models"
)

func (s *EchoServer) GetProduct(ctx echo.Context) error {
	vendorId := ctx.QueryParam("vendorId")
	product, err := s.DB.GetProduct(ctx.Request().Context(), vendorId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, product)

}

func (s *EchoServer) AddProduct(ctx echo.Context) error {
	product := new(models.Products)
	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	product, err := s.DB.AddProduct(ctx.Request().Context(), product)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, product)
}

func (s *EchoServer) GetProductByID(ctx echo.Context) error {
	ID := ctx.Param("id")
	product, err := s.DB.GetProductByID(ctx.Request().Context(), ID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, product)
}

func (s *EchoServer) UpdateProduct(ctx echo.Context) error {
	ID := ctx.Param("id")
	product := new(models.Products)
	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	if ID != product.ProductId {
		return ctx.JSON(http.StatusBadRequest, "id passed in param not matched")
	}
	product, err := s.DB.UpdateProduct(ctx.Request().Context(), product)
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
	return ctx.JSON(http.StatusOK, product)
}

func (s *EchoServer) DeleteProduct(ctx echo.Context) error {
	ID := ctx.Param("id")
	err := s.DB.DeleteProduct(ctx.Request().Context(), ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.NoContent(http.StatusResetContent)
}
