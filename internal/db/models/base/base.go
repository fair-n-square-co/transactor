package base

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PrimaryKey struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
}

type DateTime struct {
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type SoftDeleteModel struct {
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
