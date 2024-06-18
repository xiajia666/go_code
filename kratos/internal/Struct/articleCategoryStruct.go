package Struct

type ArticleCategory struct {
	Id        int       `json:"id"` //主键
	Title     string    `json:"category"      required:"true"`
	States    int       `json:"state"         required:"true"`
	Articles  []Article `gorm:"foreignKey:CateId"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`
}

func (ArticleCategory) TableName() string {
	return "lt_article_category_info"
}
