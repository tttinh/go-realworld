package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tinhtt/go-realworld/internal/adapters/postgres/sqlc"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type Users struct {
	*sqlc.Queries
}

func NewUsers(db *pgx.Conn) *Users {
	return &Users{
		Queries: sqlc.New(db),
	}
}

func (r *Users) GetByID(ctx context.Context, id int) (domain.User, error) {
	u, err := r.GetUserByID(ctx, int64(id))
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:       int(u.ID),
		Name:     u.Username,
		Email:    u.Email,
		Password: u.Password,
		Bio:      u.Bio.String,
		Image:    u.Image.String,
	}, nil
}

func (r *Users) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := r.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:       int(u.ID),
		Name:     u.Username,
		Email:    u.Email,
		Password: u.Password,
		Bio:      u.Bio.String,
		Image:    u.Image.String,
	}, nil
}

func (r *Users) Add(ctx context.Context, u domain.User) (domain.User, error) {
	param := sqlc.InsertUserParams{
		Username: u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
	dbUser, err := r.InsertUser(ctx, param)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:       int(dbUser.ID),
		Name:     dbUser.Username,
		Email:    dbUser.Email,
		Password: dbUser.Password,
		Bio:      dbUser.Bio.String,
		Image:    dbUser.Image.String,
	}, nil
}

func (r *Users) Edit(ctx context.Context, u domain.User) (domain.User, error) {
	param := sqlc.UpdateUserParams{
		ID:       int64(u.ID),
		Username: pgtype.Text{String: u.Name, Valid: len(u.Name) > 0},
		Email:    pgtype.Text{String: u.Email, Valid: len(u.Email) > 0},
		Password: pgtype.Text{String: u.Password, Valid: len(u.Password) > 0},
		Bio:      pgtype.Text{String: u.Bio, Valid: len(u.Bio) > 0},
		Image:    pgtype.Text{String: u.Image, Valid: len(u.Image) > 0},
	}
	dbUser, err := r.UpdateUser(ctx, param)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:       int(dbUser.ID),
		Name:     dbUser.Username,
		Email:    dbUser.Email,
		Password: dbUser.Password,
		Bio:      dbUser.Bio.String,
		Image:    dbUser.Image.String,
	}, nil
}

func (r *Users) GetProfile(ctx context.Context, followerID int, followingUsername string) (domain.Profile, error) {
	p, err := r.GetProfileByName(ctx, sqlc.GetProfileByNameParams{
		FollowerID:    int64(followerID),
		FollowingName: followingUsername,
	})
	if err != nil {
		return domain.Profile{}, err
	}

	return domain.Profile{
		ID:        int(p.ID),
		Name:      p.Username,
		Bio:       p.Bio.String,
		Image:     p.Image.String,
		Following: p.Following,
	}, nil
}

func (r *Users) Follow(ctx context.Context, followerID int, followingID int) error {
	return r.InsertFollow(ctx, sqlc.InsertFollowParams{
		FollowerID:  int64(followerID),
		FollowingID: int64(followingID),
	})
}

func (r *Users) Unfollow(ctx context.Context, followerID int, followingID int) error {
	return r.DeleteFollow(ctx, sqlc.DeleteFollowParams{
		FollowerID:  int64(followerID),
		FollowingID: int64(followingID),
	})
}
