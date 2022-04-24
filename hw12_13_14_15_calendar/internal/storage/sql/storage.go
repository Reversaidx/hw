package sqlstorage

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type StorageInterface struct { // TODO
	db *sql.DB
	//connect
}

type Storage struct { // TODO
	db        *sql.DB
	conString string
	//connect
}

func New(conString string) *Storage {

	return &Storage{
		db:        nil,
		conString: conString,
	}
}

func (s *Storage) Connect(ctx context.Context) (err error) {
	s.db, err = sql.Open("pgx", s.conString)
	if err != nil {
		return fmt.Errorf("cannot open pgx driver: %w", err)
	}
	return s.db.PingContext(ctx)

}

func (s *Storage) Close() error {
	// TODO
	s.db.Close()
	return nil
}
