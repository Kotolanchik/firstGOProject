package user

import "context"

type Storage interface {
	FindUser(ctx context.Context, id string) (User, error)
	Create(ctx context.Context, user User) (string, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
}
