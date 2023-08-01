package domain

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          string         `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	Name        string         `json:"name" gorm:"type:char(50);not null"`
	Sku         string         `json:"sku" gorm:"type:char(30);not null"`
	Quantity    int32          `json:"quantity" gorm:"type:int;not null"`
	Price       float32        `json:"price" gorm:"type:decimal;not null"`
	CostPrice   float32        `json:"costprice" gorm:"type:decimal;not null"`
	Weight      int32          `json:"wight" gorm:"type:int;not null"`
	Enabled     bool           `json:"enabled" gorm:"type:boolean;not null"`
	Descripcion string         `json:"descripcion" gorm:"type:char(50);not null"`
	CategoryID  string         `json:"category_id,omitempty" gorm:"-"`
	Category    *Category      `json:"category" gorm:"type:char(36);not null"`
	CreatedAt   *time.Time     `json:"-"`
	UpdateAt    *time.Time     `json:"-"`
	Deleted     gorm.DeletedAt `json:"-"`
}

type Category struct {
	ID          string         `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	Descripción string         `json:"descripcion" gorm:"type:char(50);not null"`
	CreatedAt   *time.Time     `json:"-"`
	UpdateAt    *time.Time     `json:"-"`
	Deleted     gorm.DeletedAt `json:"-"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return
}
