package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sajidcodesdotcom/kira/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, limit, offset int) ([]*models.User, error)
}

type PgUserRepository struct {
	db *pgxpool.Pool
}

func NewPgUserRepository(db *pgxpool.Pool) *PgUserRepository {
	return &PgUserRepository{db: db}
}

func (r *PgUserRepository) Create(ctx context.Context, user *models.User) error {
	query := `
	insert into users (id, email, username, full_name, password, avatar_url, role, created_at, updated_at)
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.db.Exec(ctx, query, user.ID, user.Email, user.Username, user.FullName, user.Password, user.AvatarURL, user.Role, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create user in db: %w", err)
	}
	return nil
}

func (r *PgUserRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query := `
	select id, email, username, full_name, password, avatar_url, role, created_at, updated_at
	from users
	where id = $1
	`
	user := &models.User{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Email, &user.Username, &user.FullName, &user.Password, &user.AvatarURL, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user record not found (finding by id): %w", err)
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return user, nil
}

func (r *PgUserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
	select id, email, username, full_name, password, avatar_url, role, created_at, updated_at
	from users
	where email = $1
	`

	user := &models.User{}
	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.Username, &user.FullName, &user.Password, &user.AvatarURL, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("no record found (finding by email): %w", err)
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return user, nil
}

func (r *PgUserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `
	select id, email, username, full_name, password, avatar_url, role, created_at, updated_at
	from users
	where username = $1
	`

	user := &models.User{}
	err := r.db.QueryRow(ctx, query, username).Scan(
		&user.ID, &user.Email, &user.Username, &user.FullName, &user.Password, &user.AvatarURL, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("no record found (finding by username): %w", err)
		}
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}

	return user, nil
}

func (r *PgUserRepository) Update(ctx context.Context, user *models.User) error {
	query := `
	update users
	set email=$2, username=$3, full_name=$4, password=$5, avatar_url=$6, role=$7, updated_at=Now()
	where id=$1
	`

	result, err := r.db.Exec(ctx, query, user.ID, user.Email, user.Username, user.FullName, user.Password, user.AvatarURL, user.Role, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("user not found when tried to update")
	}
	return nil
}

func (r *PgUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `
	delete from users where id=$1
	`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("user not found when tried to delete")
	}
	return nil
}

func (r *PgUserRepository) List(ctx context.Context, limit, offset int) ([]*models.User, error) {
	query := `
	select id, email, username, full_name, password, avatar_url, role, updated_at, created_at
	from users
	order by created_at desc
	limit $1 offset $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query when listing: %w", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID, &user.Email, &user.Username, &user.FullName, &user.Password, &user.AvatarURL, &user.Role, &user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("Error readinng row when listing: %w", err)
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating user rows: %w", err)
	}

	return users, nil
}
