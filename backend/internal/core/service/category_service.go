package service

import (
	"context"
	"newsportal-backend/internal/adapter/repository"
	"newsportal-backend/internal/core/domain/entity"
	"newsportal-backend/lib/conv"

	"github.com/gofiber/fiber/v2/log"
)

type CategoryService interface {
	GetCategories(ctx context.Context) ([]entity.CategoryEntity, error)
	GetCategoryByID(ctx context.Context, id int64) (*entity.CategoryEntity, error)
	CreateCategory(ctx context.Context, req entity.CategoryEntity) error
	EditCategory(ctx context.Context, req entity.CategoryEntity) error
	DeleteCategory(ctx context.Context, id int64) error
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

// CreateCategory implements CategoryService.
func (ca *categoryService) CreateCategory(ctx context.Context, req entity.CategoryEntity) error {
	slug := conv.GenerateSlug(req.Title)
	req.Slug = slug

	err := ca.categoryRepository.CreateCategory(ctx, req)
	if err != nil {
		code = "[SERVICE] CreateCategories - 1"
		log.Errorw(code, err)
		return err
	}

	return nil
}

// DeleteCategory implements CategoryService.
func (ca *categoryService) DeleteCategory(ctx context.Context, id int64) error {
	err := ca.categoryRepository.DeleteCategory(ctx, id)
	if err != nil {
		code = "[SERVICE] DeleteCategory - 1"
		log.Errorw(code, err)
		return err
	}

	return nil
}

// EditCategory implements CategoryService.
func (ca *categoryService) EditCategory(ctx context.Context, req entity.CategoryEntity) error {
	categoryData, err := ca.categoryRepository.GetCategoryByID(ctx, req.ID)
	if err != nil {
		code = "[SERVICE] EditCategory - 1"
		log.Errorw(code, err)
		return err
	}

	slug := conv.GenerateSlug(req.Title)
	if categoryData.Title == req.Title {
		slug = categoryData.Slug
	}

	req.Slug = slug

	err = ca.categoryRepository.EditCategory(ctx, req)
	if err != nil {
		code = "[SERVICE] EditCategory - 2"
		log.Errorw(code, err)
		return err
	}

	return nil
}

// GetCategories implements CategoryService.
func (ca *categoryService) GetCategories(ctx context.Context) ([]entity.CategoryEntity, error) {
	results, err := ca.categoryRepository.GetCategories(ctx)
	if err != nil {
		code = "[SERVICE] GetCategories - 1"
		log.Errorw(code, err)
		return nil, err
	}

	return results, nil
}

// GetCategoryByID implements CategoryService.
func (ca *categoryService) GetCategoryByID(ctx context.Context, id int64) (*entity.CategoryEntity, error) {
	result, err := ca.categoryRepository.GetCategoryByID(ctx, id)
	if err != nil {
		code = "[SERVICE] GetCategoryByID - 1"
		log.Errorw(code, err)
		return nil, err
	}

	return result, nil
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository: categoryRepository}
}
