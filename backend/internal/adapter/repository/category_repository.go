package repository

import (
	"context"
	"errors"
	"fmt"
	"newsportal-backend/internal/core/domain/entity"
	"newsportal-backend/internal/core/domain/model"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategories(ctx context.Context) ([]entity.CategoryEntity, error)
	GetCategoryByID(ctx context.Context, id int64) (*entity.CategoryEntity, error)
	CreateCategory(ctx context.Context, req entity.CategoryEntity) error
	EditCategory(ctx context.Context, req entity.CategoryEntity) error
	DeleteCategory(ctx context.Context, id int64) error
}

type categoryRepository struct {
	db *gorm.DB
}

// CreateCategory implements CategoryRepository.
func (ca *categoryRepository) CreateCategory(ctx context.Context, req entity.CategoryEntity) error {
	var countSlug int64
	err = ca.db.Table("categories").Where("slug = ?", req.Slug).Count(&countSlug).Error
	if err != nil {
		code = "[REPOSITORY] CreateCategory - 1"
		log.Errorw(code, err)
		return err
	}

	countSlug = countSlug + 1
	slug := fmt.Sprintf("%s-%d", req.Slug, countSlug)
	modelCategory := model.Category{
		Title:       req.Title,
		Slug:        slug,
		CreatedByID: req.User.ID,
	}

	err = ca.db.Create(&modelCategory).Error
	if err != nil {
		code = "[REPOSITORY] CreateCategories - 2"
		log.Errorw(code, err)
		return err
	}

	return nil
}

// DeleteCategory implements CategoryRepository.
func (ca *categoryRepository) DeleteCategory(ctx context.Context, id int64) error {
	var count int64
	err = ca.db.Table("contents").Where("category_id = ?", id).Count(&count).Error
	if err != nil {
		code = "[REPOSITORY] DeleteCategory - 1"
		log.Errorw(code, err)
		return err
	}

	if count > 0 {
		return errors.New("Cannot delete category with associated contents")
	}

	err = ca.db.Where("id = ?", id).Delete(&model.Category{}).Error
	if err != nil {
		code = "[REPOSITORY] DeleteCategory - 2"
		log.Errorw(code, err)
		return err
	}

	return nil
}

// EditCategory implements CategoryRepository.
func (ca *categoryRepository) EditCategory(ctx context.Context, req entity.CategoryEntity) error {
	var countSlug int64
	err := ca.db.Table("categories").Where("slug = ?", req.Slug).Count(&countSlug).Error
	if err != nil {
		code = "[REPOSITORY] EditCategory - 1"
		log.Errorw(code, err)
		return err
	}

	countSlug = countSlug + 1
	slug := req.Slug
	if countSlug == 0 {
		slug = fmt.Sprintf("%s-%d", req.Slug, countSlug)
	}
	modelCategory := model.Category{
		Title:       req.Title,
		Slug:        slug,
		CreatedByID: req.User.ID,
	}

	err = ca.db.Where("id = ?", req.ID).Updates(&modelCategory).Error
	if err != nil {
		code = "[REPOSITORY] EditCategory - 2"
		log.Errorw(code, err)
		return err
	}

	return nil
}

// GetCategories implements CategoryRepository.
func (ca *categoryRepository) GetCategories(ctx context.Context) ([]entity.CategoryEntity, error) {
	var modelCategories []model.Category

	err := ca.db.Order("created_at DESC").Preload("User").Find(&modelCategories).Error
	if err != nil {
		code = "[REPOSITORY] GetCategories - 1"
		log.Errorw(code, err)
		return nil, err
	}

	if len(modelCategories) == 0 {
		code = "[REPOSITORY] GetCategories - 2"
		err = errors.New("data not found")
		log.Errorw(code, err)
		return nil, err
	}

	var resps []entity.CategoryEntity
	for _, val := range modelCategories {
		resps = append(resps, entity.CategoryEntity{
			ID:    val.ID,
			Title: val.Title,
			Slug:  val.Slug,
			User: entity.UserEntity{
				ID:       val.User.ID,
				Name:     val.User.Name,
				Email:    val.User.Email,
				Password: val.User.Password,
			},
		})
	}

	return resps, nil
}

// GetCategoryByID implements CategoryRepository.
func (ca *categoryRepository) GetCategoryByID(ctx context.Context, id int64) (*entity.CategoryEntity, error) {
	var modelCategory model.Category
	err = ca.db.Where("id = ?", id).Preload("User").First(&modelCategory).Error
	if err != nil {
		code = "[REPOSITORY] GetCategoryByID - 1"
		log.Errorw(code, err)
		return nil, err
	}

	return &entity.CategoryEntity{
		ID:    modelCategory.ID,
		Title: modelCategory.Title,
		Slug:  modelCategory.Slug,
		User: entity.UserEntity{
			ID:    modelCategory.User.ID,
			Name:  modelCategory.User.Name,
			Email: modelCategory.User.Email,
		},
	}, nil
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}
