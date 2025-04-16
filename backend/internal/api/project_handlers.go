package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sajidcodesdotcom/kira/internal/models"
	"github.com/sajidcodesdotcom/kira/internal/services"
	"github.com/sajidcodesdotcom/kira/utils"
)

type ProjectHandler struct {
	projectRepo services.ProjectRepository
	validate    *validator.Validate
}

func NewProjectHandler(projectRepo services.ProjectRepository, validator *validator.Validate) *ProjectHandler {
	return &ProjectHandler{projectRepo: projectRepo, validate: validator}
}

func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	var projectData struct {
		Name        string `json:"name" validate:"required,min=2,max=100"`
		Description string `json:"description" validate:"required,min=10,max=500"`
		Status      string `json:"status" validate:"required,oneof=active inactive"`
	}
	if err := json.NewDecoder(r.Body).Decode(&projectData); err != nil {
		http.Error(w, "Failed to create project (not able to read body)", http.StatusInternalServerError)
		return
	}
	if err := h.validate.Struct(projectData); err != nil {
		utils.RespondWithError(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		return
	}
	ownerID, ok := r.Context().Value("user_id").(uuid.UUID)
	if !ok {
		utils.RespondWithError(w, "Failed to get user ID from context", http.StatusInternalServerError)
		return
	}
	err := h.projectRepo.Create(ctx, &models.Project{
		ID:          uuid.New(),
		Name:        projectData.Name,
		Description: projectData.Description,
		Status:      projectData.Status,
		OwnerID:     ownerID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		utils.RespondWithError(w, "Failed to create project: "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, "Project created successfully", http.StatusCreated)
}

func (h *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	var projectData struct {
		ID          uuid.UUID `json:"id" validate:"required"`
		Name        string    `json:"name" validate:"required,min=2,max=100"`
		Description string    `json:"description" validate:"required,min=10,max=500"`
		Status      string    `json:"status" validate:"required,oneof=active inactive"`
	}
	if err := json.NewDecoder(r.Body).Decode(&projectData); err != nil {
		http.Error(w, "Failed to update project (not able to read body)", http.StatusInternalServerError)
		return
	}
	if err := h.validate.Struct(projectData); err != nil {
		utils.RespondWithError(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		return
	}
	ownerID, ok := r.Context().Value("user_id").(uuid.UUID)
	if !ok {
		utils.RespondWithError(w, "Failed to get user ID from context", http.StatusInternalServerError)
		return
	}
	err := h.projectRepo.Update(ctx, &models.Project{
		ID:          projectData.ID,
		Name:        projectData.Name,
		Description: projectData.Description,
		OwnerID:     ownerID,
		Status:      projectData.Status,
	})
	if err != nil {
		utils.RespondWithError(w, "Failed to update project: "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, "Project updated successfully", http.StatusOK)
}

func (h *ProjectHandler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.RespondWithError(w, "Project ID is required", http.StatusBadRequest)
		return
	}
	projectID, err := uuid.Parse(id)
	if err != nil {
		utils.RespondWithError(w, "Invalid project ID format", http.StatusBadRequest)
		return
	}
	project, err := h.projectRepo.GetByID(ctx, projectID)
	if err != nil {
		utils.RespondWithError(w, "Failed to get project: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if project == nil {
		utils.RespondWithError(w, "Project not found", http.StatusNotFound)
		return
	}
	utils.RespondWithJSON(w, project, http.StatusOK)
}

func (h *ProjectHandler) GetProjectsByOwner(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	ownerID, ok := r.Context().Value("user_id").(uuid.UUID)
	if !ok {
		utils.RespondWithError(w, "Failed to get user ID from context", http.StatusInternalServerError)
		return
	}
	projects, err := h.projectRepo.GetByOwner(ctx, ownerID)
	if err != nil {
		utils.RespondWithError(w, "Failed to get projects: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if projects == nil {
		utils.RespondWithError(w, "No projects found for this user", http.StatusNotFound)
		return
	}
	utils.RespondWithJSON(w, projects, http.StatusOK)
}

func (h *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.RespondWithError(w, "Project ID is required", http.StatusBadRequest)
		return
	}
	projectID, err := uuid.Parse(id)
	if err != nil {
		utils.RespondWithError(w, "Invalid project ID format", http.StatusBadRequest)
		return
	}
	err = h.projectRepo.Delete(ctx, projectID)
	if err != nil {
		utils.RespondWithError(w, "Failed to delete project: "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RespondWithJSON(w, "Project deleted successfully", http.StatusOK)
}

func (h *ProjectHandler) ListProjects(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	limit := 10
	offset := 0

	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit <= 0 {
			utils.RespondWithError(w, "Invalid limit value", http.StatusBadRequest)
			return
		}
	}
	if offsetStr != "" {
		var err error
		offset, err = strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			utils.RespondWithError(w, "Invalid offset value", http.StatusBadRequest)
			return
		}
	}

	projects, err := h.projectRepo.List(ctx, limit, offset)
	if err != nil {
		utils.RespondWithError(w, "Failed to list projects: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if projects == nil {
		utils.RespondWithError(w, "No projects found", http.StatusNotFound)
		return
	}
	utils.RespondWithJSON(w, projects, http.StatusOK)
}
