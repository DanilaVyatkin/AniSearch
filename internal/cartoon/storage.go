package cartoon

import "context"

type Storage interface {
	Create(ctx context.Context, cartoon *Cartoon) error
	FindAll(ctx context.Context) (c []Cartoon, err error)
	FindOne(ctx context.Context, id string) (Cartoon, error)
	Update(ctx context.Context, cartoon *Cartoon) error
	Delete(ctx context.Context, id string) error
}
