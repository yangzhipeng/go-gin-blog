package models

type Tag struct {
	BaseModel

	Name string `form:"name" gorm:"column:name" json:"name"` // 名称
}

func (Tag) TableName() string {
	return "blog_tags"
}
