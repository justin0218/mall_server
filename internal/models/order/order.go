package order

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	OrderCode      string    `json:"order_code"`
	ThirdOrderCode string    `json:"third_order_code"`
	GoodsId        int64     `json:"goods_id"`
	GoodsName      string    `json:"goods_name"`
	SkuId          int64     `json:"sku_id"`
	SkuName        string    `json:"sku_name"`
	Count          int64     `json:"count"`
	Price          int64     `json:"price"`
	DealPrice      int       `json:"deal_price"`
	Uid            int64     `json:"uid"`
	Status         int       `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Model struct {
	Db   *gorm.DB
	Name string
}

func NewModel(db *gorm.DB) *Model {
	return &Model{
		Db:   db,
		Name: "orders",
	}
}

func (s *Model) Create(in Order) (err error) {
	db := s.Db.Table(s.Name)
	err = db.Create(&in).Error
	if err != nil {
		return
	}
	return
}

func (s *Model) GetByStatus(status int) (ret []Order, err error) {
	db := s.Db.Table(s.Name)
	err = db.Where("status = ?", status).Find(&ret).Error
	if err != nil {
		return
	}
	return
}

func (s *Model) UpdateStatusByOrderCode(orderCode string, status int) (err error) {
	db := s.Db.Table(s.Name)
	err = db.Where("order_code = ?", orderCode).UpdateColumn("status", status).Error
	if err != nil {
		return
	}
	return
}
