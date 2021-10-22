package service

import (
	. "gin-blog/internal/models"
	. "gin-blog/internal/pkg/common"
	"github.com/jinzhu/gorm"
)

type ArticleService struct {
	Article
	CommonParams
}

// 查询参数
func (s *ArticleService) QueryParams() map[string]interface{} {
	where := make(map[string]interface{})
	if s.Status >= 0 {
		where["status"] = s.Status
	}
	if s.Id > 0 {
		where["id"] = s.Id
	}

	return where
}

// 文章列表
func (s *ArticleService) FindList(where map[string]interface{}) ([]*Article, error) {
	var entities []*Article
	offset := GetOffset(s.Limit, s.Page)
	db := NewQuery()
	if len(s.Title) > 0 {
		db = db.Where("title LIKE ?", s.Title+"%")
	}
	err := db.Model(&Article{}).Preload("User").Preload("Tag").Preload("Category").Where(where).Offset(offset).Limit(s.Limit).Find(&entities).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return entities, nil
}

// 文章总数
func (s *ArticleService) Count(where map[string]interface{}) (int, error) {
	var total int
	db := NewQuery()
	if len(s.Title) > 0 {
		db = db.Where("title LIKE ?", s.Title+"%")
	}
	if err := db.Model(&Article{}).Where(where).Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

// 新增文章
func (s *ArticleService) AddArticle() error {
	var entity Article
	entity.Title = s.Title
	entity.TagID = s.TagID
	entity.Status = s.Status
	entity.Desc = s.Desc
	entity.Content = s.Content
	entity.CategoryID = s.CategoryID
	entity.UserID = s.UserID
	if err := NewQuery().Create(&entity).Error; err != nil {
		return err
	}

	return nil
}

// 更新文章
func (s *ArticleService) EditArticle(id int) error {
	var maps = make(map[string]interface{})
	maps["title"] = s.Title
	maps["tag_id"] = s.TagID
	maps["desc"] = s.Desc
	maps["status"] = s.Status
	maps["content"] = s.Content
	maps["category_id"] = s.CategoryID

	err := NewQuery().Model(&Article{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return err
	}

	return nil
}

// 删除文章
func (s *ArticleService) DelArticle(id int) error {
	if err := NewQuery().Where("id = ?", id).Delete(&Article{}).Error; err != nil {
		return err
	}

	return nil
}

// 查找单条文章
func (s *ArticleService) FindById(id int) (bool, *Article) {
	var entity Article
	err := NewQuery().Select("id").Where("id = ?", id).Find(&entity).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, nil
	}

	if entity.ID > 0 {
		return true, &entity
	}

	return false, nil
}
