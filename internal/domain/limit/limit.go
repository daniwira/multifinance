package domainlimit

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Limit struct {
	ID        uint64         `gorm:"primary_key" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	Name      string         `gorm:"column:name;index" json:"name"`
	Tenor1    float64        `gorm:"column:tenor1;index" json:"tenor1"`
	Tenor2    float64        `gorm:"column:tenor2;index" json:"tenor2"`
	Tenor3    float64        `gorm:"column:tenor3;index" json:"tenor3"`
	Tenor4    float64        `gorm:"column:tenor4;index" json:"tenor4"`
}

func (l *Limit) GetTenorValue(month int) (float64, error) {
	switch month {
	case 1:
		return l.Tenor1, nil
	case 2:
		return l.Tenor2, nil
	case 3:
		return l.Tenor3, nil
	case 4:
		return l.Tenor4, nil
	default:
		return 0.0, fmt.Errorf("invalid month for retrieving tenor value")
	}
}
