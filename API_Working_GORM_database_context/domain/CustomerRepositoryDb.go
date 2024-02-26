package domain

import (
	"api/logger"
	"context"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomerRepositoryDb struct {
	client *gorm.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	customers := make([]Customer, 0)
	result := d.client.Find(&customers)
	if result.Error != nil {
		logger.Error("Error while querying customer table: " + result.Error.Error())
		return nil, result.Error
	}

	return customers, nil
}

func (d CustomerRepositoryDb) GetCustomerById(ctx context.Context, id int) (*Customer, error) {
	var c Customer
	result := d.client.First(&c, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			logger.Error("Customer not found: " + result.Error.Error())
			return nil, result.Error
		}
		logger.Error("Error while querying the customer table: " + result.Error.Error())
		return nil, result.Error
	}

	return &c, nil
}

func (d CustomerRepositoryDb) CreateCustomer(newCustomer Customer) (*Customer, error) {
	// createSql := `INSERT INTO customers (id, name, city, zipcode, date_of_birth, status)
	//     VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, city, zipcode, date_of_birth, status`

	result_new := d.client.Create(&newCustomer)

	if result_new.Error != nil {

		logger.Error("Error while creating the customer:")
		return nil, result_new.Error
	}

	logger.Info("Customer created successfully")
	return &newCustomer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dsn := "user=postgres password=Meghu@2001 dbname=Customerdb sslmode=disable"
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//Set connection pool settings
	sqlDB, err := client.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
