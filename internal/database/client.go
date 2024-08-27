package database

import (
	"context"
	"fmt"
	"time"

	"github.com/sesudhanshu/Go_Microservice/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DatabaseClient interface {
	Ready() bool
	GetAllCustomer(ctx context.Context, emailAddress string) ([]models.Customer, error)
	AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	GetCustomerByID(ctx context.Context, ID string) (*models.Customer, error)
	UpdateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	DeleteCustomer(ctx context.Context, ID string) error

	GetProduct(ctx context.Context, vendorId string) ([]models.Products, error)
	AddProduct(ctx context.Context, product *models.Products) (*models.Products, error)
	GetProductByID(ctx context.Context, ID string) (*models.Products, error)
	UpdateProduct(ctx context.Context, product *models.Products) (*models.Products, error)
	DeleteProduct(ctx context.Context, ID string) error

	GetService(ctx context.Context) ([]models.Services, error)
	AddService(ctx context.Context, service *models.Services) (*models.Services, error)
	GetServiceByID(ctx context.Context, ID string) (*models.Services, error)
	UpdateService(ctx context.Context, service *models.Services) (*models.Services, error)
	DeleteService(ctx context.Context, ID string) error

	GetAllVendors(ctx context.Context) ([]models.Vendor, error)
	AddVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error)
	GetVendorByID(ctx context.Context, ID string) (*models.Vendor, error)
	UpdateVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error)
	DeleteVendor(ctx context.Context, ID string) error
}

type Client struct {
	DB *gorm.DB
}

func (c Client) Ready() bool {
	var ready string
	txn := c.DB.Raw("SELECT 1 as ready").Row().Scan(&ready)
	if txn != nil {
		fmt.Println("Fail to connect to DB")
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}

func NewDatabaseClient() (DatabaseClient, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		"localhost",
		"postgres",
		"postgres",
		"postgres",
		5432,
		"disable")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "wisdom.",
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}
	client := Client{
		DB: db,
	}
	return client, nil
}
