package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	AvatarURL string    `json:"avatar_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(ID uuid.UUID, fullName, email, password, username, role, avatarURL string) *User {
	now := time.Now()
	return &User{
		ID:        ID,
		FullName:  fullName,
		Email:     email,
		Password:  password,
		Username:  username,
		AvatarURL: avatarURL,
		Role:      role,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

type Project struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerID     uuid.UUID `json:"owner_id"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewProject(id uuid.UUID, name, description string, ownerId uuid.UUID, status string) *Project {
	now := time.Now()
	return &Project{
		ID:          id,
		Name:        name,
		Description: description,
		OwnerID:     ownerId,
		Status:      status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ProjectID   uuid.UUID  `json:"project_id"`
	AssigneeID  *uuid.UUID `json:"assignee_id"`
	Status      string     `json:"status"`
	Priority    string     `json:"priority"`
	DueDate     *time.Time `json:"due_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func NewTask(id uuid.UUID, title, description string, productId uuid.UUID, assigneeId *uuid.UUID, status, priority string, dueDate *time.Time) *Task {
	now := time.Now()
	return &Task{
		ID:          id,
		Title:       title,
		Description: description,
		ProjectID:   productId,
		AssigneeID:  assigneeId,
		Status:      status,
		Priority:    priority,
		DueDate:     dueDate,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
