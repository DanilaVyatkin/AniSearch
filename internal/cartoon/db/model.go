package cartoon

import (
	"MakeAnAPI/internal/cartoon"
	"database/sql"
)

type Cartoon struct {
	ID          int64          `json:"id"`
	Name        sql.NullString `json:"name"`
	Genre       sql.NullString `json:"genre"`
	Rating      sql.NullString `json:"rating"`
	Description sql.NullString `json:"description"`
}

func (c Cartoon) ToDomain() cartoon.Cartoon {
	crt := cartoon.Cartoon{
		ID: c.ID,
	}
	if c.Name.Valid {
		crt.Name = string(c.Name.String)
	}
	if c.Genre.Valid {
		crt.Genre = string(c.Genre.String)
	}
	if c.Rating.Valid {
		crt.Rating = string(c.Genre.String)
	}
	if c.Description.Valid {
		crt.Description = string(c.Genre.String)
	}

	return crt
}
