package service

import (
	"fmt"
	"os"
	"path"

	"github.com/Zrossiz/Home_menu/backend/internal/dto"
	"github.com/google/uuid"
)

type AttachmentService struct {
	attachmentPostgresStorage AttachmentPostgresStorage
}

type AttachmentPostgresStorage interface {
	Create(payload dto.CreateAttachmentDTO) error
	GetAllByDish(dishID int) ([]dto.AttachmentDTO, error)
}

func NewAttachmentService(
	attachmentPostgresStorage AttachmentPostgresStorage,
) *AttachmentService {
	return &AttachmentService{
		attachmentPostgresStorage: attachmentPostgresStorage,
	}
}

func (a *AttachmentService) Create(payload dto.CreateAttachmentDTO) (string, error) {
	key := a.generateKey()
	payload.Key = key + payload.Ext

	err := a.attachmentPostgresStorage.Create(payload)
	if err != nil {
		return "", fmt.Errorf("error create attachment: %v", err)
	}

	return key, nil
}

func (a *AttachmentService) GetAllByDish(dishID int) ([]dto.AttachmentDTO, error) {
	return a.attachmentPostgresStorage.GetAllByDish(dishID)
}

func (a *AttachmentService) DeleteAllByDish(dishID int) error {
	allAttachments, err := a.GetAllByDish(dishID)
	if err != nil {
		return fmt.Errorf("error get attachments for delete: %v", err)
	}
	applicationPath := path.Join(os.Getenv("APPLICATION_PATH"), "uploads")

	for _, attachment := range allAttachments {
		attachmentPath := path.Join(applicationPath, attachment.Key)
		err = os.Remove(attachmentPath)
		if err != nil {
			err = fmt.Errorf("error remove file: %v", err)
		}
	}

	return err
}

func (a *AttachmentService) generateKey() string {
	id := uuid.New()
	return id.String()
}
