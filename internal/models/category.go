package models

type Category struct {
	BaseModel

	Name string `form:"title" gorm:"column:name;type:varchar(32)" json:"name"` // 名称
	Desc string `form:"desc" gorm:"column:desc;type:varchar(256)" json:"desc"` // 描述
}

func (Category) TableName() string {
	return "blog_categories"
}
