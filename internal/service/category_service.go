package service

import (
	. "gin-blog/internal/models"
	. "gin-blog/internal/pkg/common"
	"github.com/jinzhu/gorm"
)

type CategoryService struct {
	Category
	CommonParams
}

// 查询参数
func (s *CategoryService) QueryParams() map[string]interface{} {
	where := make(map[string]interface{})
	if s.Id > 0 {
		where["id"] = s.Id
	}

	return where
}

// 分类列表
func (s *CategoryService) FindList(where map[string]interface{}) ([]*Category, error) {
	var entities []*Category
	offset := GetOffset(s.Limit, s.Page)
	err := NewQuery().Model(&Category{}).Where(where).Offset(offset).Limit(s.Limit).Find(&entities).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return entities, nil
}

// 分类总数
func (s *CategoryService) Count(where map[string]interface{}) (int, error) {
	var total int
	if err := NewQuery().Model(&Category{}).Where(where).Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

// 新增分类
func (s *CategoryService) AddCategory() error {
	var entity Category
	entity.Name = s.Name
	entity.Status = s.Status
	entity.Desc = s.Desc
	if err := NewQuery().Create(&entity).Error; err != nil {
		return err
	}

	return nil
}

// 更新分类
func (s *CategoryService) EditCategory(id int) error {
	var maps = make(map[string]interface{})
	maps["name"] = s.Name
	maps["status"] = s.Status
	maps["desc"] = s.Desc

	err := NewQuery().Model(&Category{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return err
	}

	return nil
}

// 删除分类
func (s *CategoryService) DelCategory(id int) error {
	if err := NewQuery().Where("id = ?", id).Delete(&Category{}).Error; err != nil {
		return err
	}

	return nil
}

// 查找单条分类
func (s *CategoryService) FindById(id int) (bool, *Category) {
	var entity Category
	err := NewQuery().Select("id").Where("id = ?", id).Find(&entity).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, nil
	}

	if entity.ID > 0 {
		return true, &entity
	}

	return false, nil
}
