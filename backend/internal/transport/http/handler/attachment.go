package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Zrossiz/Home_menu/backend/internal/dto"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type AttachmentHandler struct {
	logger  *zap.Logger
	service AttachmentService
}

type AttachmentService interface {
	Create(payload dto.CreateAttachmentDTO) (string, error)
	GetAllByDish(dishID int) ([]dto.AttachmentDTO, error)
}

func NewAttachmentHandler(service AttachmentService, logger *zap.Logger) *AttachmentHandler {
	return &AttachmentHandler{service: service, logger: logger}
}

func (a *AttachmentHandler) Create(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(20 << 20)
	if err != nil {
		http.Error(rw, "error parsing form", http.StatusBadRequest)
		return
	}

	uploadDir := filepath.Join(os.Getenv("APPLICATION_PATH"), "uploads")
	os.MkdirAll(uploadDir, os.ModePerm)

	files := r.MultipartForm.File["files"]

	dishStringID := chi.URLParam(r, "dishID")
	dishIntID, err := strconv.Atoi(dishStringID)
	if err != nil {
		http.Error(rw, "invalid dish id", http.StatusBadRequest)
		return
	}

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(rw, "open file error", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		ext := filepath.Ext(fileHeader.Filename)

		key, err := a.service.Create(dto.CreateAttachmentDTO{
			DishID: dishIntID,
			Ext:    ext,
		})
		if err != nil {
			a.logger.Error("error create attachment", zap.Error(err))
			http.Error(rw, "internal server error", http.StatusInternalServerError)
			return
		}

		dstPath := filepath.Join(uploadDir, key+ext)
		dst, err := os.Create(dstPath)
		if err != nil {
			http.Error(rw, "create file error", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(rw, "write file errro", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(rw, "Загружен файл: %s\n", fileHeader.Filename)
	}
}

func (a *AttachmentHandler) GetOne(rw http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	if key == "" {
		http.Error(rw, "missing key param", http.StatusBadRequest)
		return
	}

	uploadDir := filepath.Join(os.Getenv("APPLICATION_PATH"), "uploads")
	filePath := filepath.Join(uploadDir, key)

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(rw, "file not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	contentType := "application/octet-stream"
	ext := filepath.Ext(filePath)
	switch ext {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	case ".gif":
		contentType = "image/gif"
	}

	stat, err := file.Stat()
	if err != nil {
		http.Error(rw, "internal error", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", contentType)
	rw.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))
	rw.Header().Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, key))

	_, err = io.Copy(rw, file)
	if err != nil {
		a.logger.Error("error streaming file", zap.Error(err))
	}
}
