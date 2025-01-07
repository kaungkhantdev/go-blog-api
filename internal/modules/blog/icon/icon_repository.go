package icon

import (
	"gorm.io/gorm"
)

type IconRepository struct {
	db *gorm.DB
}

func NewIconRepository(db *gorm.DB) IconRepositoryInterface {
	return &IconRepository{db: db}
}

func (repo *IconRepository) CreateIcon(data *IconEntity) (IconEntity, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return IconEntity{}, err
	}

	return *data, nil
}

func (repo *IconRepository) UpdateIcon(id int, data *IconEntity) (IconEntity, error) {
	var icon IconEntity
	if err := repo.db.First(&icon, id).Error; err != nil {
		return IconEntity{}, err
	}

	if err := repo.db.Model(&icon).Updates(data).Error; err != nil {
		return IconEntity{}, err
	}
	return icon, nil
}

func (repo *IconRepository) FindByName(name string) (IconEntity, error) {
	var icon IconEntity
	if err := repo.db.Where("name = ?", name).First(&icon).Error; err != nil {
		return IconEntity{}, err
	}

	return icon, nil
}

func (repo *IconRepository) FindByIdIcon(id int) (IconEntity, error) {
	var icon IconEntity
	if err := repo.db.First(&icon, id).Error; err != nil {
		return IconEntity{}, err
	}
	return icon, nil
}
