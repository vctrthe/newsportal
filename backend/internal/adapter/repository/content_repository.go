package repository

import (
	"context"
	"fmt"
	"math"
	"newsportal-backend/internal/core/domain/entity"
	"newsportal-backend/internal/core/domain/model"
	"strings"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ContentRepository interface {
	GetContents(ctx context.Context, query entity.QueryString) ([]entity.ContentEntity, int64, int64, error)
	GetContentByID(ctx context.Context, id int64) (*entity.ContentEntity, error)
	CreateContent(ctx context.Context, req entity.ContentEntity) error
	UpdateContent(ctx context.Context, req entity.ContentEntity) error
	DeleteContent(ctx context.Context, id int64) error
}

type contentRepository struct {
	db *gorm.DB
}

// CreateContent implements ContentRepository.
func (co *contentRepository) CreateContent(ctx context.Context, req entity.ContentEntity) error {
	tags := strings.Join(req.Tags, ",")
	modelContent := model.Content{
		Title:       req.Title,
		Excerpt:     req.Excerpt,
		Description: req.Description,
		Image:       req.Image,
		Tags:        tags,
		Status:      req.Status,
		CategoryID:  req.CategoryID,
		CreatedByID: req.CreatedByID,
	}

	err = co.db.Create(&modelContent).Error
	if err != nil {
		code = "[REPOSITORY] CreateContent - 1"
		log.Errorw(code, err)
		return err
	}

	return nil
}

// DeleteContent implements ContentRepository.
func (co *contentRepository) DeleteContent(ctx context.Context, id int64) error {
	err := co.db.Where("id = ?", id).Delete(&model.Content{}).Error
	if err != nil {
		code = "[REPOSITORY] DeleteContent - 1"
		log.Errorw(code, err)
		return err
	}

	return nil
}

// GetContentByID implements ContentRepository.
func (co *contentRepository) GetContentByID(ctx context.Context, id int64) (*entity.ContentEntity, error) {
	var modelContent model.Content
	err := co.db.Where("id = ?", id).Preload(clause.Associations).First(&modelContent).Error
	if err != nil {
		code = "[REPOSITORY] GetContentByID - 1"
		log.Errorw(code, err)
		return nil, err
	}

	tags := strings.Split(modelContent.Tags, ",")
	resp := entity.ContentEntity{
		ID:          modelContent.ID,
		Title:       modelContent.Title,
		Excerpt:     modelContent.Excerpt,
		Description: modelContent.Description,
		Image:       modelContent.Image,
		Tags:        tags,
		Status:      modelContent.Status,
		CategoryID:  modelContent.CategoryID,
		CreatedByID: modelContent.CreatedByID,
		CreatedAt:   modelContent.CreatedAt,
		Category: entity.CategoryEntity{
			ID:    modelContent.CategoryID,
			Title: modelContent.Category.Title,
			Slug:  modelContent.Category.Slug,
		},
		User: entity.UserEntity{
			ID:       modelContent.User.ID,
			Name:     modelContent.User.Name,
			Email:    modelContent.User.Email,
			Password: modelContent.User.Password,
		},
	}

	return &resp, nil
}

// GetContents implements ContentRepository.
func (co *contentRepository) GetContents(ctx context.Context, query entity.QueryString) ([]entity.ContentEntity, int64, int64, error) {
	var modelContents []model.Content
	var countData int64

	order := fmt.Sprintf("%s %s", query.OrderBy, query.OrderType)
	offset := (query.Page - 1) * query.Limit
	status := ""
	if query.Status != "" {
		status = query.Status
	}

	sqlMain := co.db.Preload(clause.Associations).
		Where("title ilike ? OR excerpt ilike ? OR description ilike ?", "%"+query.Search+"%", "%"+query.Search+"%", "%"+query.Search+"%").
		Where("status LIKE ?", "%"+status+"%")

	if query.CategoryID > 0 {
		sqlMain = sqlMain.Where("category_id = ?", query.CategoryID)
	}

	err = sqlMain.Model(&modelContents).Count(&countData).Error
	if err != nil {
		code = "[REPOSITORY] GetContents - 1"
		log.Errorw(code, err)
		return nil, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(countData) / float64(query.Limit)))

	err = sqlMain.Order(order).Limit(query.Limit).Offset(offset).Find(&modelContents).Error
	if err != nil {
		code = "[REPOSITORY] GetContents - 2"
		log.Errorw(code, err)
		return nil, 0, 0, err
	}

	resps := []entity.ContentEntity{}
	for _, val := range modelContents {
		tags := strings.Split(val.Tags, ",")
		resp := entity.ContentEntity{
			ID:          val.ID,
			Title:       val.Title,
			Excerpt:     val.Excerpt,
			Description: val.Description,
			Image:       val.Image,
			Tags:        tags,
			Status:      val.Status,
			CategoryID:  val.CategoryID,
			CreatedByID: val.CreatedByID,
			CreatedAt:   val.CreatedAt,
			Category: entity.CategoryEntity{
				ID:    val.CategoryID,
				Title: val.Category.Title,
				Slug:  val.Category.Slug,
			},
			User: entity.UserEntity{
				ID:       val.User.ID,
				Name:     val.User.Name,
				Email:    val.User.Email,
				Password: val.User.Password,
			},
		}

		resps = append(resps, resp)
	}

	return resps, countData, int64(totalPages), nil
}

// UpdateContent implements ContentRepository.
func (co *contentRepository) UpdateContent(ctx context.Context, req entity.ContentEntity) error {
	tags := strings.Join(req.Tags, ",")
	modelContent := model.Content{
		Title:       req.Title,
		Excerpt:     req.Excerpt,
		Description: req.Description,
		Image:       req.Image,
		Tags:        tags,
		Status:      req.Status,
		CategoryID:  req.CategoryID,
		CreatedByID: req.CreatedByID,
	}

	err = co.db.Where("id = ?", req.ID).Updates(&modelContent).Error
	if err != nil {
		code = "[REPOSITORY] UpdateContent - 1"
		log.Errorw(code, err)
		return err
	}

	return nil
}

func NewContentRepository(db *gorm.DB) ContentRepository {
	return &contentRepository{db: db}
}
