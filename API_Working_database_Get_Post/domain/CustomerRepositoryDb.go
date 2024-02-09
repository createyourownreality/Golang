package domain

import (
	"api/logger"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findAllSql := "SELECT id, name, city, zipcode, date_of_birth, status FROM customers"

	customers := make([]Customer, 0)
	if err := d.client.Select(&customers, findAllSql); err != nil {
		logger.Error("Error while querying customer table: " + err.Error())
		return nil, err
	}

	return customers, nil
}

func (d CustomerRepositoryDb) GetCustomerById(id int) (*Customer, error) {
	findSql := "SELECT id, name, city, zipcode, date_of_birth, status FROM customers WHERE id = $1"

	var c Customer
	if err := d.client.Get(&c, findSql, id); err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while scanning the customer table: " + err.Error())
			return nil, err
		}
		logger.Error("Error while scanning the customer table: " + err.Error())
		return nil, err
	}

	return &c, nil
}

func (d CustomerRepositoryDb) CreateCustomer(newCustomer Customer) (*Customer, error) {
	createSql := `INSERT INTO customers (id, name, city, zipcode, date_of_birth, status)
        VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, city, zipcode, date_of_birth, status`

	var createdCustomer Customer
	if err := d.client.QueryRowx(createSql, newCustomer.Id, newCustomer.Name, newCustomer.City, newCustomer.Zipcode, newCustomer.DateofBirth, newCustomer.Status).
		StructScan(&createdCustomer); err != nil {
		logger.Error("Error while creating customer:")
		return nil, err
	}

	logger.Info("Customer created successfully")
	return &createdCustomer, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dsn := "user=postgres password=Meghu@2001 dbname=Customerdb sslmode=disable"
	client, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
