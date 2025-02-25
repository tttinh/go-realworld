package domain

import "context"

type UserRepo interface {
	FindById(ctx context.Context, id int) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	Insert(ctx context.Context, u User) (User, error)
	Update(ctx context.Context, u User) (User, error)
}
