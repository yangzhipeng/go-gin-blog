package service

import (
	. "gin-blog/internal/models"
	. "gin-blog/internal/pkg/common"
	"github.com/jinzhu/gorm"
)

type TagService struct {
	Tag
	CommonParams
}

// 查询参数
func (s *TagService) QueryParams() map[string]interface{} {
	where := make(map[string]interface{})
	if s.Id > 0 {
		where["id"] = s.Id
	}
	if s.Status >= 0 {
		where["status"] = s.Status
	}

	return where
}

// 标签列表
func (s *TagService) FindList(where map[string]interface{}) ([]*Tag, error) {
	var entities []*Tag
	offset := GetOffset(s.Limit, s.Page)
	db := NewQuery()
	if len(s.Name) > 0 {
		db = db.Where("name LIKE ?", s.Name+"%")
	}
	err := db.Model(&Tag{}).Where(where).Offset(offset).Limit(s.Limit).Find(&entities).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return entities, nil
}

// 标签总数
func (s *TagService) Count(where map[string]interface{}) (int, error) {
	var total int
	db := NewQuery()
	if len(s.Name) > 0 {
		db = db.Where("name LIKE ?", s.Name+"%")
	}
	if err := db.Model(&Tag{}).Where(where).Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

// 新增标签
func (s *TagService) AddTag() error {
	var entity Tag
	entity.Name = s.Name
	entity.Status = s.Status
	if err := NewQuery().Create(&entity).Error; err != nil {
		return err
	}

	return nil
}

// 更新标签
func (s *TagService) EditTag(id int) error {
	var maps = make(map[string]interface{})
	maps["name"] = s.Name
	maps["status"] = s.Status

	err := NewQuery().Model(&Tag{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return err
	}

	return nil
}

// 删除标签
func (s *TagService) DelTag(id int) error {
	if err := NewQuery().Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}

	return nil
}

// 查找单条标签
func (s *TagService) FindById(id int) (bool, *Tag) {
	var entity Tag
	err := NewQuery().Select("id").Where("id = ?", id).Find(&entity).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, nil
	}

	if entity.ID > 0 {
		return true, &entity
	}

	return false, nil
}
