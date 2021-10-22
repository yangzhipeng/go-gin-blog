package models

type User struct {
	BaseModel

	Username string `form:"username" gorm:"column:username;type:varchar(20)" json:"username"` // 用户名
	Password string `gorm:"column:password;type:varchar(20)" json:"password"`                 // 密码
	Desc     string `gorm:"column:desc;type:varchar(512)" json:"desc"`                        // 描述
	Role     string `gorm:"column:role;type:varchar(20)" json:"role"`                         // 角色
}

func (User) TableName() string {
	return "blog_users"
}
