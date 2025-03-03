package pgrepo

import (
	"context"

	"github.com/tinhtt/go-realworld/internal/domain"
)

func (r *Articles) GetAllComments(ctx context.Context, id int) ([]domain.Comment, error) {
	return nil, nil
}

func (r *Articles) AddComment(ctx context.Context, c domain.Comment) (domain.Comment, error) {
	return domain.Comment{}, nil
}

func (r *Articles) RemoveComment(ctx context.Context, id int) error {
	return nil
}
