package services

import (
	"go-blog-api/internal/icon/interfaces"
	"go-blog-api/internal/icon/models"
)

type IconService struct {
	repo interfaces.IconRepositoryInterface
}

func NewIconService(repo *interfaces.IconRepositoryInterface) *IconService {
	return &IconService{repo: *repo}
}

func (service *IconService) CreateIcon(data *models.Icon) (models.Icon, error) {
	newIcon, err := service.repo.CreateIcon(data)
	if err != nil {
		return models.Icon{}, err
	}

	return newIcon, nil
}

func (service *IconService) UpdateIcon(id int, data *models.Icon) (models.Icon, error) {
	updateIcon, err := service.repo.UpdateIcon(id, data)
	if err != nil {
		return models.Icon{}, err
	}

	return updateIcon, nil
}

func (service *IconService) FindByName(name string) (models.Icon, error) {
	icon, err := service.repo.FindByName(name)
	if err != nil {
		return models.Icon{}, err
	}

	return icon, nil
}

func (service *IconService) FindbyIdIcon(id int) (models.Icon, error) {
	icon, err := service.repo.FindByIdIcon(id)
	if err != nil {
		return models.Icon{}, err
	}

	return icon, nil
}
