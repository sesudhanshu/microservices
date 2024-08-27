package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/sesudhanshu/Go_Microservice/internal/dberrors"
	"github.com/sesudhanshu/Go_Microservice/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (c Client) GetProduct(ctx context.Context, vendorId string) ([]models.Products, error) {
	var product []models.Products
	result := c.DB.WithContext(ctx).
		Where(models.Products{VendorId: vendorId}).
		Find(&product)
	return product, result.Error
}

func (c Client) AddProduct(ctx context.Context, product *models.Products) (*models.Products, error) {
	product.ProductId = uuid.NewString()
	result := c.DB.WithContext(ctx).
		Create(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return product, nil
}

func (c Client) GetProductByID(ctx context.Context, ID string) (*models.Products, error) {
	product := &models.Products{}
	result := c.DB.WithContext(ctx).
		Where(&models.Products{ProductId: ID}).
		First(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{}
		}
		return nil, result.Error
	}
	return product, nil
}

func (c Client) UpdateProduct(ctx context.Context, product *models.Products) (*models.Products, error) {
	var products []models.Products
	result := c.DB.WithContext(ctx).Model(&products).
		Clauses(clause.Returning{}).
		Where(&models.Products{ProductId: product.ProductId}).
		Updates(models.Products{
			Name:     product.Name,
			Price:    product.Price,
			VendorId: product.VendorId,
		})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, &dberrors.NotFoundError{Entity: "product", ID: product.ProductId}
	}
	return &products[0], nil
}

func (c Client) DeleteProduct(ctx context.Context, ID string) error {
	return c.DB.WithContext(ctx).Delete(&models.Products{ProductId: ID}).Error
}
