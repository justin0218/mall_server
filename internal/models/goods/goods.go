package goods

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Goods struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Images    string    `json:"images"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Image struct {
	Url string `json:"url"`
}

type Model struct {
	Db   *gorm.DB
	Name string
}

func NewModel(db *gorm.DB) *Model {
	return &Model{
		Db:   db,
		Name: "goods",
	}
}

func (s *Model) GetOneById(goodsId int) (ret Goods, err error) {
	db := s.Db.Table(s.Name)
	err = db.Where("id = ?", goodsId).First(&ret).Error
	if err != nil {
		return
	}
	return
}
