package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

//for this interface we create this code to add DB stack
type DBLogger struct{}

func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

// function to receive a pointer to pg_option
// we need to have a connection with database
func NewReceive(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)

	return db
}
