package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	connParams = os.Getenv("DB_CONN_PARAMS")
)

func init() {

	c, cErr := Connection()
	if cErr != nil {
		log.Fatal(
			cErr,
			"creation of the db connection failed",
		)
	}
	defer c.Close()

	_, t1Err := c.Exec(
		`
		CREATE TABLE IF NOT EXISTS post (
			id          VARCHAR(100)  NOT NULL PRIMARY KEY,
			title       VARCHAR(100)  NOT NULL,
			author      VARCHAR(100)  NOT NULL,
			content     VARCHAR(2000) NOT NULL,
			comments_on BOOLEAN       NOT NULL
		)
		`,
	)
	if t1Err != nil {
		log.Fatalf(
			"creation   table   post    failed, %s", t1Err.Error(),
		)
	}

	_, t2Err := c.Exec(
		`
		CREATE TABLE IF NOT EXISTS comment (
			id          VARCHAR(100)  NOT NULL PRIMARY KEY,
			post_id     VARCHAR(100)  NOT NULL,
			parent_id   VARCHAR(100),
			author      VARCHAR(100)  NOT NULL,
			content     VARCHAR(2000) NOT NULL,
			created_at  INTEGER       NOT NULL,

			FOREIGN KEY (post_id) REFERENCES post(id),
			FOREIGN KEY (parent_id) REFERENCES comment(id)
		)
		`,
	)
	if t2Err != nil {
		log.Fatalf(
			"creation  table   comment  failed, %s", t2Err.Error(),
		)
	}

}

func Connection() (*sql.DB, error) {

	c, cErr := sql.Open("postgres", connParams)
	if cErr != nil {
		return nil, cErr
	}

	if pErr := c.Ping(); pErr != nil {
		return nil, pErr
	}

	return c, nil

}
