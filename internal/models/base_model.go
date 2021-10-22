package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	StatusOff int = iota
	StatusOn
)

type BaseModel struct {
	ID        int        `gorm:"primary_key;auto_increment" json:"id"`                         // id
	Status    int        `form:"status,default=1" gorm:"column:status;type:int" json:"status"` // 状态
	CreatedAt CustomTime `gorm:"type:datetime;" json:"created_at"`                             // 创建时间
	UpdatedAt CustomTime `gorm:"type:datetime;" json:"updated_at"`                             // 更新时间
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`                                       // 删除时间
}

// 自定义时间
type CustomTime time.Time

func (t *CustomTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	str := string(data)
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = CustomTime(t1)
	return err
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t CustomTime) Value() (driver.Value, error) {
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t *CustomTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		*t = CustomTime(vt) // 字符串转成 time.Time 类型
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *CustomTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}
