package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Zrossiz/Home_menu/backend/internal/dto"
	"github.com/Zrossiz/Home_menu/backend/internal/helpers"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type CategoryHandler struct {
	service CategoryService
	logger  *zap.Logger
}

type CategoryService interface {
	Create(payload dto.CreateCategoryDTO) error
	Delete(categoryID int) error
	GetAll() ([]dto.CategoryDTO, error)
	Update(categoryID int, payload dto.UpdateCategoryDTO) error
}

func NewCategoryHandler(service CategoryService, logger *zap.Logger) *CategoryHandler {
	return &CategoryHandler{
		service: service,
		logger:  logger,
	}
}

func (c *CategoryHandler) Create(rw http.ResponseWriter, r *http.Request) {
	var body dto.CreateCategoryDTO

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(rw, "invalid request body", http.StatusBadRequest)
		return
	}

	err = c.service.Create(body)
	if err != nil {
		c.logger.Error("error create category", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

func (c *CategoryHandler) Update(rw http.ResponseWriter, r *http.Request) {
	categoryStringID := chi.URLParam(r, "categoryID")
	categoryIntID, err := strconv.Atoi(categoryStringID)
	if err != nil {
		http.Error(rw, "invalid category id", http.StatusBadRequest)
		return
	}

	var body dto.UpdateCategoryDTO
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(rw, "invalid body", http.StatusBadRequest)
		return
	}

	err = c.service.Update(categoryIntID, body)
	if err != nil {
		c.logger.Error("error update category", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (c *CategoryHandler) Delete(rw http.ResponseWriter, r *http.Request) {
	categoryStringID := chi.URLParam(r, "categoryID")
	categoryIntID, err := strconv.Atoi(categoryStringID)
	if err != nil {
		http.Error(rw, "invalid category id", http.StatusBadRequest)
		return
	}

	err = c.service.Delete(categoryIntID)
	if err != nil {
		c.logger.Error("error delete cateogry", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (c *CategoryHandler) GetAll(rw http.ResponseWriter, r *http.Request) {
	categories, err := c.service.GetAll()
	if err != nil {
		c.logger.Error("error get all categories", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}

	err = helpers.SendSuccess(rw, categories, http.StatusOK)
	if err != nil {
		c.logger.Error("error encode response", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}
}
