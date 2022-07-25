package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"ozonTest/pkg/models"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() *PostgresStorage {
	dbinfo := "postgres://postgres:postgres@db:5432/postgres?sslmode=disable"
	//dbinfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", DB_USER, DB_PASSWORD, HOST, PORT, DB_NAME)
	fmt.Println(dbinfo)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &PostgresStorage{db: db}
}

func (storage *PostgresStorage) CreateShortLink(link models.Link) (string, error) {
	db := setupDB()

	_, err := db.Exec("INSERT INTO linksdb(long_url, short_url) VALUES($1, $2);", link.Long, link.Short)
	if err != nil {
		return "", err
	}

	return link.Short, nil
}

func (storage *PostgresStorage) GetLongLink(short string) (string, error) {
	db := setupDB()

	query := fmt.Sprintf(`select * from linksDB where short_url='%s'`, short)
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

	return "", fmt.Errorf("there is no long link for current short")
}

func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("postgres://postgres:postgres@db:5432/postgres?sslmode=disable")
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		panic(err)
	}
	return db
}
