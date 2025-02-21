package service

import (
	"context"
	"newsportal-backend/config"
	"newsportal-backend/internal/adapter/cloudflare"
	"newsportal-backend/internal/adapter/repository"
	"newsportal-backend/internal/core/domain/entity"

	"github.com/gofiber/fiber/v2/log"
)

type ContentService interface {
	GetContents(ctx context.Context) ([]entity.ContentEntity, error)
	GetContentByID(ctx context.Context, id int64) (*entity.ContentEntity, error)
	CreateContent(ctx context.Context, req entity.ContentEntity) error
	UpdateContent(ctx context.Context, req entity.ContentEntity) error
	DeleteContent(ctx context.Context, id int64) error
	UploadImageR2(ctx context.Context, req entity.FileUploadEntity) (string, error)
}

type contentService struct {
	contentRepository repository.ContentRepository
	cfg               *config.Config
	r2                cloudflare.CloudflareR2Adapter
}

// CreateContent implements ContentService.
func (co *contentService) CreateContent(ctx context.Context, req entity.ContentEntity) error {
	err = co.contentRepository.CreateContent(ctx, req)
	if err != nil {
		code := "[SERVICE] CreateContent - 1"
		log.Errorw(code, err)
		return err
	}

	return nil
}

// DeleteContent implements ContentService.
func (co *contentService) DeleteContent(ctx context.Context, id int64) error {
	err := co.contentRepository.DeleteContent(ctx, id)
	if err != nil {
		code := "[SERVICE] DeleteContent - 1"
		log.Errorw(code, err)
		return err
	}

	return nil
}

// GetContentByID implements ContentService.
func (co *contentService) GetContentByID(ctx context.Context, id int64) (*entity.ContentEntity, error) {
	result, err := co.contentRepository.GetContentByID(ctx, id)
	if err != nil {
		code := "[SERVICE] GetContentByID - 1"
		log.Errorw(code, err)
		return nil, err
	}

	return result, nil
}

// GetContents implements ContentService.
func (co *contentService) GetContents(ctx context.Context) ([]entity.ContentEntity, error) {
	results, err := co.contentRepository.GetContents(ctx)
	if err != nil {
		code := "[SERVICE] GetContents - 1"
		log.Errorw(code, err)
		return nil, err
	}

	return results, nil
}

// UpdateContent implements ContentService.
func (co *contentService) UpdateContent(ctx context.Context, req entity.ContentEntity) error {
	err = co.contentRepository.UpdateContent(ctx, req)
	if err != nil {
		code := "[SERVICE] UpdateContent - 1"
		log.Errorw(code, err)
		return err
	}

	return nil
}

// UploadImageR2 implements ContentService.
func (co *contentService) UploadImageR2(ctx context.Context, req entity.FileUploadEntity) (string, error) {
	urlImage, err := co.r2.UploadImage(&req)
	if err != nil {
		code := "[SERVICE] UploadImageR2 - 1"
		log.Errorw(code, err)
		return "", err
	}

	return urlImage, nil
}

func NewContentService(contentRepository repository.ContentRepository, cfg *config.Config, r2 cloudflare.CloudflareR2Adapter) ContentService {
	return &contentService{
		contentRepository: contentRepository,
		cfg:               cfg,
		r2:                r2,
	}
}
