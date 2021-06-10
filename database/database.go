package database

import (
	"database/sql"
	"fmt"

	"github.com/SomeshSunariwal/GraphQL_implementation/config"
	"github.com/SomeshSunariwal/GraphQL_implementation/modal"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type Database struct {
	client *sql.DB
}

func DB_INIT() *sql.DB {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode='disable'",
		config.HOST, config.PG_PORT, config.PG_USER, config.PG_PASSWORD, config.PG_DATABASE_NAME)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Info("ERROR : ", err)
		return nil
	}

	return db
}

func Client() Database {
	return Database{
		client: DB_INIT(),
	}
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
		return book, err
	}
	defer database.client.Close()
	return book, nil
}

func (database *Database) UpdateItem() (string, error) {
	return "", nil
}

func (database *Database) DeleteItem() (string, error) {
	return "", nil
}

func (database *Database) GetItems() ([]modal.BookModal, error) {
	books := []modal.BookModal{}
	query := "SELECT id, book_name, author, seller, available from books"

	rows, err := database.client.Query(query)
	fmt.Println("..", err)
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
		availabilityZoneQuery := "SELECT books.id, address FROM availability_zone JOIN books on books.book_name = availability_zone.book_name WHERE books.book_name=$1"
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

func (database *Database) GetItemByID() (string, error) {
	return "", nil
}
