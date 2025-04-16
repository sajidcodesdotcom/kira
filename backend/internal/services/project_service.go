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

type ProjectRepository interface {
	Create(ctx context.Context, project *models.Project) error
	Update(ctx context.Context, project *models.Project) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Project, error)
	GetByOwner(ctx context.Context, ownerID uuid.UUID) (*models.Project, error)
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, limit, offset int) ([]*models.Project, error)
}

type PgProjectRepository struct {
	db *pgxpool.Pool
}

func NewPgProjectRepository(db *pgxpool.Pool) *PgProjectRepository {
	return &PgProjectRepository{
		db: db,
	}
}

func (r *PgProjectRepository) Create(ctx context.Context, project *models.Project) error {
	query := `
	insert into projects (id, name, description, status, owner_id, created_at, updated_at)
	values($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Exec(context.Background(), query, project.ID, project.Name, project.Description, project.Status, project.OwnerID, project.CreatedAt, project.UpdatedAt)
	if err != nil {
		return fmt.Errorf("Failed to create project in db: %w", err)
	}
	return nil
}

func (r *PgProjectRepository) Update(ctx context.Context, project *models.Project) error {
	query := `
	update projects set name = $1, description = $2, status = $3, owner_id = $4, updated_at = $5
	where id = $6
	`
	_, err := r.db.Exec(ctx, query, project.Name, project.Description, project.Status, project.OwnerID, project.UpdatedAt, project.ID)
	if err != nil {
		return fmt.Errorf("failed to update project in db: %w", err)
	}
	return nil
}

func (r *PgProjectRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Project, error) {
	query := `
	select id, name, description, status, owner_id, created_at, updated_at
	from projects
	where id = $1
	`
	project := &models.Project{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&project.ID, &project.Name, &project.Description, &project.Status, &project.OwnerID, &project.CreatedAt, &project.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("project record not found (finding by id): %w", err)
		}
		return nil, fmt.Errorf("failed to get project by id: %w", err)
	}
	return project, nil

}

func (r *PgProjectRepository) GetByOwner(ctx context.Context, ownerID uuid.UUID) (*models.Project, error) {
	query := `
	select id, name, description, status, owner_id, created_at, updated_at
	from projects
	where owner_id = $1
	`
	project := &models.Project{}
	err := r.db.QueryRow(ctx, query, ownerID).Scan(
		&project.ID, &project.Name, &project.Description, &project.Status, &project.OwnerID, &project.CreatedAt, &project.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("project record not found (finding by owner id): %w", err)
		}
		return nil, fmt.Errorf("failed to get project by owner id: %w", err)
	}
	return project, nil
}

func (r *PgProjectRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `
	delete from projects
	where id = $1
	`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete project from db: %w", err)
	}
	return nil
}

func (r *PgProjectRepository) List(ctx context.Context, limit, offset int) ([]*models.Project, error) {
	query := `
	select id, name, description, status, owner_id, created_at, updated_at
	from projects
	limit $1 offset $2
	`
	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list projects: %w", err)
	}
	defer rows.Close()

	var projects []*models.Project
	for rows.Next() {
		project := &models.Project{}
		err := rows.Scan(
			&project.ID, &project.Name, &project.Description, &project.Status, &project.OwnerID, &project.CreatedAt, &project.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan project: %w", err)
		}
		projects = append(projects, project)
	}
	return projects, nil
}
