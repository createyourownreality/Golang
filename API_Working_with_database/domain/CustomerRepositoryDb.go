package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	findAllSql := "select id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer

		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customer " + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) GetCustomerById(id int) (Customer, error) {
	findSql := "SELECT id, name, city, zipcode, date_of_birth, status FROM customers WHERE id = $1"

	rows := d.client.QueryRow(findSql, id)
	var c Customer
	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)

	if err != nil {
		log.Println("Error while scaning the customer table " + err.Error())
		return c, err
	}
	return c, nil

}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dsn := "user=postgres password=Meghu@2001 dbname=Customerdb sslmode=disable"
	client, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
