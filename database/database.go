package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/SomeshSunariwal/GraphQL_implementation/config"
	"github.com/SomeshSunariwal/GraphQL_implementation/modal"
	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Database struct {
	client *sql.DB
}

type Config struct {
	Config config.ENV_CONFIG
}

func DB_INIT() *sql.DB {
	//GETTING Config Details from Environment Variable
	databaseConfig := DatabaseConfig()

	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode='disable'",
		databaseConfig.HOST, 5432, databaseConfig.PG_USER, databaseConfig.PG_PASSWORD, databaseConfig.PG_DATABASE_NAME)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Info("ERROR : ", err)
		return nil
	}

	return db
}

func DatabaseConfig() *config.ENV_CONFIG {
	HOST := os.Getenv("HOST")
	PG_USER := os.Getenv("PG_USER")
	PG_PASSWORD := os.Getenv("PG_PASSWORD")
	PG_DATABASE_NAME := os.Getenv("PG_DATABASE_NAME")
	if HOST == "" || PG_USER == "" || PG_PASSWORD == "" || PG_DATABASE_NAME == "" {
		return &config.ENV_CONFIG{
			HOST:             config.HOST,
			PG_USER:          config.PG_USER,
			PG_PASSWORD:      config.PG_PASSWORD,
			PG_DATABASE_NAME: config.PG_DATABASE_NAME,
		}
	}
	return &config.ENV_CONFIG{
		HOST:             HOST,
		PG_USER:          PG_USER,
		PG_PASSWORD:      PG_PASSWORD,
		PG_DATABASE_NAME: PG_DATABASE_NAME,
	}
}

func Client() Database {
	return Database{
		client: DB_INIT(),
	}
}

func (database *Database) Health() error {
	err := database.client.Ping()
	if err != nil {
		return errors.New("DB has Error")
	}
	return nil
}

func (database *Database) AddItem(userRequest *modal.PostBook) (modal.BookModal, error) {
	book := modal.BookModal{}
	details := modal.DetailsModal{}
	availabilities := []modal.AvailabilityModal{}
	availability := modal.AvailabilityModal{}
	query := `WITH insert_into_books as ( insert into books (book_name, author, seller, available)
											VALUES ($1, $2, $3, $4)
						   					RETURNING id, book_name, author, seller, available), 
			insert_into_availability_zone as ( insert into availability_zone(book_name, address)
											VALUES ($1, $5)
											RETURNING book_name, address)
			SELECT ib.id, ib.book_name, author, seller, available, address from insert_into_books ib
			JOIN insert_into_availability_zone iaz ON iaz.book_name = ib.book_name`

	err := database.client.QueryRow(query, userRequest.BookName, userRequest.Author, userRequest.Seller, userRequest.Available, userRequest.Location).Scan(
		&book.ID,
		&book.BookName,
		&details.Author,
		&details.Seller,
		&book.Available,
		&availability.Location,
	)

	availabilities = append(availabilities, availability)
	book.Details = details
	book.Availability = availabilities

	if err != nil {
		log.Info("Error : ", err)
		return book, err
	}
	defer database.client.Close()
	return book, nil
}

func (database *Database) UpdateItem(updateValues *modal.PostBook) (modal.BookModal, error) {
	book := modal.BookModal{}
	details := modal.DetailsModal{}
	query := `UPDATE books SET author=COALESCE($1 , author),
								available=COALESCE($2, available),
								seller=COALESCE($3, seller) WHERE book_name=$4
								RETURNING id, book_name, author, seller, available`

	err := database.client.QueryRow(query, updateValues.Author, updateValues.Available, updateValues.Seller, updateValues.BookName).Scan(
		&book.ID,
		&book.BookName,
		&details.Author,
		&details.Seller,
		&book.Available,
	)

	book.Details = details

	if err != nil {
		log.Info("Error : ", err)
		return book, err
	}
	defer database.client.Close()
	return book, nil
}

func (database *Database) DeleteItem(bookName string) (modal.BookModal, error) {
	book := modal.BookModal{}
	details := modal.DetailsModal{}
	var stringsArr []sql.NullString

	query := `WITH delete_books as ( DELETE FROM books WHERE book_name=$1 RETURNING id, book_name, author, seller, available), 
					delete_availability_zone as ( DELETE FROM availability_zone WHERE book_name= $1 RETURNING book_name, address)
						SELECT ib.id, ib.book_name, author, seller, available, array_agg(address) FROM delete_books ib
						JOIN delete_availability_zone iaz ON iaz.book_name = ib.book_name
						GROUP BY ib.book_name, ib.id, ib.author, ib.seller, ib.available`

	err := database.client.QueryRow(query, bookName).Scan(
		&book.ID,
		&book.BookName,
		&details.Author,
		&details.Seller,
		&book.Available,
		pq.Array(&stringsArr),
	)

	availabilityZones := []modal.AvailabilityModal{}
	for _, value := range stringsArr {
		availabilityZone := modal.AvailabilityModal{}
		value, _ := value.Value()
		availabilityZone.Location = value.(string)
		availabilityZones = append(availabilityZones, availabilityZone)
	}

	book.Availability = availabilityZones
	book.Details = details

	if err != nil {
		log.Info("Error : ", err)
		return book, nil
	}
	return book, nil
}

func (database *Database) GetItems() ([]modal.BookModal, error) {
	books := []modal.BookModal{}
	query := "SELECT id, book_name, author, seller, available from books"

	rows, err := database.client.Query(query)

	if err != nil {
		log.Info("Error : ", err)
		return nil, err
	}
	defer database.client.Close()

	for rows.Next() {
		book := modal.BookModal{}
		availabilityZones := []modal.AvailabilityModal{}
		details := modal.DetailsModal{}
		rows.Scan(
			&book.ID,
			&book.BookName,
			&details.Author,
			&details.Seller,
			&book.Available,
		)
		book.Details = details
		availabilityZoneQuery := "SELECT books.id, address FROM availability_zone JOIN books ON books.book_name = availability_zone.book_name WHERE books.book_name=$1"
		rows, err := database.client.Query(availabilityZoneQuery, book.BookName)
		if err != nil {
			log.Info("Error : ", err)
			return nil, err
		}
		for rows.Next() {
			availabilityZone := modal.AvailabilityModal{}
			rows.Scan(
				&availabilityZone.ID,
				&availabilityZone.Location,
			)
			availabilityZones = append(availabilityZones, availabilityZone)
		}
		book.Availability = availabilityZones
		books = append(books, book)
	}

	return books, nil
}

func (database *Database) GetItemByID(bookName string) (modal.BookModal, error) {
	book := modal.BookModal{}
	details := modal.DetailsModal{}
	var stringsArr []sql.NullString

	query := `SELECT books.id, books.book_name, author, seller, available, array_agg(address) FROM books 
						JOIN availability_zone on availability_zone.book_name = books.book_name 
						GROUP By books.id, books.book_name 
						HAVING books.book_name=$1`

	err := database.client.QueryRow(query, bookName).Scan(
		&book.ID,
		&book.BookName,
		&details.Author,
		&details.Seller,
		&book.Available,
		pq.Array(&stringsArr),
	)

	defer database.client.Close()

	availabilityZones := []modal.AvailabilityModal{}
	for _, value := range stringsArr {
		availabilityZone := modal.AvailabilityModal{}
		value, _ := value.Value()
		availabilityZone.Location = value.(string)
		availabilityZones = append(availabilityZones, availabilityZone)
	}

	book.Availability = availabilityZones
	book.Details = details

	if err != nil {
		log.Info("Error : ", err)
		return book, err
	}
	return book, nil
}
