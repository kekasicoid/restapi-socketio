package table

//IdOrang     Orang  `json:"id_orang" gorm:"foreignKey:id" swaggerignore:"true"`

type Asset struct {
	Id          int `json:"id" gorm:"primarykey"`
	OrangID     int
	Orang       Orang  `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	IdProduct   int    `json:"id_product" gorm:"column:id_product"`
	Tittle      string `json:"tittle" gorm:"column:tittle;size:255"`
	Description string `json:"description" gorm:"column:description;size:255"`
	Price       int    `json:"price" gorm:"column:price"`
	Brand       string `json:"brand" gorm:"column:Brand;size:255"`
	Category    string `json:"category" gorm:"column:Category;size:255"`
	Thumbnail   string `json:"thumbnail" gorm:"column:thumbnail;type:text"`
	ModelDate   `swaggerignore:"true"`
}

func (t Asset) TableName() string {
	return "assets"
}
