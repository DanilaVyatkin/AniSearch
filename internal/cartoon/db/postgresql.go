package cartoon

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

func NewDB(client postgesql.Client) cartoon.Storage {
	return &db{
		client: client,
	}
}

const (
	sqlStatementCreate = `INSERT INTO "AniSearch".public.anime (name, genre, rating, description)
						VALUES ($1, $2, $3, $4) 
						RETURNING id`
	sqlStatementFindAll = `SELECT id, name, genre, rating, description FROM "AniSearch".public.anime`
	sqlStatementFindOne = `SELECT id, name FROM "AniSearch".public.anime WHERE id = $1`
	sqlStatementUpdate  = `UPDATE "AniSearch".public.anime
						SET name = $2, genre = $3, rating = $4, description = $5
						WHERE id = $1`
	sqlStatementDelete = `DELETE FROM "AniSearch".public.anime WHERE id = $1`
)

func (d *db) Create(ctx context.Context, cartoon *cartoon.Cartoon) error {
	if err := d.client.QueryRow(ctx, sqlStatementCreate, cartoon.Name, cartoon.Rating, cartoon.Genre, cartoon.Description).
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
	rows, err := d.client.Query(ctx, sqlStatementFindAll)
	if err != nil {
		return nil, err
	}

	cartoons := make([]cartoon.Cartoon, 0)
	for rows.Next() {
		var cart Cartoon

		err = rows.Scan(&cart.ID, &cart.Name, &cart.Genre, &cart.Rating, &cart.Description)
		if err != nil {
			return nil, err
		}

		cartoons = append(cartoons, cart.ToDomain())
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return cartoons, nil
}

func (d *db) FindOne(ctx context.Context, id string) (cartoon.Cartoon, error) {
	var cart cartoon.Cartoon
	err := d.client.QueryRow(ctx, sqlStatementFindOne, id).Scan(&cart.ID, &cart.Name)
	if err != nil {
		return cartoon.Cartoon{}, err
	}

	return cart, nil
}

func (d *db) Update(ctx context.Context, cartoon *cartoon.Cartoon) error {
	_, err := d.client.Exec(ctx, sqlStatementUpdate, cartoon.ID, cartoon.Name, cartoon.Rating, cartoon.Genre,
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
	_, err := d.client.Exec(ctx, sqlStatementDelete, id)
	if err != nil {
		log.Fatalf("failed delete element, Error: %s", err)
	}

	return nil
}
