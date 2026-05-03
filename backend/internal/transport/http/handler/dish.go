package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Zrossiz/Home_menu/backend/internal/apperrors"
	"github.com/Zrossiz/Home_menu/backend/internal/dto"
	"github.com/Zrossiz/Home_menu/backend/internal/helpers"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type DishHandler struct {
	dishService       DishService
	attachmentService AttachmentService
	logger            *zap.Logger
}

type DishService interface {
	Create(payload dto.CreateDishDTO) error
	Delete(dishID int) error
	GetAllByCategory(categoryID int) ([]dto.DishDTO, error)
	GetOne(dishID int) (*dto.DishDTO, error)
	Update(dishID int, payload dto.UpdateDishDTO) error
	GetPublicPathsForImages(keys []dto.AttachmentDTO) []string
	Find(search string) ([]dto.DishDTO, error)
}

func NewDishHandler(dishService DishService, attachmentService AttachmentService, logger *zap.Logger) *DishHandler {
	return &DishHandler{
		dishService:       dishService,
		attachmentService: attachmentService,
		logger:            logger,
	}
}

func (d *DishHandler) Find(rw http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")

	dishes, err := d.dishService.Find(search)
	if err != nil {
		d.logger.Error("find dishes error", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}

	err = helpers.SendSuccess(rw, dishes, http.StatusOK)
	if err != nil {
		d.logger.Error("error encode response", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (d *DishHandler) Create(rw http.ResponseWriter, r *http.Request) {
	var body dto.CreateDishDTO

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(rw, "invalid request body", http.StatusBadRequest)
		return
	}

	err = d.dishService.Create(body)
	if err != nil {
		d.logger.Error("error create dish", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

func (d *DishHandler) GetOne(rw http.ResponseWriter, r *http.Request) {
	dishStringID := chi.URLParam(r, "dishID")
	dishIntID, err := strconv.Atoi(dishStringID)
	if err != nil {
		http.Error(rw, "invalid dish id", http.StatusBadRequest)
		return
	}

	dish, err := d.dishService.GetOne(dishIntID)
	if err != nil {
		if err == apperrors.ErrNotFound {
			http.Error(rw, "not found", http.StatusNotFound)
			return
		}
		d.logger.Error("error get dish by id", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}

	attachments, err := d.attachmentService.GetAllByDish(dish.ID)
	if err != nil {
		d.logger.Error("error get attachments", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}

	var images []string
	for _, v := range attachments {
		images = append(images, v.Key)
	}

	result := dto.GetDishDTO{
		DishDTO: *dish,
		Images:  images,
	}

	err = helpers.SendSuccess(rw, result, http.StatusOK)
	if err != nil {
		d.logger.Error("error encode response", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}
}

func (d *DishHandler) Update(rw http.ResponseWriter, r *http.Request) {
	dishStringID := chi.URLParam(r, "dishID")
	dishIntID, err := strconv.Atoi(dishStringID)
	if err != nil {
		http.Error(rw, "invalid dish id", http.StatusBadRequest)
		return
	}

	var body dto.UpdateDishDTO
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(rw, "invalid dish id", http.StatusBadRequest)
		return
	}

	err = d.dishService.Update(dishIntID, body)
	if err != nil {
		d.logger.Error("error update dish", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (d *DishHandler) Delete(rw http.ResponseWriter, r *http.Request) {
	dishStringID := chi.URLParam(r, "dishID")
	dishIntID, err := strconv.Atoi(dishStringID)
	if err != nil {
		http.Error(rw, "invalid dish id", http.StatusBadRequest)
		return
	}

	err = d.dishService.Delete(dishIntID)
	if err != nil {
		d.logger.Error("error delete dish", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (d *DishHandler) GetAllByCategory(rw http.ResponseWriter, r *http.Request) {
	categoryStringID := chi.URLParam(r, "categoryID")
	categoryIntID, err := strconv.Atoi(categoryStringID)
	if err != nil {
		http.Error(rw, "invalid category id", http.StatusBadRequest)
		return
	}

	data, err := d.dishService.GetAllByCategory(categoryIntID)
	if err != nil {
		d.logger.Error("error get all dishes by category", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}

	err = helpers.SendSuccess(rw, data, http.StatusOK)
	if err != nil {
		d.logger.Error("error send all dishes by category", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}
}
