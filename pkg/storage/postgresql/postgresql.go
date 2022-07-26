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

func NewPostgresStorage(user, password, host, port, dbName string) *PostgresStorage {
	dbinfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)
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

func (s *PostgresStorage) CreateShortLink(link models.Link) (string, error) {
	query := fmt.Sprintf("INSERT INTO linksdb(long_url, short_url) VALUES($1, $2);", link.Long, link.Short)
	_, err := s.db.Exec(query)
	if err != nil {
		return "", err
	}

	return link.Short, nil
}

func (s *PostgresStorage) GetLongLink(short string) (string, error) {
	query := fmt.Sprintf(`select * from linksDB where short_url='%s'`, short)
	rows, err := s.db.Query(query)
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
	}

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
