package table

import (
	"time"

	"github.com/kekasicoid/restapi-socketio/enum"
	"gorm.io/gorm"
)

type ModelDate struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Orang struct {
	Id           int               `json:"id" gorm:"primarykey"`
	Nama         string            `json:"nama" gorm:"column:nama;size:255"`
	JenisKelamin enum.JenisKelamin `json:"jenis_kelamin" gorm:"column:jenis_kelamin;type:int;size:1;default:0;comment:0 Laki-laki, 1 wanita"`
	OrangTua     int               `json:"orang_tua" gorm:"column:orang_tua;size:255;default:null;"`
	Anak         []Orang           `gorm:"foreignkey:OrangTua"`
	ModelDate    `swaggerignore:"true"`
}

func (t Orang) TableName() string {
	return "orang"
}
