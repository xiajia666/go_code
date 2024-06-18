package Struct

type Article struct {
	Id              int             `json:"id"`
	ArticleName     string          `json:"articleName"    required:"true"`
	CateId          int             `json:"article_category_id"` //外键
	ArticleCategory ArticleCategory `gorm:"foreignKey:CateId"`
	CreatedAt       string          `json:"createdAt"`
	UpdatedAt       string          `json:"updatedAt"`
	DeletedAt       string          `json:"deletedAt"`
}

func (Article) TableName() string {
	return "lt_article_info"
}
