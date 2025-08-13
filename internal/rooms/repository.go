package rooms

import "gorm.io/gorm"

type RoomRepository struct {
	db *gorm.DB
}

func (r *RoomRepository) Create(dto RoomModel) error {
	return r.db.Create(&dto).Error
}

func (r *RoomRepository) Update(dto RoomModel) error {
	return r.db.Save(&dto).Error
}

func (r *RoomRepository) Get(dto RoomModel) ([]RoomModel, error) {
	var results []RoomModel
	err := r.db.Where(&dto).Find(&results).Error
	return results, err
}

func (r *RoomRepository) Delete(id uint) error {
	return r.db.Delete(&gorm.Model{ID: id}).Error
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{db: db}
}
