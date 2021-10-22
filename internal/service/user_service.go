package service

import (
	"encoding/base64"
	. "gin-blog/internal/models"
	. "gin-blog/internal/pkg/common"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type UserService struct {
	User
	CommonParams
}

func (s *UserService) QueryParams() map[string]interface{} {
	where := make(map[string]interface{})
	if s.Id > 0 {
		where["id"] = s.Id
	}

	return where
}

// 列表
func (s *UserService) FindList(where map[string]interface{}) ([]*User, error) {
	var entities []*User
	offset := GetOffset(s.Limit, s.Page)
	err := NewQuery().Model(&User{}).Where(where).Offset(offset).Limit(s.Limit).Find(&entities).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return entities, nil
}

// 总数
func (s *UserService) Count(where map[string]interface{}) (int, error) {
	var total int
	if err := NewQuery().Model(&User{}).Where(where).Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

// 新增
func (s *UserService) AddUser() error {
	var entity User
	entity.Status = StatusOn
	entity.Username = s.Username
	entity.Password = s.Encrypt(entity.Password)
	entity.Desc = s.Desc
	entity.Role = "member"
	if err := NewQuery().Create(&entity).Error; err != nil {
		return err
	}

	return nil
}

// 更新
func (s *UserService) EditUser(id int) error {
	var maps = make(map[string]interface{})
	maps["username"] = s.Username
	maps["password"] = s.Encrypt(s.Password)
	maps["status"] = s.Status
	maps["desc"] = s.Desc
	maps["role"] = "member"

	err := NewQuery().Model(&User{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return err
	}

	return nil
}

// 删除
func (s *UserService) DelUser(id int) error {
	if err := NewQuery().Where("id = ?", id).Delete(&User{}).Error; err != nil {
		return err
	}

	return nil
}

// 查找单条
func (s *UserService) FindById(id int) (bool, *User) {
	var entity User
	err := NewQuery().Select("id").Where("id = ?", id).Find(&entity).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, nil
	}

	if entity.ID > 0 {
		return true, &entity
	}

	return false, nil
}

// 查找单条
func (s *UserService) FindByUserName(username string) (bool, *User) {
	var entity User
	err := NewQuery().Select("id").Where("username = ?", username).Find(&entity).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, nil
	}

	if entity.ID > 0 {
		return true, &entity
	}

	return false, nil
}

// 加密
func (s *UserService) Encrypt(password string) string {
	const KeyLen = 8
	salt := make([]byte, 8)
	salt = []byte{5, 48, 12, 33, 85, 6, 10, 9}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	finalPw := base64.StdEncoding.EncodeToString(HashPw)
	return finalPw
}

// 校验
func (s *UserService) CheckLogin(username string, password string) (code int, user *User) {
	var entity User
	NewQuery().Where("username = ?", username).First(&entity)
	if entity.ID <= 0 {
		return 404, nil
	}
	if s.Encrypt(password) != entity.Password {
		return 403, nil
	}
	return 200, &entity
}
