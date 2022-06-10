package db

import (
	"MakeAnAPI/internal/cartoon"
	"MakeAnAPI/pkg/client/postgesql"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"log"
)

type db struct {
	client postgesql.Client
}

func (d *db) Create(ctx context.Context, cartoon *cartoon.Cartoon) error {
	q := `
		INSERT INTO "AniSearch".public.anime (name, genre, rating, description)
		VALUES ($1, $2, $3, $4) 
		RETURNING id
		`

	if err := d.client.QueryRow(ctx, q, cartoon.Name, cartoon.Rating, cartoon.Genre, cartoon.Description).
		Scan(&cartoon.ID); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			fmt.Println(fmt.Sprintf("SQL error: %s, Detail: %s, Where: %s, Code: %s, SQLstate: %s", pgErr.Message,
				pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			return nil
		}
		return err
	}
	return nil
}

func (d *db) FindAll(ctx context.Context) (c []cartoon.Cartoon, err error) {
	q := `SELECT id, name, genre, rating, description FROM "AniSearch".public.anime`

	rows, err := d.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	name := make([]cartoon.Cartoon, 0)
	for rows.Next() {
		var cart cartoon.Cartoon

		err = rows.Scan(&cart.ID, &cart.Name, &cart.Genre, &cart.Rating, &cart.Description)
		if err != nil {
			return nil, err
		}

		name = append(name, cart)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return name, nil
}

func (d *db) FindOne(ctx context.Context, id string) (cartoon.Cartoon, error) {
	q := `SELECT id, name FROM "AniSearch".public.anime WHERE id = $1`

	var cart cartoon.Cartoon
	err := d.client.QueryRow(ctx, q, id).Scan(&cart.ID, &cart.Name)
	if err != nil {
		return cartoon.Cartoon{}, err
	}

	return cart, nil
}

func (d *db) Update(ctx context.Context, cartoon *cartoon.Cartoon) error {
	sqlStatement := `
			UPDATE "AniSearch".public.anime
			SET name = $2, genre = $3, rating = $4, description = $5
			WHERE id = $1
			`

	_, err := d.client.Exec(ctx, sqlStatement, cartoon.ID, cartoon.Name, cartoon.Rating, cartoon.Genre,
		cartoon.Description)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			fmt.Println(fmt.Sprintf("SQL error: %s, Detail: %s, Where: %s, Code: %s, SQLstate: %s", pgErr.Message,
				pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			return nil
		}
		return err
	}

	return nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	sqlStatement := `
			DELETE FROM "AniSearch".public.anime WHERE id = $1
			`

	_, err := d.client.Exec(ctx, sqlStatement, id)
	if err != nil {
		log.Fatalf("failed delete element, Error: %s", err)
	}

	return nil
}

func NewDB(client postgesql.Client) cartoon.Storage {
	return &db{
		client: client,
	}
}
