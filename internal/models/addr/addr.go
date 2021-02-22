package addr

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Addr struct {
	Uid       int64     `json:"uid"`
	OrderCode string    `json:"order_code"`
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Province  string    `json:"province"`
	City      string    `json:"city"`
	Region    string    `json:"region"`
	Addr      string    `json:"addr"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Model struct {
	Db   *gorm.DB
	Name string
}

func NewModel(db *gorm.DB) *Model {
	return &Model{
		Db:   db,
		Name: "addrs",
	}
}

func (s *Model) Create(in Addr) (err error) {
	db := s.Db.Table(s.Name)
	err = db.Create(&in).Error
	if err != nil {
		return
	}
	return
}
