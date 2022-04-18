package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"strings"
	"template/app/internal/apperror"
	template2 "template/app/internal/template"
	"template/app/pkg/client/postgresql"
	"template/app/pkg/logging"
	"time"
)

type db struct {
	client postgresql.Client
	logger logging.Logger
}

func NewStorage(c postgresql.Client, l logging.Logger) template2.Storage {
	return &db{
		client: c,
		logger: l,
	}
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", "")
}

func (r *db) Create(ctx context.Context, user *template2.User) error {
	q := `
		INSERT INTO "users" ("username", "password", "email") 
		VALUES ($1, $2, $3) 
		RETURNING id
	`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	r.logger.Trace(fmt.Sprintf("SQL query: %s", formatQuery(q)))
	if err := r.client.QueryRow(ctx, q, user.Username, user.Password, user.Email).Scan(&user.ID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			r.logger.Error(
				"SQL Error: ", pgErr.Message,
				", Detail: ", pgErr.Detail,
				", Where: ", pgErr.Where,
				", Code: ", pgErr.Code,
				", SQLState: ", pgErr.SQLState(),
			)
			if strings.Contains(pgErr.Message, "duplicate key value violates unique") {
				return apperror.NewAppError(
					"That email is already registered in the system",
					"Y-000004",
					"enter another email",
				)
			}
			return err
		}
		return err
	}
	return nil
}
