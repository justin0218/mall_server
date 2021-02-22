package sku

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Sku struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Cover     string    `json:"cover"`
	GoodsId   int       `json:"goods_id"`
	Price     int       `json:"price"`
	Inventory int       `json:"inventory"`
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
		Name: "skus",
	}
}

func (s *Model) GetMultiByGoodsId(goodsId int) (ret []Sku, err error) {
	db := s.Db.Table(s.Name)
	err = db.Where("goods_id = ?", goodsId).Find(&ret).Error
	if err != nil {
		return
	}
	return
}

func (s *Model) DecrInventory(id, buyNum int) (err error) {
	db := s.Db.Table(s.Name)
	err = db.Where("id = ?", id).UpdateColumn("inventory", gorm.Expr("inventory - ?", buyNum)).Error
	if err != nil {
		return
	}
	return
}

func (s *Model) GetById(id int) (ret Sku, err error) {
	db := s.Db.Table(s.Name)
	err = db.Where("id = ?", id).First(&ret).Error
	if err != nil {
		return
	}
	return
}
