package users

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dikopylov/highload-architect/internal/model/types"
	"github.com/jmoiron/sqlx"
)

var ErrUserNotFound = errors.New("user not found")

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, id types.UserID) (*User, error)
}

type pgsqlRepository struct {
	db sqlx.DB
}

func NewPgsqlRepository(db sqlx.DB) Repository {
	return &pgsqlRepository{db: db}
}

func (r *pgsqlRepository) CreateUser(ctx context.Context, user *User) error {
	var userUUID types.UserID
	const query = `
insert into users (first_name, second_name, birthdate, biography, city, password, age) 
values ($1, $2, $3, $4, $5, $6, $7)
returning id
`
	err := r.db.GetContext(
		ctx,
		&userUUID,
		query,
		user.FirstName,
		user.SecondName,
		user.Birthdate,
		user.Biography,
		user.City,
		user.Password,
		user.Age,
	)
	if err != nil {
		return err
	}

	user.ID = userUUID

	return nil
}

func (r *pgsqlRepository) GetUserByID(ctx context.Context, id types.UserID) (*User, error) {
	const query = `
select id, first_name, second_name, birthdate, biography, city, password, age
from users 
where id = $1
`
	user := &User{}
	err := r.db.GetContext(ctx, user, query, id.String())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	if user.ID.IsNil() {
		return nil, ErrUserNotFound
	}

	return user, nil
}
