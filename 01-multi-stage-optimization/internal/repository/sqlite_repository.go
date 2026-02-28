package repository

import (
	"database/sql"
	"minimalrestapi/internal/entity"
)

type LogRepository struct {
	db *sql.DB
}

func NewLogRepository(db *sql.DB) *LogRepository {
	return &LogRepository{db: db}
}

func (r *LogRepository) Save(log *entity.Log) error {
	_, err := r.db.Exec("INSERT INTO logs (id, data) VALUES (?, ?)", log.ID, log.Data)
	return err
}

func (r *LogRepository) Bootstrap() error {
	_, err := r.db.Exec("CREATE TABLE IF NOT EXISTS logs (id TEXT, data TEXT)")
	return err
}
