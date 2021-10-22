package models

type Article struct {
	BaseModel

	UserID     int      `form:"user_id" gorm:"index;type:int" json:"user_id"`                              // 用户id
	TagID      int      `form:"tag_id" gorm:"index;column:tag_id;type:int" json:"tag_id"`                  // 标签id
	CategoryID int      `form:"category_id" gorm:"index;column:category_id;type:int" json:"category_id"`   // 分类id
	Title      string   `form:"title" gorm:"column:title;type:varchar(200)" json:"title"`                  // 标题
	Desc       string   `form:"desc" gorm:"column:desc;type:varchar(512)" json:"desc"`                     // 描述
	Content    string   `form:"content" gorm:"column:content;type:longtext"  json:"content"`               // 内容
	User       User     `gorm:"association_autoupdate:false;association_autocreate:false" json:"user"`     // 用户关联
	Tag        Tag      `gorm:"association_autoupdate:false;association_autocreate:false" json:"tag"`      // 标签关联
	Category   Category `gorm:"association_autoupdate:false;association_autocreate:false" json:"category"` // 分类关联
}

func (Article) TableName() string {
	return "blog_articles"
}
