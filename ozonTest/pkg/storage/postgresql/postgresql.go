package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"ozonTest/pkg/models"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "qweryGood"
	DB_NAME     = "postgres"
)

const (
	PORT = 5432
	HOST = "database"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() *PostgresStorage {
	dbinfo := fmt.Sprintf("host=%s, port=%d, user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	return &PostgresStorage{db: db}
}

func (storage *PostgresStorage) CreateShortLink(link models.Link) (string, error) {
	db := setupDB()
	defer db.Close()

	var lastInsertID int
	err := db.QueryRow("INSERT INTO linksDB(long_link, short_link) VALUES($1, $2);", link.Long, link.Short).Scan(&lastInsertID)

	fmt.Println(lastInsertID)
	if err == nil {
		return "", fmt.Errorf("что-то пошло не так при занесении новой ссылки в postgres")
	}

	return link.Short, nil
}

func (storage *PostgresStorage) GetLongLink(short string) (string, error) {
	db := setupDB()
	defer db.Close()

	query := fmt.Sprintf(`select * from linksDB where short_link='%s'`, short)
	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var long string
		var short string
		err = rows.Scan(&id, &long, &short)
		fmt.Println(long, short)
		return long, nil
	}

	return "", fmt.Errorf("для заданной короткой ссылки нет длинной")
}

func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		panic(err)
	}
	return db
}
